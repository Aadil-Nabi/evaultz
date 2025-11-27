package userhandlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUpHandler(c *gin.Context) {
	type userDetails struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
		Tenant   string `json:"tenant"`
		Team     string `json:"team"`
	}

	var body userDetails

	// Parse input
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Hash password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// Run everything inside a transaction (atomic)
	err = configs.DB.Transaction(func(tx *gorm.DB) error {

		// 1️⃣ Find or create tenant
		var tenant models.Tenant
		if err := tx.Where("name = ?", body.Tenant).First(&tenant).Error; err != nil {
			// Create tenant if not exists
			if errors.Is(err, gorm.ErrRecordNotFound) {
				tenant = models.Tenant{Name: body.Tenant}
				if err := tx.Create(&tenant).Error; err != nil {
					return fmt.Errorf("failed to create tenant: %w", err)
				}
			} else {
				return err
			}
		}

		// 2️⃣ Find or create team (if provided)
		var teamID *uuid.UUID = nil

		if strings.TrimSpace(body.Team) != "" {
			var team models.Team

			if err := tx.Where("name = ? AND tenant_id = ?", body.Team, tenant.ID).First(&team).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					team = models.Team{
						Name:     body.Team,
						TenantID: tenant.ID,
					}
					if err := tx.Create(&team).Error; err != nil {
						return fmt.Errorf("failed to create team: %w", err)
					}
				} else {
					return err
				}
			}

			// Assign team ID pointer
			tid := team.ID
			teamID = &tid
		}

		// 3️⃣ Ensure email is unique within the tenant
		var existing models.User
		if err := tx.Where("email = ? AND tenant_id = ?", body.Email, tenant.ID).First(&existing).Error; err == nil {
			return fmt.Errorf("user with this email already exists in tenant")
		}

		// 4️⃣ Create user
		user := models.User{
			Email:    body.Email,
			Password: string(hashedPass),
			Username: body.Username,
			TenantID: tenant.ID,
			TeamID:   teamID,
		}

		if err := tx.Create(&user).Error; err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		return nil
	})

	// Handle transaction error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Success
	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

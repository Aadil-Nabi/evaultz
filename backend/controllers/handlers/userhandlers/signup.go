package userhandlers

import (
	"errors"
	"fmt"
	"log"
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

	// Parse JSON
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Hash Password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// Wrap everything in a transaction
	err = configs.DB.Transaction(func(tx *gorm.DB) error {

		// 1️⃣ Find or create Tenant
		var tenant models.Tenant
		if err := tx.Where("name = ?", body.Tenant).First(&tenant).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				tenant = models.Tenant{Name: body.Tenant}
				if err := tx.Create(&tenant).Error; err != nil {
					return fmt.Errorf("failed to create tenant: %w", err)
				}
			} else {
				return err
			}
			log.Println("TENANT CREATED is :", tenant)
		}

		// 2️⃣ (Optional) Find or create Team
		var teamPtr *uuid.UUID = nil

		if strings.TrimSpace(body.Team) != "" {

			var team models.Team
			err := tx.Where("name = ? AND tenant_id = ?", body.Team, tenant.ID).First(&team).Error

			if err != nil {
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

			// Assign pointer
			teamID := team.ID
			teamPtr = &teamID
		}

		// 3️⃣ Validate duplicate user inside same tenant
		var existing models.User
		if err := tx.Where("email = ? AND tenant_id = ?", body.Email, tenant.ID).First(&existing).Error; err == nil {
			return fmt.Errorf("user with this email already exists in this tenant")
		}

		// 4️⃣ Create User
		user := models.User{
			Email:    body.Email,
			Password: string(hashedPass),
			Username: body.Username,
			TenantID: tenant.ID,
			TeamID:   teamPtr, // pointer now!
		}

		if err := tx.Create(&user).Error; err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

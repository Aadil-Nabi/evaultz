package userhandlers

import (
	"net/http"

	"github.com/Aadil-Nabi/evaultz/configs"
	"github.com/Aadil-Nabi/evaultz/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DeleteUser(c *gin.Context) {
	// 1️⃣ Extract userID from context (set by middleware)
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	userID := userIDVal.(uuid.UUID)

	// 2️⃣ Fetch user with files
	var user models.User
	err := configs.DB.Preload("Files").First(&user, "id = ?", userID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// 3️⃣ Begin transaction
	tx := configs.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to start transaction"})
		return
	}

	// 4️⃣ Delete all S3 files owned by user. **********____IMPORTANT____*********
	// for _, file := range user.Files {
	// 	if err := awsS3Delete(file.Key); err != nil {
	// 		tx.Rollback()
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"error":   "failed to delete file from S3",
	// 			"details": err.Error(),
	// 		})
	// 		return
	// 	}
	// }

	// 5️⃣ Delete file records
	if err := tx.Where("owner_id = ?", user.ID).Delete(&models.File{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete file records"})
		return
	}

	// 6️⃣ Delete user (team and tenant remain untouched)
	if err := tx.Delete(&models.User{}, "id = ?", user.ID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	// 7️⃣ Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to commit transaction"})
		return
	}

	// 8️⃣ Clear JWT cookie
	c.SetCookie("jwt", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}

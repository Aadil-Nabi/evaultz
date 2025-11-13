package filehandlers

import (
	"context"
	"net/http"

	"github.com/Aadil-Nabi/evaultz/utility/awsclient"
	"github.com/gin-gonic/gin"
)

func DownloadHander(c *gin.Context) {

	bucketService, err := awsclient.NewBucketBasics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	// if file name is passed in a param insted of a JSON body use below.
	// /api/v1/download/:filename
	filename := c.Param("filename")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "filename param is required",
		})
	}

	// key := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), fileDetails.filename)
	// key := fileDetails.Filename
	key := filename

	file, err := bucketService.DownloadFile(context.TODO(), key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.Header("Content-Disposition", "attachment; filename="+key)
	c.Data(http.StatusOK, "application/octet-stream", file)

	// c.JSON(http.StatusOK, gin.H{
	// 	"result": "file downloaded successfully",
	// 	"file":   file,
	// })

}

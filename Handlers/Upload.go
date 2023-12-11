package Handlers

import (
	"anonim-fileshare/Utils"
	
	"io"
	"net/http"
	"os"
	"path/filepath"
	

	"github.com/gin-gonic/gin"
)
func Upload(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	newFilename := Utils.ID()
	storagePath := filepath.Join("shared", newFilename)
	out, err := os.Create(storagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

    url := "http://27.0.0.1:5000/dw/"+newFilename 
 
	c.JSON(http.StatusOK, gin.H{
		"message": "Dosya başarıyla yüklendi", 
		"filename": newFilename,
		"url": url,
	})
}
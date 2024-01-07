package handlers

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

const chars = "abcdefghijklmnopqrstuvwxyz0123456789"

func generateFolderName() string {
  b := make([]byte,6) 
	rand.Read(b)
	for i := range b {
		b[i] = chars[int(b[i])%len(chars)]
	}
	return string(b)

}

func UploadFile(c *gin.Context){
		fileName := c.Param("filename")
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

  		// Save the raw data to a file
    folderName := generateFolderName()
    writePath := fmt.Sprintf("uploads/%s/%s", folderName, fileName)
    os.Mkdir("uploads/" + folderName, 0777)
		err = os.WriteFile(writePath, body, 0777)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

    downloadUrl := fmt.Sprintf("http://localhost:8080/uploads/%s/%s ", folderName, fileName)

		c.String(200, downloadUrl)
}

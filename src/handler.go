package src

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "hello world",
	})
}

func Deploy(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// get signature header
	header := c.GetHeader("X-Hub-Signature-256")
	if header == "" {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "missing header",
		})
		return
	}

	secret := os.Args[2]

	if valid := verifySignature(secret, header, string(body)); valid != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": valid.Error(),
		})
		return
	}

	pullRepo(os.Args[1])

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "done",
	})
}

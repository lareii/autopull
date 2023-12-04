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

	Log(2, "hello world", true)
}

func Deploy(c *gin.Context) {
	// check if req is just ping
	if ping := c.GetHeader("X-GitHub-Event"); ping == "ping" {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "pong",
		})

		Log(2, "pong", true)
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err.Error(),
		})

		Log(0, err.Error(), true)
		return
	}

	// get signature header
	header := c.GetHeader("X-Hub-Signature-256")
	if header == "" {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "missing header",
		})

		Log(0, "missing header", true)
		return
	}

	secret := os.Args[2]

	if valid := verifySignature(secret, header, string(body)); valid != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": valid.Error(),
		})

		Log(0, valid.Error(), true)
		return
	}

	pullRepo(os.Args[1])

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "done",
	})

	Log(2, "autopull triggered. pulling the repo", true)
}

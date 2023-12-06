package handlers

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lareii/autopull/utils"
)

/*
 * Deploy is a handler for the /deploy endpoint.

 * If the request is ping, it returns "pong" message.
 * Else, it pulls and returns "done" message the repo if signature is valid.
 */

func Deploy(c *gin.Context) {
	// Check if the request is ping.
	if c.GetHeader("X-GitHub-Event") == "ping" {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "pong",
		})

		utils.Log(2, "pong", true)

		return
	}

	// Get signature header and check if it is empty.
	header := c.GetHeader("X-Hub-Signature-256")

	if header == "" {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "missing header",
		})

		utils.Log(0, "missing header", true)
		return
	}

	// Get request body and check if there is an error.
	body, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": err.Error(),
		})

		utils.Log(0, err.Error(), true)

		return
	}

	// Get secret from the command line arguments.
	secret := os.Args[2]

	// Check if our signature is valid, if not return error.
	if err := utils.Signature(secret, header, string(body)); err != nil {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": err.Error(),
		})

		utils.Log(0, err.Error(), true)
		return
	}

	// Pull the repo and return "done" message.
	utils.Pull(os.Args[1])

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "done",
	})

	// And finally log the event.
	utils.Log(2, "autopull triggered. pulling the repo", true)
}

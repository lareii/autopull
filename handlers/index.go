package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lareii/autopull/utils"
)

/*
 * Index is a handler for the / endpoint.

 * It returns "hello world" message for testing purposes.
 */

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "hello world",
	})

	utils.Log(2, "hello world", true)
}

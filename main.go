package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/lareii/autopull/handlers"
	"github.com/lareii/autopull/utils"
)

func init() {
	utils.Log(2, "initializing", true)

	// Check if the required parameters are present.
	if args := os.Args[1:]; len(args) != 2 {
		utils.Log(0, "some parameters are missing", true)
		os.Exit(0)
	}

	// Check if the specified directory is a git repository.
	if err := utils.Check(os.Args[1]); err != nil {
		utils.Log(0, err.Error(), true)
		os.Exit(0)
	}

	utils.Clear()

	repo, _ := utils.Repository(os.Args[1])
	utils.Log(2, "directory: "+os.Args[1], true)
	utils.Log(2, "repository: "+repo, true)

	println()

	for {
		print("is everything correct? (y/n): ")

		var ok string
		fmt.Scan(&ok)

		if ok[0] == 'y' {
			break
		} else if ok[0] == 'n' {
			os.Exit(0)
		} else {
			println("nuh uh")
		}
	}

	println()

	gin.SetMode("debug")
	gin.DefaultWriter = io.Discard
}

func main() {
	utils.Log(2, "server listening on :3000", true)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", handlers.Index)
	router.POST("/deploy", handlers.Deploy)

	router.Run(":3000")
}

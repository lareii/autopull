package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"

	s "github.com/lareii/autopull/src"
)

func init() {
	println("Initializing...")

	gin.SetMode("debug")
	gin.DefaultWriter = io.Discard
}

func main() {
	if args := os.Args[1:]; len(args) != 2 {
		return
	}

	println("Listening on :3000")

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", s.Index)
	router.POST("/deploy", s.Deploy)

	router.Run(":3000")
}

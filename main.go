package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	s "github.com/lareii/autopull/src"
)

func init() {
	println("Initializing...")
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env duzelt abi")
	}

	/*
		mode := os.Getenv("GIN_MODE")
		gin.SetMode(mode)
		gin.DefaultWriter = io.Discard
	*/
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

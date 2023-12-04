package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"

	s "github.com/lareii/autopull/src"
)

func init() {
	s.Clear()

	s.Log(2, "initializing", true)

	if args := os.Args[1:]; len(args) != 2 {
		s.Log(0, "some parameters are missing", true)
		os.Exit(0)
	}

	if err := s.CheckDir(os.Args[1]); err != nil {
		s.Log(0, err.Error(), true)
		os.Exit(0)
	}

	repo, _ := s.GetRepoURL(os.Args[1])
	s.Log(2, "directory: "+os.Args[1], true)
	s.Log(2, "repository: "+repo, true)

	println()
	println("Is everything correct?")
	println("[0] Yes")
	println("[1] Cancel")
	println()
	print("> ")

	var ok string
	fmt.Scan(&ok)

	if ok == "1" {
		os.Exit(0)
	}

	println()

	gin.SetMode("debug")
	gin.DefaultWriter = io.Discard
}

func main() {
	s.Log(2, "Server listening on :3000.", true)

	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", s.Index)
	router.POST("/deploy", s.Deploy)

	router.Run(":3000")
}

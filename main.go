package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tajale72/file-converter/internal/router"
)

func main() {
	r := gin.Default()
	r.GET("/", router.Fie)
	router.Fi
}

package main

import (
	"context"
	"log"

	"go-mongo/internal/app"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	app.Routes(r, context.Background())

	log.Fatal(r.Run(":8000"))
}

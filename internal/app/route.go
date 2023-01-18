package app

import (
	"context"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, ctx context.Context) error {
	app, err := NewApp(ctx)
	if err != nil {
		return err
	}

	router.GET("/users", app.UserHandler.All)
	router.POST("/users", app.UserHandler.Insert)
	router.GET("/users/:id", app.UserHandler.Load)
	router.PUT("/users/:id", app.UserHandler.Update)
	router.DELETE("/users/:id", app.UserHandler.Delete)

	return nil
}

package app

import (
	"context"
	"log"
	"os"

	"go-mongo/internal/handler"
	"go-mongo/internal/service"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type ApplicationContext struct {
	UserHandler *handler.UserHandler
}

func NewApp(ctx context.Context) (*ApplicationContext, error) {
	var _ = godotenv.Load(".env")
	uri := os.Getenv("uri")

	clientOptions := options.Client().ApplyURI(uri)
	client, _ := mongo.NewClient(clientOptions)

	err := client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database("gin-mongo")

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database!")
	} else {
		log.Println("Connected to the database!")
	}

	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(userService)

	return &ApplicationContext{
		UserHandler: userHandler,
	}, nil
}

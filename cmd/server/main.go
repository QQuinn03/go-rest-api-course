package main

import (
	"context"
	"fmt"

	"github.com/QQuinn03/go-rest-api-course/internal/comment"
	"github.com/QQuinn03/go-rest-api-course/internal/db"
	transportHttp "github.com/QQuinn03/go-rest-api-course/internal/transport/http"
)

// Run - is responsible for instantiation and sets up our go application
func Run() error {
	fmt.Println("Starting Up Our APP")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate databse")
		return err
	}

	cmtService := comment.NewService((db)) //return service struct{store}
	cmtService.GettComment(context.Background(),
		"71c5d074-b6cf-11ec-b909-0242ac120002")

	httpHandler := transportHttp.NewHandler(cmtService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil

}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}

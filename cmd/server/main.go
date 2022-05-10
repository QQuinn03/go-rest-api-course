package main

import (
	"fmt"

	"github.com/QQuinn03/go-rest-api-course/internal/comment"
	"github.com/QQuinn03/go-rest-api-course/internal/db"
	transportHttp "github.com/QQuinn03/go-rest-api-course/internal/transport/http"
)

// App - the struct which contains things like
// pointers to database connections
/*type App struct {
}*/

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
	/*cmtService.PostComment(context.Background(),
		comment.Comment{
			ID:     "71c5d074-b6cf-11ec-b909-0242ac120002",
			Slug:   "TEST if this works",
			Author: "QQ",
			Body:   "hi!",
		},
	)*/

	httpHandler := transportHttp.NewHandler(cmtService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}

	/*fmt.Println(cmtService.GetComment(context.Background(),
		"71c5d074-b6cf-11ec-b909-0242ac120002",
	))*/

	/*if err := db.Ping(context.Background()); err != nil {
		return err
	}*/

	return nil

}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}

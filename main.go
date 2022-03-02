package main

import (
	"fmt"
	"net/http"

	transportApp "github.com/QQuinn03/go-rest-api-course/internal/transport/http"
)

//this struct includes things like pointers
//to databse connections
type App struct{}

//Run-sets up our application
func (app *App) Run() error {
	fmt.Println("setting up our App")

	hanlder := transportApp.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Fail to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go rest api")
	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("error on rest api")
		fmt.Println(err)
	}

}

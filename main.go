package main

import "fmt"

//this struct includes things like pointers
//to databse connections
type App struct{}

//Run-sets up our application
func (app *App) Run() error {
	fmt.Println("setting up our App")
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

package main

import "fmt"

type App struct{}

func (*App) Run() error {
	fmt.Println("Setting up our APP")
	return nil

}
func main() {
	//fmt.Println("h")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our rest api")
		fmt.Println(err)
	}
}

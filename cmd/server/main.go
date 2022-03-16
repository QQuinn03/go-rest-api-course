package main

import "fmt"

// App - the struct which contains things like
// pointers to database connections
type App struct {
}

// Run - is responsible for instantiation and sets up our go application
func Run() error {
	fmt.Println("Starting Up Our APP")

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}

}

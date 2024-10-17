package main

import "github.com/devops-journey/src/app"

func main() {
	router := app.SetUpRouter()
	router.Run(":8080")
}

package main

import "github.com/devops-journey/src/app" 


func main() {
	router := app.SetUpRouter()
	router.GET("/albums", app.GetAlbums)

	router.Run(":8000")
}

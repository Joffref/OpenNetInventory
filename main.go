package main

import "github.com/joffref/netrman/routes"

func main() {
	router := routes.SetupRouter()
	router.Run()
}

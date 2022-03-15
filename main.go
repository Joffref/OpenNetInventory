package main

import "github.com/joffref/openNetInventory/routes"

func main() {
	router := routes.SetupRouter()
	router.Run()
}

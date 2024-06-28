package main

import (
	"fmt"
	"go-chat/routes"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Hello World")
	routes.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))

}

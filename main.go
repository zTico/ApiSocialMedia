package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Loading()

	fmt.Println(config.ConectDb)

	fmt.Println("Rodando api")

	router := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}

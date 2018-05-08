package main

import (
	"DTmocker/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.ShowApiList)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

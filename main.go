package main

import (
	"DTmocker/controller"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/", controller.ShowApiList)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"log"
	"net/http"

	"github.com/FraserTooth/go-examine/cmd/webpageanalyser"
)

func main() {
	http.HandleFunc("/", webpageanalyser.AnalyseWebpage)
	address := ":8000"
	log.Println("Starting server on address", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}

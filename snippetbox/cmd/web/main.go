package main

import (
	"net/http"
	"log"
)

func main(){

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)


	log.Println("Starting server on 4248")
	err := http.ListenAndServe(":4248", mux)
	log.Fatal(err)

}

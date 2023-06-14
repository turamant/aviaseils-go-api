package main

import (

	"log"
	"net/http"

)

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/company", showCompany)
	mux.HandleFunc("/company/create", createCompany)
	
	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}


package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sinhamanav03/backend-api-assignment/controller"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/find-pairs", controller.FindPairController).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("Failed to start server", err)
	}

}

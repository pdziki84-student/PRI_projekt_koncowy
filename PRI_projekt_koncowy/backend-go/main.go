package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pri/todos/Controllers"

	"github.com/gorilla/mux"
)

func main() {
	lf, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer lf.Close()
	log.SetOutput(lf)

	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("content-type", "application/json")
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE")
			w.Header().Add("Access-Control-Allow-Headers", "*")
		})
	r.HandleFunc("/items", Controllers.GetTasks).Methods("GET")
	r.HandleFunc("/items", Controllers.CreateTask).Methods("POST")
	r.HandleFunc("/items/{id}", Controllers.DeleteTask).Methods("DELETE")
	fmt.Println("Server listening")
	http.ListenAndServe(":8081", r)
}

package main

import(
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"handlers"
)
func main(){
	r:= mux.NewRouter()

	r.HandleFunc("api/commits",handlers.GetCommits).Methods("GET")
	r.HandleFunc("/api/contributors", handlers.GetContributors).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
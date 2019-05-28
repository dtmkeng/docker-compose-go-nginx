package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type cars struct {
	Version int
	Typecar string
	Price   int
	Color   string
}

const (
	port = "9001"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	fmt.Println("Server start at " + port + " ...")
	http.ListenAndServe(":"+port, r)

}
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	carCrate := cars{
		Version: 1,
		Typecar: "SUV",
		Price:   1000000,
		Color:   "Red",
	}
	json.NewEncoder(w).Encode(carCrate)
}

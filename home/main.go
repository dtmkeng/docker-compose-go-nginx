package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type homes struct {
	Version  int
	Typehome string
	Price    int
	Location string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	fmt.Println("Server start at 9000 ...")
	http.ListenAndServe(":9000", r)

}
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	homeCreate := homes{
		Version:  1,
		Typehome: "single",
		Price:    12236222,
		Location: "Thanland",
	}
	json.NewEncoder(w).Encode(homeCreate)
}

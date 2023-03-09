package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Gowapi struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var gowapis []Gowapi

func getGowapis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gowapis)
}

func deleteGowapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range gowapis {
		if item.ID == params["id"] {
			gowapis = append(gowapis[:index], gowapis[index+1:]...)
			break
		}
	}

}

func getGowapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range gowapis {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createGowapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var gowapi Gowapi
	_ = json.NewDecoder(r.Body).Decode(&gowapi)
	gowapi.ID = strconv.Itoa(rand.Intn(100000000))
	gowapis = append(gowapis, gowapi)
	json.NewEncoder(w).Encode(gowapi)
}

func updateGowapi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range gowapis {
		if item.ID == params["id"] {
			gowapis = append(gowapis[:index], gowapis[index+1:]...)
			var gowapi Gowapi
			_ = json.NewDecoder(r.Body).Decode(&gowapi)
			gowapi.ID = params["id"]
			gowapis = append(gowapis, gowapi)
			json.NewEncoder(w).Encode(gowapi)

		}
	}
}

func main() {
	r := mux.NewRouter()

	gowapis = append(gowapis, Gowapi{ID: "1", Isbn: "432877", Title: "Gow1", Director: &Director{Firstname: "Lexa", Lastname: "Otosaka"}})
	gowapis = append(gowapis, Gowapi{ID: "2", Isbn: "432876", Title: "Gow2", Director: &Director{Firstname: "Yuda", Lastname: "Motoyasu"}})
	gowapis = append(gowapis, Gowapi{ID: "3", Isbn: "432875", Title: "Gow3", Director: &Director{Firstname: "Emperoro", Lastname: "Kurayami"}})

	r.HandleFunc("/Gowapi", getGowapis).Methods("GET")
	r.HandleFunc("/Gowapi/{id}", getGowapi).Methods("GET")
	r.HandleFunc("/Gowapi", createGowapi).Methods("POST")
	r.HandleFunc("/Gowapi/{id}", updateGowapi).Methods("PUT")
	r.HandleFunc("/Gowapi/{id}", deleteGowapi).Methods("DELETE")

	fmt.Printf("Starting Server At Port 8000 ( server jalan di port 8000 ya )")
	log.Fatal(http.ListenAndServe(":8000", r))
}

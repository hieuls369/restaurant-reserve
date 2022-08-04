package main

import (
	controllers "example/restaurant-reserved/internal/controllers/reserve-controller"
	"example/restaurant-reserved/internal/providers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	providers.ProviderHandler()
	r.HandleFunc("/reserve", controllers.GetReserves()).Methods("GET")
	r.HandleFunc("/reserve", controllers.CreateReserve()).Methods("POST")
	r.HandleFunc("/reserve/{phoneNumber}", controllers.UpdateReserve()).Methods("PUT")
	r.HandleFunc("/reserve", controllers.DeleteReserve()).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

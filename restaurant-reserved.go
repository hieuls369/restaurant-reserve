package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Table struct {
	ID        string `json:"ID"`
	MaxPeople int    `json:"maxPeople"`
}

type Reserve struct {
	ID           string `json:"reserveID"`
	AmountPeople int    `json:"amountPeople"`
	PhoneNumber  string `json:"phoneNumber"`
	Date         string `json:"date"`
	TableID      string `json:"tableID"`
}

var ReserveList []Reserve
var TableList []Table

func getReserves(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["date"])
	if params["date"] != "" {
		var filterByDate []Reserve
		for _, item := range ReserveList {
			if item.Date == params["date"] {
				filterByDate = append(filterByDate, item)
			}
		}
		json.NewEncoder(w).Encode(filterByDate)
		return
	}
	json.NewEncoder(w).Encode(ReserveList)
}

func deleteReserve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range ReserveList {
		if item.PhoneNumber == params["phoneNumber"] {
			ReserveList = append(ReserveList[:index], ReserveList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(ReserveList)
}

func getReserve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range ReserveList {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createReserve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reserve Reserve
	_ = json.NewDecoder(r.Body).Decode(&reserve)
	reserve.ID = strconv.Itoa(rand.Intn(10000000))
	for _, item := range ReserveList {
		if item.Date == reserve.Date && item.TableID == reserve.TableID {
			fmt.Printf("The table has been reserved")
			return
		}
	}
	ReserveList = append(ReserveList, reserve)
	json.NewEncoder(w).Encode(reserve)
}

func updateReserve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range ReserveList {
		if item.PhoneNumber == params["phoneNumber"] {
			ReserveList = append(ReserveList[:index], ReserveList[index+1:]...)
			var reserve Reserve
			_ = json.NewDecoder(r.Body).Decode(&reserve)
			reserve.ID = params["id"]
			ReserveList = append(ReserveList, reserve)
			json.NewEncoder(w).Encode(reserve)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	TableList = append(TableList, Table{ID: "1", MaxPeople: 3})
	TableList = append(TableList, Table{ID: "2", MaxPeople: 2})
	TableList = append(TableList, Table{ID: "3", MaxPeople: 4})
	ReserveList = append(ReserveList, Reserve{ID: "1", AmountPeople: 2, PhoneNumber: "0332534234", Date: "21/02/2022", TableID: "1"})
	ReserveList = append(ReserveList, Reserve{ID: "2", AmountPeople: 1, PhoneNumber: "0732975234", Date: "21/02/2022", TableID: "2"})
	ReserveList = append(ReserveList, Reserve{ID: "3", AmountPeople: 3, PhoneNumber: "0162329234", Date: "04/11/2022", TableID: "3"})

	r.HandleFunc("/reserve", getReserves).Methods("GET").Queries("date", "{date}")
	r.HandleFunc("/reserve/{id}", getReserve).Methods("GET")
	r.HandleFunc("/reserve", createReserve).Methods("POST")
	r.HandleFunc("/reserve/{phoneNumber}", updateReserve).Methods("PUT")
	r.HandleFunc("/reserve/{phoneNumber}", deleteReserve).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

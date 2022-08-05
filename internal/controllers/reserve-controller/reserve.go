package reserve_controller

import (
	"context"
	"encoding/json"
	user_reserve "example/restaurant-reserved/internal/infrastructure/repositories/user-reserve"
	models "example/restaurant-reserved/internal/share-domain/reserve/reserve-model"
	"example/restaurant-reserved/tools/response"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetReserves() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var reserves []models.ReserveModel
		results, err := user_reserve.GetReserves(ctx)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			response := response.ReserveResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx) {
			var reserve models.ReserveModel
			if err = results.Decode(&reserve); err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				response := response.ReserveResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
				json.NewEncoder(w).Encode(response)
			}
			reserves = append(reserves, reserve)
		}

		w.WriteHeader(http.StatusOK)
		response := response.ReserveResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": reserves}}
		json.NewEncoder(w).Encode(response)
	}
}

func DeleteReserve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		id := r.URL.Query().Get("id")
		phonnNumber := r.URL.Query().Get("phoneNumber")

		fmt.Println(id, phonnNumber)
		result, err := user_reserve.DeleteReserve(ctx, id, phonnNumber)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			response := response.ReserveResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		if result.DeletedCount < 1 {
			w.WriteHeader(http.StatusNotFound)

			response := response.ReserveResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := response.ReserveResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}}
		json.NewEncoder(w).Encode(response)
	}
}

func CreateReserve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var reserve models.ReserveModel

		if err := json.NewDecoder(r.Body).Decode(&reserve); err != nil {
			w.WriteHeader(http.StatusBadRequest)

			response := response.ReserveResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		newReserve := models.ReserveModel{
			ReserveID:    reserve.ReserveID,
			AmountPeople: reserve.AmountPeople,
			PhoneNumber:  reserve.PhoneNumber,
			Date:         reserve.Date,
			TableID:      reserve.TableID,
		}
		result, err := user_reserve.CreateReserve(ctx, newReserve)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			response := response.ReserveResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		response := response.ReserveResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(w).Encode(response)
	}

}

func UpdateReserve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		params := mux.Vars(r)
		phoneNumber := params["phoneNumber"]
		var reserve models.ReserveModel

		if err := json.NewDecoder(r.Body).Decode(&reserve); err != nil {
			w.WriteHeader(http.StatusBadRequest)

			response := response.ReserveResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		result, err := user_reserve.UpdateReserve(ctx, phoneNumber, reserve)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			response := response.ReserveResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := response.ReserveResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(w).Encode(response)
	}
}

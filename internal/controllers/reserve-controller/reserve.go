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
		var reserves []models.ReserveModel
		defer cancel()

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

		// reserveList, _ := ReserveController.reserveService.GetReserves()
		// fmt.Println(reserveList)
		// params := mux.Vars(r)
		// if params["date"] != "" {
		// 	var filterByDate []models.ReserveModel
		// 	for _, item := range ReserveList {
		// 		if item.Date == params["date"] {
		// 			filterByDate = append(filterByDate, item)
		// 		}
		// 	}
		// 	json.NewEncoder(w).Encode(filterByDate)
		// 	return
		// }
		// json.NewEncoder(w).Encode(ReserveList)
	}
}

func DeleteReserve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := r.URL.Query().Get("id")
		phonnNumber := r.URL.Query().Get("phoneNumber")
		fmt.Println(id, phonnNumber)

		defer cancel()

		result, err := user_reserve.DeleteReserve(ctx, id)
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

		w.WriteHeader(http.StatusOK)
		response := response.ReserveResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}}
		json.NewEncoder(w).Encode(response)
	}
}

func CreateReserve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var reserve models.ReserveModel
		defer cancel()

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
		w.WriteHeader(http.StatusCreated)
		response := response.ReserveResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(w).Encode(response)
		// 	var reserve reserve_model.ReserveModel
		// 	_ = json.NewDecoder(r.Body).Decode(&reserve)
		// 	reserve.ID = strconv.Itoa(rand.Intn(10000000))
		// 	for _, item := range ReserveList {
		// 		if item.Date == reserve.Date && item.TableID == reserve.TableID {

		// 			fmt.Printf("The table has been reserved")
		// 			return
		// 		}
		// 	}
		// 	ReserveList = append(ReserveList, reserve)
		// 	json.NewEncoder(w).Encode(reserve)
	}

}

func UpdateReserve() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		phoneNumber := params["phoneNumber"]
		// phoneNumberConv, _ := strconv.ParseInt(phoneNumber, 0, 0)
		var reserve models.ReserveModel
		defer cancel()

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

		w.WriteHeader(http.StatusOK)
		response := response.ReserveResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(w).Encode(response)
	}
}

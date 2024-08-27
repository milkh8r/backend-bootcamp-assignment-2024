package handler

import (
	"avito-backend-bootcamp/internal/domain"
	"avito-backend-bootcamp/internal/usecase"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HouseHandler struct {
	houseUseCase *usecase.HouseUseCase
}

func NewHouseHandler(houseUseCase *usecase.HouseUseCase) *HouseHandler {
	return &HouseHandler{houseUseCase: houseUseCase}
}

func (h *HouseHandler) CreateHouse(w http.ResponseWriter, r *http.Request) {
	var house domain.House
	err := json.NewDecoder(r.Body).Decode(&house)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userRole := r.Header.Get("User-Role") // Assuming the user role is passed in the header

	err = h.houseUseCase.CreateHouse(&house, userRole)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(house)
}

func (h *HouseHandler) GetHouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid house ID", http.StatusBadRequest)
		return
	}

	house, err := h.houseUseCase.GetHouse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(house)
}

func (h *HouseHandler) ListHouses(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	houses, err := h.houseUseCase.ListHouses(page, pageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(houses)
}

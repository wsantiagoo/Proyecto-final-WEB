// Manejadores Gestion de autos
package handlers

import (
	"ColombiaCar/internal/services"
	"encoding/json"
	"net/http"
)

type CarHandler struct {
	carService services.CarService
}

func NewCarHandler(carService services.CarService) *CarHandler {
	return &CarHandler{carService: carService}
}

func (h *CarHandler) SearchCars(w http.ResponseWriter, r *http.Request) {
	filters := r.URL.Query()

	filterMap := make(map[string]string)
	for key, values := range filters {
		if len(values) > 0 {
			filterMap[key] = values[0]
		}
	}

	cars, err := h.carService.SearchCars(filterMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"resfulsimple/internal/models"
	"strconv"
)

// ProductHandler CRUD для продуктів
func ProductHandler(w http.ResponseWriter, r *http.Request) { //nolint:funlen
	switch r.Method {
	case http.MethodGet:
		if getProducts(w, r) {
			return
		}
	case http.MethodPost:
		if createProduct(w, r) {
			return
		}
	case http.MethodPut, http.MethodPatch:
		if updateProduct(w, r) {
			return
		}
	case http.MethodDelete:
		if deleteProduct(w, r) {
			return
		}
	default:
		http.Error(w, "Метод не підтримується", http.StatusMethodNotAllowed)
	}
}

func deleteProduct(w http.ResponseWriter, r *http.Request) bool {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Не вказано id", http.StatusBadRequest)
		return true
	}
	products := models.GetData()
	pid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, fmt.Errorf("can't parse product id:%w", err).Error(), http.StatusBadRequest)
		return true
	}
	products.Delete(pid)
	w.WriteHeader(http.StatusOK)
	return false
}

func updateProduct(w http.ResponseWriter, r *http.Request) bool {
	//id := r.URL.Query().Get("id")
	//if id == "" {
	//	http.Error(w, "Не вказано id", http.StatusBadRequest)
	//	return true
	//}
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}
	products := models.GetData()
	if err := products.Update(product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}
	w.WriteHeader(http.StatusAccepted)
	return false
}

func createProduct(w http.ResponseWriter, r *http.Request) bool {
	var newProduct models.Product
	data, _ := io.ReadAll(r.Body)
	_ = r.Body.Close()
	if err := json.Unmarshal(data, &newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("can't parse request: %v: %s", err, string(data))
		return true
	}
	products := models.GetData()
	if products.Find(newProduct.ID) != nil {
		http.Error(w, "Товар з таким id вже існує", http.StatusAlreadyReported)
		return true
	}
	products.Add(newProduct)
	w.WriteHeader(http.StatusCreated)
	return false
}

func getProducts(w http.ResponseWriter, r *http.Request) bool {
	id := r.URL.Query().Get("id")
	switch id {
	case "":
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(models.GetData().Get()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		pid, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, fmt.Errorf("can't parse product id:%w", err).Error(), http.StatusBadRequest)
			return true
		}
		product := models.GetData().Find(pid)
		if product == nil {
			http.Error(w, "Товар не знайдено", http.StatusNotFound)
			return true
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	return false
}

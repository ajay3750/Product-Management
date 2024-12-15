package handlers

import (
	"encoding/json"
	"net/http"
	"product-management/models"
	"product-management/queue"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		logrus.WithError(err).Error("Failed to decode request body")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	productID, err := models.AddProduct(product)
	if err != nil {
		logrus.WithError(err).Error("Failed to add product")
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	for _, imageURL := range product.ProductImages {
		err = queue.PublishMessage("image_queue", imageURL)
		if err != nil {
			logrus.WithError(err).Error("Failed to publish image URL")
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"product_id": productID,
	})
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	product, err := models.GetProductByID(id)
	if err != nil {
		logrus.WithError(err).Error("Product not found")
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func GetProductsByUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	products, err := models.GetProductsByUser(userID)
	if err != nil {
		logrus.WithError(err).Error("Failed to fetch products")
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		logrus.WithError(err).Error("Failed to decode request body")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	err = models.UpdateProduct(id, product)
	if err != nil {
		logrus.WithError(err).Error("Failed to update product")
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

package models

import "time"

type Product struct {
	ID                      int       `json:"id"`
	UserID                  int       `json:"user_id"`
	ProductName             string    `json:"product_name"`
	ProductDesc             string    `json:"product_description"`
	ProductPrice            float64   `json:"product_price"`
	ProductImages           []string  `json:"product_images"`
	CompressedProductImages []string  `json:"compressed_product_images"`
	CreatedAt               time.Time `json:"created_at"`
}

func AddProduct(product Product) (int, error) {
	return 1, nil
}
func GetProductByID(id string) (Product, error) {
	return Product{
		ID:           1,
		ProductName:  "Samsung S21 Ultra",
		ProductDesc:  "Flagship smartphone with high-end features",
		ProductPrice: 1200.0,
	}, nil
}
func UpdateProduct(id string, updatedProduct Product) error {
	return nil
}

func GetProductsByUser(userID string) ([]Product, error) {
	return []Product{
		{ID: 1, ProductName: "Samsung S21 Ultra", ProductDesc: "Flagship smartphone with high-end features", ProductPrice: 1200.0},
	}, nil
}

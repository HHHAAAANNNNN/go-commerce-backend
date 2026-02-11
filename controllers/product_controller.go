package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HHHAAAANNNNN/go-commerce-backend/config"
	"github.com/HHHAAAANNNNN/go-commerce-backend/models"
	"github.com/HHHAAAANNNNN/go-commerce-backend/utils"
	"github.com/gorilla/mux"
)

// GetAllProducts - GET /api/products
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, name, price, stock, category, rating, created_at FROM products ORDER BY created_at DESC`

	rows, err := config.DB.Query(query)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to fetch products")
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock,
			&product.Category, &product.Rating, &product.CreatedAt)
		if err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to scan product")
			return
		}
		products = append(products, product)
	}

	utils.SuccessResponse(w, "Products fetched successfully", products)
}

// GetProductByID - GET /api/products/{id}
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	query := `SELECT id, name, price, stock, category, rating, created_at FROM products WHERE id = ?`

	var product models.Product
	err := config.DB.QueryRow(query, id).Scan(
		&product.ID, &product.Name, &product.Price, &product.Stock,
		&product.Category, &product.Rating, &product.CreatedAt,
	)
	if err != nil {
		utils.ErrorResponse(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.SuccessResponse(w, "Product fetched successfully", product)
}

// CreateProduct - POST /api/products
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req models.ProductCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validation
	if req.ID == "" || req.Name == "" || req.Price <= 0 {
		utils.ErrorResponse(w, http.StatusBadRequest, "ID, name, and price are required")
		return
	}

	query := `INSERT INTO products (id, name, price, stock, category, rating) VALUES (?, ?, ?, ?, ?, ?)`
	_, err = config.DB.Exec(query, req.ID, req.Name, req.Price, req.Stock, req.Category, req.Rating)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.CreatedResponse(w, "Product created successfully", req)
}

// UpdateProduct - PUT /api/products/{id}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req models.ProductUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	query := `UPDATE products SET name = ?, price = ?, stock = ?, category = ?, rating = ? WHERE id = ?`
	result, err := config.DB.Exec(query, req.Name, req.Price, req.Stock, req.Category, req.Rating, id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to update product")
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		utils.ErrorResponse(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.SuccessResponse(w, "Product updated successfully", nil)
}

// DeleteProduct - DELETE /api/products/{id}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	query := `DELETE FROM products WHERE id = ?`
	result, err := config.DB.Exec(query, id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		utils.ErrorResponse(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.SuccessResponse(w, "Product deleted successfully", nil)
}

// SearchProducts - GET /api/products/search?q=keyword
func SearchProducts(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("q")
	if keyword == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Search keyword is required")
		return
	}

	query := `SELECT id, name, price, stock, category, rating, created_at 
              FROM products 
              WHERE name LIKE ? OR category LIKE ?
              ORDER BY rating DESC`

	searchPattern := "%" + keyword + "%"
	rows, err := config.DB.Query(query, searchPattern, searchPattern)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to search products")
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Stock,
			&product.Category, &product.Rating, &product.CreatedAt)
		if err != nil {
			continue
		}
		products = append(products, product)
	}

	utils.SuccessResponse(w, "Search completed", products)
}

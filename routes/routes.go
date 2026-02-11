package routes

import (
	"net/http"

	"github.com/HHHAAAANNNNN/go-commerce-backend/controllers"
	"github.com/HHHAAAANNNNN/go-commerce-backend/middlewares"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    // Apply global middlewares
    router.Use(middlewares.Logger)
    router.Use(middlewares.CORS)

    // API routes
    api := router.PathPrefix("/api").Subrouter()

    // Health check
    api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"status":"OK","message":"Server is running"}`))
    }).Methods("GET")

    // User routes
    api.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
    api.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
    api.HandleFunc("/users", controllers.CreateUser).Methods("POST")
    api.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
    api.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

    // Product routes
    api.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
    api.HandleFunc("/products/search", controllers.SearchProducts).Methods("GET")
    api.HandleFunc("/products/{id}", controllers.GetProductByID).Methods("GET")
    api.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
    api.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
    api.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")

    return router
}
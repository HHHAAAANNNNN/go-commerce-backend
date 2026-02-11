package models

import "time"

type Product struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Price     int       `json:"price"`
    Stock     int       `json:"stock"`
    Category  string    `json:"category"`
    Rating    float64   `json:"rating"`
    CreatedAt time.Time `json:"created_at"`
}

type ProductCreateRequest struct {
    ID       string  `json:"id"`
    Name     string  `json:"name"`
    Price    int     `json:"price"`
    Stock    int     `json:"stock"`
    Category string  `json:"category"`
    Rating   float64 `json:"rating"`
}

type ProductUpdateRequest struct {
    Name     string  `json:"name,omitempty"`
    Price    int     `json:"price,omitempty"`
    Stock    int     `json:"stock,omitempty"`
    Category string  `json:"category,omitempty"`
    Rating   float64 `json:"rating,omitempty"`
}
package utils

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Success bool        `json:"success"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

// JSONResponse - Send JSON response
func JSONResponse(w http.ResponseWriter, statusCode int, response Response) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(response)
}

// SuccessResponse - Send success response
func SuccessResponse(w http.ResponseWriter, message string, data interface{}) {
    JSONResponse(w, http.StatusOK, Response{
        Success: true,
        Message: message,
        Data:    data,
    })
}

// ErrorResponse - Send error response
func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
    JSONResponse(w, statusCode, Response{
        Success: false,
        Error:   message,
    })
}

// CreatedResponse - Send 201 Created response
func CreatedResponse(w http.ResponseWriter, message string, data interface{}) {
    JSONResponse(w, http.StatusCreated, Response{
        Success: true,
        Message: message,
        Data:    data,
    })
}
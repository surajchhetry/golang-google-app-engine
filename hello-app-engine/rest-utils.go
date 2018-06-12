package main

import (
	"encoding/json"
	"net/http"
)

const SUCCESS_MESSAGE = "SUCCESS"
const ERROR_MESSAGE = "ERROR"

type DefaultResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(DefaultResponse{Code: 0, Message: SUCCESS_MESSAGE})
}

func OkWithMessage(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(DefaultResponse{Code: 0, Message: message})
}

func OkWithData(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(DefaultResponse{Code: 0, Message: SUCCESS_MESSAGE, Data: data})
}

func Error(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(DefaultResponse{Code: http.StatusBadRequest, Message: ERROR_MESSAGE})
	w.WriteHeader(http.StatusBadRequest)
}

func ErrorWithMessage(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(DefaultResponse{Code: http.StatusBadRequest, Message: message})
	w.WriteHeader(http.StatusBadRequest)
}

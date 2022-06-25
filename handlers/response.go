package handlers

import (
	"encoding/json"
	"net/http"
)

type ResponseInterface interface {
	Send(w http.ResponseWriter)
}

// @Description Структура HTTP ответа метода POST /v1/auth/register
type InsertResponse struct {
	ID string `json:"id" example:"123"`
}

func (p InsertResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&p)
}

// @Description Структура HTTP ответа ERROR
type ErrResponse struct {
	HTTPStatusCode int    `json:"status_code,omitempty"`
	ErrorText      string `json:"error,omitempty"`
}

func (err *ErrResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HTTPStatusCode)
	json.NewEncoder(w).Encode(&err)
}

type UserResponse struct {
	Phone string `json:"phone" example:"+79167003020"`
	Login string `json:"login" example:"rubella19"`
	Name  string `json:"name" example:"Анастасия"`
	Birth string `json:"birth" example:"2000-07-28"`
	Tag   string `json:"tg" example:"@Rubella19"`
	Email string `json:"email" example:"anastasia.a.krasnova@gmail.com"`
}

func NewUserResponse() *UserResponse {
	return &UserResponse{}
}

func (ur *UserResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&ur)
}

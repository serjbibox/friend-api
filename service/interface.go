package service

import (
	"net/http"
)

type Service interface {
	Get() *UserService
	Create(r *http.Request) *UserService
	Modify(r *http.Request) *UserService
	Init(r *http.Request) *UserService
}

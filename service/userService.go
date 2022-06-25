package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/serjbibox/friend-api/model"
)

const (
	REGISTER_MODE = iota
	LOGIN_MODE
	LOGOUT_MODE
)

type UserService struct {
	User      *model.User
	Err       error
	ErrStatus int
	Mode      int
}

func New() *UserService {
	return &UserService{
		User:      &model.User{},
		Err:       nil,
		ErrStatus: 0,
		Mode:      REGISTER_MODE,
	}
}

func (s *UserService) Get() *UserService {
	return s
}
func (s *UserService) Init(r *http.Request) *UserService {
	if s.Err != nil {
		return s
	}
	switch s.Mode {
	case REGISTER_MODE:
		log.Println("регистрация нового пользователя")
	case LOGIN_MODE:
		log.Println("вход пользователя")
	}
	if err := json.NewDecoder(r.Body).Decode(&s.User); err != nil {
		s.Err = fmt.Errorf("%w", err)
		return s
	}
	return s
}

func (s *UserService) Modify(r *http.Request) *UserService {
	return s
}

func (s *UserService) Create(r *http.Request) *UserService {

	return s
}

package dao

import (
	"github.com/serjbibox/friend-api/service"
)

type UserDAO struct {
	UserService *service.UserService
}

func New(s *service.UserService) *UserDAO {
	return &UserDAO{s}
}

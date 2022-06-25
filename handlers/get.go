package handlers

import (
	"errors"
	"log"
	"net/http"

	"github.com/serjbibox/friend-api/dao"
	"github.com/serjbibox/friend-api/service"
)

// @Summary  Получает запись из таблицы friend_user по ID записи в query parameter
// @Tags         /v1/user
// @Produce      json
// @Param    id   query      int  true  "friend_user PRIMARY KEY ID"
// @Success  200  {object}  handlers.UserResponse
// @Failure  400  {object}  handlers.ErrResponse
// @Failure  503  {object}  handlers.ErrResponse
// @Router   /v1/user [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	s := service.New()
	dao := dao.New(s)
	s.ErrStatus = http.StatusServiceUnavailable
	s.User.ID = r.URL.Query().Get("id")
	if s.User.ID == "" {
		s.Err = errors.New("отсутствует параметр id")
		SendErr(w, s.ErrStatus, s.Err)
		return
	}
	dao.GetUser()
	if dao.UserService.Err != nil {
		log.Println(s.ErrStatus, s.Err)
		SendErr(w, s.ErrStatus, s.Err)
		return
	}
	ur := NewUserResponse()
	ur.Phone = s.User.Phone
	ur.Login = s.User.Login
	ur.Name = s.User.Name
	ur.Birth = s.User.Birth
	ur.Tag = s.User.Tag
	ur.Email = s.User.Email
	SendHttp(w, ur)
}

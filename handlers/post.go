package handlers

import (
	"log"
	"net/http"

	"github.com/serjbibox/friend-api/dao"
	"github.com/serjbibox/friend-api/service"
)

// @Summary  Создаёт новую запись в таблице friend_user
// @Tags     /v1/auth/register
// @Accept   json
// @Produce  json
// @Param    input  body      model.User  true  "регистрация карточки пользователя"
// @Success  200    {object}  handlers.InsertResponse
// @Failure  400    {object}  handlers.ErrResponse
// @Failure  503    {object}  handlers.ErrResponse
// @Router   /v1/auth/register [post]
func NewUser(w http.ResponseWriter, r *http.Request) {
	s := service.New()
	dao := dao.New(s)
	s.ErrStatus = http.StatusServiceUnavailable
	s.Mode = REGISTER_MODE
	dao.UserService.
		Init(r).
		ValidateFields().
		ValidateData()
	dao.Create()
	if dao.UserService.Err != nil {
		log.Println(s.ErrStatus, s.Err)
		SendErr(w, s.ErrStatus, s.Err)
		return
	}
	SendHttp(w,
		InsertResponse{
			ID: s.User.ID,
		})
}

// @Summary  Авторизация пользователя
// @Tags     /v1/auth/login
// @Accept   json
// @Produce  json
// @Param    input  body      model.UserLogin  true  "запрос на авторизацию"
// @Success  200    {object}  handlers.InsertResponse
// @Failure  400    {object}  handlers.ErrResponse
// @Failure  503    {object}  handlers.ErrResponse
// @Router   /v1/auth/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	s := service.New()
	dao := dao.New(s)
	s.ErrStatus = http.StatusServiceUnavailable
	s.Mode = LOGIN_MODE
	dao.UserService.
		Init(r).
		ValidateFields().
		ValidateData()
	dao.UserVerification()
	if dao.UserService.Err != nil {
		log.Println(s.ErrStatus, s.Err)
		SendErr(w, s.ErrStatus, s.Err)
		return
	}
	SendHttp(w,
		InsertResponse{
			ID: s.User.ID,
		})
}

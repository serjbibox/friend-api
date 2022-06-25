package model

// @Description Структура карточки пользователя в теле запроса POST /v1/auth/register
type User struct {
	ID       string `json:"id" example:"11234"`
	Phone    string `json:"phone" example:"+79167003020"`
	Login    string `json:"login" example:"rubella19"`
	Password string `json:"password" example:"1Qwerty!"`
	Name     string `json:"name" example:"Анастасия"`
	Birth    string `json:"birth" example:"2000-07-28"`
	Tag      string `json:"tg" example:"@Rubella19"`
	Email    string `json:"email" example:"anastasia.a.krasnova@gmail.com"`
}

// @Description Структура JSON метода POST /v1/auth/login
type UserLogin struct {
	Login    string `json:"login" example:"rubella19"`
	Password string `json:"password" example:"1Qwerty!"`
}

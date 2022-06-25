package handlers

import (
	"fmt"
	"net/http"
)

const (
	REGISTER_MODE = iota
	LOGIN_MODE
	LOGOUT_MODE
)

func SendHttp(w http.ResponseWriter, v ResponseInterface) {
	v.Send(w)
}

func SendErr(w http.ResponseWriter, statusCode int, err error) {
	er := ErrResponse{
		HTTPStatusCode: statusCode,
		ErrorText:      fmt.Sprint(err),
	}
	er.Send(w)
}

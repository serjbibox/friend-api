package dao

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/serjbibox/friend-api/service"
)

func (ud *UserDAO) GetUser() *service.UserService {
	if ud.UserService.Err != nil {
		return ud.UserService
	}
	conn, err := DbConnect()
	if err != nil {
		ud.UserService.Err = fmt.Errorf("%w", err)
		return ud.UserService
	}
	defer conn.Close(context.Background())

	u := ud.UserService.User
	id, err := strconv.Atoi(u.ID)
	if err != nil {
		ud.UserService.Err = fmt.Errorf("%w", err)
		return ud.UserService
	}
	ls := false
	query := "SELECT login_status FROM " + db.tables["users"] + " WHERE id = $1;"
	err = conn.QueryRow(context.Background(), query, id).Scan(&ls)
	if err != nil {
		ud.UserService.Err = fmt.Errorf("%w", err)
		return ud.UserService
	}
	if !ls {
		ud.UserService.Err = errors.New("необходимо войти в систему")
		return ud.UserService
	}
	query = "SELECT phone, login, name, birth, tg, email FROM " + db.tables["users"] + " WHERE id = $1;"
	err = conn.QueryRow(context.Background(), query, id).
		Scan(&u.Phone, &u.Login, &u.Name, &u.Birth, &u.Tag, &u.Email)
	if err != nil {
		ud.UserService.Err = fmt.Errorf("%w", err)
		return ud.UserService
	}

	return ud.UserService
}

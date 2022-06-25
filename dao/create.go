package dao

import (
	"context"
	"fmt"

	"github.com/serjbibox/friend-api/service"

	"golang.org/x/crypto/bcrypt"
)

func (ud *UserDAO) Create() *service.UserService {
	if ud.UserService.Err != nil {
		return ud.UserService
	}
	conn, err := DbConnect()
	if err != nil {
		ud.UserService.Err = fmt.Errorf("%w", err)
		return ud.UserService
	}
	defer conn.Close(context.Background())
	query := "INSERT INTO " + db.tables["users"] + " (date_added, phone, login, password, name, birth, tg, email) " +
		"VALUES (current_timestamp, $1, $2, $3, $4, $5, $6, $7) RETURNING ID;"
	u := ud.UserService.User
	id := 0
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		ud.UserService.Err = fmt.Errorf("%w", err)
		return ud.UserService
	}
	err = conn.QueryRow(context.Background(), query, u.Phone, u.Login, hashedPassword, u.Name, u.Birth, u.Tag, u.Email).Scan(&id)
	if err != nil {
		ud.UserService.Err = fmt.Errorf("%w", err)
		return ud.UserService
	}
	u.ID = fmt.Sprint(id)
	return ud.UserService
}

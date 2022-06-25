package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/serjbibox/friend-api/service"

	"golang.org/x/crypto/bcrypt"
)

func (ud *UserDAO) UserVerification() *service.UserService {
	if ud.UserService.Err != nil {
		return ud.UserService
	}
	conn, err := DbConnect()
	if err != nil {
		ud.UserService.Err = fmt.Errorf("%w", err)
		return ud.UserService
	}
	defer conn.Close(context.Background())
	hashedPassword := ""
	id := 0
	u := ud.UserService.User
	query := "SELECT password, id FROM " + db.tables["users"] + " WHERE login = $1 AND id = (SELECT MAX(id) from " + db.tables["users"] + ");"
	log.Println("чтение по логину", ud.UserService.User.Login)
	err = conn.QueryRow(context.Background(), query, ud.UserService.User.Login).Scan(&hashedPassword, &id)
	if err != nil {
		ud.UserService.Err = fmt.Errorf("%w", err)
		return ud.UserService
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password))
	if err != nil {
		ud.UserService.Err = fmt.Errorf("%w", err)
		return ud.UserService
	}

	query = "UPDATE " + db.tables["users"] + " SET login_status = ($1) WHERE id = $2 RETURNING ID;"
	err = conn.QueryRow(context.Background(), query, true, id).Scan(&id)
	if err != nil {
		ud.UserService.Err = fmt.Errorf("%w", err)
		return ud.UserService
	}
	u.ID = fmt.Sprint(id)
	log.Println("вход выполнен, id =", id)
	return ud.UserService
}

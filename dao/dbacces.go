package dao

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func DbConnect() (*pgx.Conn, error) {
	var err error
	if DbConnString == "" {
		if DbConnString, err = ConnStringConfig(); err != nil {
			return nil, err
		}
	}
	conn, err := pgx.Connect(context.Background(), DbConnString)
	if err != nil {
		return nil, err
	}
	log.Println(DbConnString)
	log.Println("Установлено соединение с БД")
	return conn, nil
}

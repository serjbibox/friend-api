package dao

import (
	"database/sql"
	"log"
	"os"

	"github.com/spf13/viper"
)

const (
	FRIEND_DB_PASS = "FRIEND_DB_PASS"
)

type PostgresSettings struct {
	tables map[string]string
}

func DbPrepare() {
	db = &PostgresSettings{tables: map[string]string{}}
	db.tables["users"] = viper.GetString("db_tables.users_table")
}

var db *PostgresSettings
var DB *sql.DB
var DbConnString string

func ConnStringConfig() (string, error) {
	connString := "postgres://" +
		viper.GetString("db.username") +
		":" +
		readEnv(FRIEND_DB_PASS) +
		"@" +
		viper.GetString("db.host") +
		":" +
		viper.GetString("db.port") +
		"/" +
		viper.GetString("db.name") +
		"?sslmode=" +
		viper.GetString("db.sslmode")

	log.Println(connString)

	return connString, nil

}

func readEnv(key string) string {
	if env, ok := os.LookupEnv(key); !ok {
		return ""
	} else {
		return env
	}
}

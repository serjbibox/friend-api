package main

import (
	"log"

	_ "golang.org/x/crypto/bcrypt"

	_ "github.com/jackc/pgtype"
	_ "github.com/jackc/pgx/v4"
	"github.com/serjbibox/friend-api/dao"
	_ "github.com/serjbibox/friend-api/docs"
	_ "github.com/serjbibox/friend-api/handlers"
	"github.com/serjbibox/friend-api/server"
	"github.com/spf13/viper"
)

// @title          API для платформы ДРУГ
// @version        1.0
// @description    API для взаимодействия с сервером платформы ДРУГ.
// @contact.name   API Support
// @contact.email  serj_bibox@mail.ru
// @BasePath
func main() {
	err := initConfig()
	if err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	dao.DbPrepare()
	log.Panic(server.New(viper.GetString("port")).Run())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

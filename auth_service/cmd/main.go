package main

import (
	"auth_service"
	"auth_service/pkg/handler"
	"auth_service/pkg/repository"
	"auth_service/pkg/service"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	errInitConfig := initConfig()

	if errInitConfig != nil {
		logrus.Fatalf("Error load config file: %s", errInitConfig.Error())
	}

	db, errDB := repository.NewSQLDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})

	if errDB != nil {
		logrus.Fatalf("Error init DB: %s", errDB.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(auth_service.Server)
	err := srv.Run(viper.GetString("port"), handlers.InitRoutes())

	if err != nil {
		logrus.Fatalf("Error start HTTP server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("/home/andrey/go/auth_service/configs/")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

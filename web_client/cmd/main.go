package main

import (
	"web_client"
	"web_client/pkg/handler"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	err_init_config := initConfig()

	if err_init_config != nil {
		logrus.Fatalf("Ошибка загрузки файла конфигурации сервера: %s", err_init_config.Error())
	}

	handlers := handler.NewHandler()

	srv := new(web_client.Server)
	err := srv.Run(viper.GetString("port"), handlers.InitRoutes())

	if err != nil {
		logrus.Fatalf("Ошибка во время запуска http сервера: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("/home/andrey/go/web_client/configs/")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

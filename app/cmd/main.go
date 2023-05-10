package main

import (
	"log"

	"github.com/FarrukhMahkamov/teamly_career/internal/repository"
	"github.com/FarrukhMahkamov/teamly_career/internal/service"
	handler "github.com/FarrukhMahkamov/teamly_career/internal/transport/http"
	"github.com/FarrukhMahkamov/teamly_career/pkg"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	server := new(pkg.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

package main

import (
	"log"

	"github.com/FarrukhMahkamov/teamly_career/internal/repository"
	"github.com/FarrukhMahkamov/teamly_career/internal/service"
	handler "github.com/FarrukhMahkamov/teamly_career/internal/transport/http"
	"github.com/FarrukhMahkamov/teamly_career/pkg"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := pkg.NewPostgresDB(pkg.PostgresConfig{
		Host:     viper.GetString("DBHOST"),
		Port:     viper.GetString("DBPORT"),
		Username: viper.GetString("DBUSERNAME"),
		DBName:   viper.GetString("DBNAME"),
		Password: viper.GetString("DBPASSWORD"),
		SSLMode:  viper.GetString("SSLMODE"),
	})

	if err != nil {
		log.Fatalf("error initializing db: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
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

// migrate create -ext sql -dir ./schema/migrations -seq init

package main

import (
	"os"

	"github.com/FarrukhMahkamov/teamly_career/internal/repository"
	"github.com/FarrukhMahkamov/teamly_career/internal/service"
	handler "github.com/FarrukhMahkamov/teamly_career/internal/transport/http"
	"github.com/FarrukhMahkamov/teamly_career/pkg"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	//set formatter for logrus
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//set config file
	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	//set env variables
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//set db connection
	db, err := pkg.NewPostgresDB(pkg.PostgresConfig{
		Host:     viper.GetString("DBHOST"),
		Port:     viper.GetString("DBPORT"),
		Username: viper.GetString("DBUSERNAME"),
		DBName:   viper.GetString("DBNAME"),
		Password: os.Getenv("DBPASSWORD"),
		SSLMode:  viper.GetString("SSLMODE"),
	})

	//check if db connection is successful
	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}

	//set repositories, services and handlers
	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	server := new(pkg.Server)

	//run server
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

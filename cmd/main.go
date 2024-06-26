package main

import (
	"os"

	app "github.com/ShawaDev/Dealer"
	"github.com/ShawaDev/Dealer/pkg/handler"
	"github.com/ShawaDev/Dealer/pkg/repository"
	"github.com/ShawaDev/Dealer/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	//Setting logrus formater
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//Enabling env files
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}

	//Enabling config file
	if err := InitConfig(); err != nil {
		logrus.Fatal(err)
	}

	//Connecting to DB
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatal(err)
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	//Creating server
	server := new(app.Server)

	//Running server
	if err := server.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}

}

// Reading config.yaml file
func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

package settings

import (
	"fmt"
	"os"
	"strconv"
)

type Settings struct {
	Application
	Postgres
}

// TOOD: implement loading the environment variables
func Load() (Settings, error) {
	db_port, err1 := strconv.Atoi(os.Getenv("DB_PORT"))
	app_port, err2 := strconv.Atoi(os.Getenv("APP_PORT"))
	if err1 != nil && err2 != nil {
		fmt.Errorf("No DB_PORT and No App_PORT")
	}
	return Settings{
		Postgres: Postgres{
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_Name"),
			Port:     db_port,
			Password: os.Getenv("DB_PASSWORD"),
			User:     os.Getenv("DB_PASSWORD"),
		},
		Application: Application{
			Port: app_port,
		},
	}, nil
}

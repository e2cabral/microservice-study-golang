package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"microservice-golang/src/domain/entities"
	"microservice-golang/src/infra/helpers"
)

func Connect() (*gorm.DB, error) {
	env := helpers.Env{}
	host, err := env.GetVariable("DB_HOST")
	user, err := env.GetVariable("DB_USER")
	password, err := env.GetVariable("DB_PASSWORD")
	dbname, err := env.GetVariable("DB_NAME")
	port, err := env.GetVariable("DB_PORT")

	if err != nil {
		return nil, err
	}

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		host, user, password, dbname, port,
	)

	connection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	println("Database connected.")

	return connection, nil
}

func Migrate(database *gorm.DB) error {
	err := database.AutoMigrate(
		&entities.Product{},
		&entities.Review{},
		&entities.Category{},
	)
	if err != nil {
		return err
	}

	println("Tables have been migrated.")

	return nil
}

func InitializeDatabaseConnection() error {
	database, err := Connect()
	if err != nil {
		return err
	}

	if err := Migrate(database); err != nil {
		return err
	}

	return nil
}

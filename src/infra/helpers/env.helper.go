package helpers

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

type Env struct {
}

func (e *Env) loadEnvFile() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	return nil
}

func (e *Env) GetVariable(name string) (interface{}, error) {
	if err := e.loadEnvFile(); err != nil {
		return nil, err
	}

	variable := os.Getenv(name)

	if len(strings.TrimSpace(variable)) != 0 {
		return variable, nil
	}

	return nil, errors.New("environment variable not found")
}

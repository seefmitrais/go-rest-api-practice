package config

import (
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) {
	err := godotenv.Load(filenames...)
	if err != nil {
		panic(err)
	}
}

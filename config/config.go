package config

import (
	"equal_dark_crawler/utils"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Scylla struct {
		Host string
		Port int
	}
}

// Values : config values
var Values config

// Init : read .env file
func Init(filenames ...string) {
	err := godotenv.Load(filenames...)
	if err != nil {
		panic(err)
	}

	Values.Scylla.Host = os.Getenv("SCYLLA_HOST")
	Values.Scylla.Port = utils.ParseInt(os.Getenv("SCYLLA_PORT"))
}

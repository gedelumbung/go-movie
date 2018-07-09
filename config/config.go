package conf

import (
	"os"

	"github.com/kelseyhightower/envconfig"

	"github.com/joho/godotenv"
)

type SiteConfiguration struct {
	Env string `envconfig:"ENV" default:"local"`
}

type LogConfiguration struct {
	Dir string `envconfig:"DIR" default:"./storage/log"`
}

type Configuration struct {
	SITE SiteConfiguration
	API  struct {
		Host string `envconfig:"HOST" default:":5000"`
	}
	DB struct {
		Driver string `envconfig:"DRIVER" required:"true"`
		Mysql  struct {
			URL string `envconfig:"URL" required:"true"`
		}
		Sqlite3 struct {
			URL string `envconfig:"URL" required:"true"`
		}
	}
	LOG LogConfiguration
}

func loadEnvironment(filename string) error {
	var err error
	if filename != "" {
		err = godotenv.Load(filename)
	} else {
		err = godotenv.Load()
		if os.IsNotExist(err) {
			return nil
		}
	}
	return err
}

func LoadConfig(filename string) (*Configuration, error) {
	if err := loadEnvironment(filename); err != nil {
		return nil, err
	}

	config := new(Configuration)
	if err := envconfig.Process("GOMOVIE", config); err != nil {
		return nil, err
	}
	return config, nil
}

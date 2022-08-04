package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

const (
	goEnv       = "GO_ENV"
	port        = "PORT"
	mongoDbUri  = "MONGO_DB_URI"
	mongoDbName = "MONGO_DB_NAME"
)

type (
	IConfigService interface {
		GetMongoDbUri() string
	}
	Config struct {
		GoEnv       string
		Port        string
		MongoDbUri  string
		MongoDbName string
	}
)

var defaults = map[string]string{
	goEnv:       "",
	port:        "8000",
	mongoDbUri:  "",
	mongoDbName: "",
}

func (cs Config) GetMongoDbUri() string { return cs.MongoDbUri }

func LoadConfig() (Config, error) {

	for k, v := range defaults {
		viper.SetDefault(k, v)
	}

	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Could not read config file: %v", err)
	}

	config := Config{
		GoEnv:       viper.GetString(goEnv),
		Port:        viper.GetString(port),
		MongoDbUri:  viper.GetString(mongoDbUri),
		MongoDbName: viper.GetString(mongoDbName),
	}

	fmt.Println(config)
	return config, nil

}

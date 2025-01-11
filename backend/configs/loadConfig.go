package configs

import (
	"encoding/json"

	"os"
)

type Config struct {
	Port  int    `json:"port"`
	Dsn   string `json:"dsn"`
	Debug bool   `json:"debug"`
}

func LoadConfig(configfile string) Config {

	config := Config{}
	text, err := os.ReadFile(configfile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(text, &config)
	if err != nil {
		panic(err)
	}

	return config
}

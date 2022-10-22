package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Port int    `json:"port"`
	Dsn  string `json:"dsn"`
}

func LoadConfig() Config {
	fmt.Println("config")
	fmt.Print("\n\n")
	config := Config{}
	text, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(text, &config)
	if err != nil {
		panic(err)
	}

	return config
}

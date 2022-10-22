package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func LoadConfig() string {
	fmt.Println("config")
	fmt.Print("\n\n")
	config := map[string]string{}
	text, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(text, &config)
	if err != nil {
		panic(err)
	}
	dsn, ok := config["dsn"]
	if !ok {
		panic(errors.New("config dsn not found"))
	}
	println(dsn)
	return dsn
}

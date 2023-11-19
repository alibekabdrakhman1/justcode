package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile("lecture11/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading yaml file: %v", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("error decoding into config struct: %v", err)
	}
	fmt.Println(config.HttpServer)
	fmt.Println(config.Postgres)
}

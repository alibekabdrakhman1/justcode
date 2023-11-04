package main

import (
	"fmt"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/applicator"
	"github.com/alibekabdrakhman/justcode/lecture12/internal/auth/config"
	"github.com/spf13/viper"
	"log"
)

func main() {
	cfg, err := loadConfig("lecture12/config/user")
	if err != nil {
		log.Fatalf("failed to load config err: %v", err)
	}

	app := applicator.New(&cfg)
	log.Fatal(app.Run())
}
func loadConfig(path string) (config config.Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("failed to ReadInConfig err: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to Unmarshal config err: %w", err)
	}

	return config, nil
}

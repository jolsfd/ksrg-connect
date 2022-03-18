package main

import (
	"fmt"
	"log"

	"github.com/jolsfd/ksrg-connect/cmd"
	"github.com/spf13/viper"
)

func main() {
	// Setup config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./public")

	// Read config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("no config file", err)
		} else {
			log.Fatal(err)
		}
	}

	// Unmarshall config
	err := viper.Unmarshal(&cmd.AppConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Debug info
	fmt.Printf("Config: %v\n", cmd.AppConfig)

	// Start server
	cmd.StartServer()
}

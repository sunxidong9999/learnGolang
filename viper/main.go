package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func readConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	fmt.Println("Config file used:", viper.ConfigFileUsed())
	fmt.Println("Server Address:", viper.GetString("server.addr"))
	fmt.Println("Server Port:", viper.GetInt("server.port"))
	fmt.Println("Syslog Address:", viper.GetString("syslog.addr"))
	fmt.Println("Syslog Port:", viper.GetInt("syslog.port"))
	fmt.Println("Use Viper:", viper.GetBool("useViper"))
	fmt.Println("Price:", viper.GetFloat64("basicPrice"))

	err = viper.WriteConfigAs("new_config.json") // Save the current config to a new file
	if err != nil {
		log.Fatalf("Error writing config file: %s \n", err)
	} else {
		fmt.Println("New config file saved as new_config.json")
	}
}

func main() {
	readConfig()
}

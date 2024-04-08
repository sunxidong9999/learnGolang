package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server struct {
	Addr string `json:"addr"`
	Port int32  `json:"port"`
}

/*
type Syslog struct {
	Addr string
	Port int32
}
*/

type Config struct {
	Server Server `json:"server"`
	Syslog Server `json:"syslog"`
}

func LoadConfig(filename string, conf *Config) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Read file failed: %v\n", err)
		return err
	}

	err = json.Unmarshal(data, conf)
	if err != nil {
		fmt.Printf("json parse failed: %v\n", err)
		return err
	}

	return nil
}

func SaveToJson(conf Config, f string) error {
	data, err := json.MarshalIndent(conf, "", "    ")
	if err != nil {
		fmt.Printf("json marshal failed: %v\n", err)
		return err
	}

	err = os.WriteFile(f, data, 0644)
	if err != nil {
		fmt.Printf("Write file failed: %v\n", err)
		return err
	}

	return nil
}

func main() {
	var conf Config

	err := LoadConfig("./config.json", &conf)
	if err != nil {
		fmt.Println("Load config failed: ", err)
	}
	fmt.Println("config:", conf)
	fmt.Println(conf.Server.Addr)
	fmt.Println(conf.Server.Port)
	fmt.Println(conf.Syslog.Addr)
	fmt.Println(conf.Syslog.Port)

	fmt.Println("Save to new config file: conf.json:")
	err = SaveToJson(conf, "conf.json")
	if err != nil {
		fmt.Println("Save to json file failed:", err)
	}
}

package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Instance *Config

type Config struct {
	AppName string `yaml:"AppName"`
	Mysql   struct {
		Host string `yaml:"Host"`
		Port string `yaml:"Port"`
		User string `yaml:"User"`
		Pwd  string `yaml:"Pwd"`
	} `yaml:"Mysql"`
	MongoDB struct {
		Host string `yaml:"Host"`
		Port string `yaml:"Port"`
		User string `yaml:"User"`
		Pwd  string `yaml:"Pwd"`
	} `yaml:"MongoDB"`
}

func Init(filename string) *Config {
	Instance = &Config{}

	config := viper.New()
	config.SetConfigName("app")    // name of config file (without extension)
	config.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	config.AddConfigPath(filename) // optionally look for config in the working directory
	err := config.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	if err := config.Unmarshal(Instance); err != nil {
		fmt.Println(err)
	}
	fmt.Println(Instance.Mysql.Host)

	config.WatchConfig()
	config.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed", e.Name)
	})

	return Instance
}

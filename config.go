package vodka

import (
	"fmt"

	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigType("toml")   // or viper.SetConfigType("YAML")
	viper.SetConfigName("config") // name of config file (without extension)

	viper.AddConfigPath("/etc/tadinefis/") // path to look for the config file in
	viper.AddConfigPath("$HOME/.tadinefis")
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

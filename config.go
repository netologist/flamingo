package flamingo

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

func loadConfig() {
	viper.SetConfigType("toml")   // or viper.SetConfigType("YAML")
	viper.SetConfigName("config") // name of config file (without extension)

	viper.AddConfigPath("./config/") // optionally look for config in the working directory
	err := viper.ReadInConfig()      // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	loggerSettings()
}

func loggerSettings() {
	format := viper.GetString("log.format")
	level := viper.GetString("log.level")
	output := viper.GetString("log.output")

	if format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	if output == "stderr" {
		logrus.SetOutput(os.Stderr)
	} else {
		logrus.SetOutput(os.Stdout)
	}

	if level == "panic" {
		logrus.SetLevel(logrus.PanicLevel)
	} else if level == "error" {
		logrus.SetLevel(logrus.ErrorLevel)
	} else if level == "warning" {
		logrus.SetLevel(logrus.WarnLevel)
	} else if level == "info" {
		logrus.SetLevel(logrus.InfoLevel)
	} else if level == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.ErrorLevel)
	}

	fmt.Println("LogLevel:", logrus.GetLevel())
	fmt.Println("LogOut:", output)
	fmt.Println("LogFormat:", format)
}

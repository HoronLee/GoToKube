package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	Data Config
)

func InitConfig() {
	configPath := "config.toml"
	viper.SetConfigFile(configPath)
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	setDefaults()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Printf("Configuration file does not exist, automatically create a default configuration file")
		createDefaultConfig(configPath)
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading configuration file" + err.Error())
	} else {
		log.Printf("Read configuration file")
	}

	if err := viper.Unmarshal(&Data); err != nil {
		log.Printf("Unable to decode into struct" + err.Error())
	}
	fmt.Println(Data)
}

func setDefaults() {
	defaults := map[string]interface{}{
		"Web.Enable":            false,
		"Web.ListeningAddr":     ":8080",
		"Kubernetes.Enable":     false,
		"Kubernetes.ConfigPath": "",
		"Database.Type":         "sqlite",
		"Database.Path":         "data.db",
		"Database.Addr":         "",
		"Database.User":         "",
		"Database.Password":     "",
		"Database.Name":         "",
		"Common.LogDir":         "",
		"Common.TermEnable":     false,
		"Auth.User":             "root",
		"Auth.Pass":             "123456",
	}
	for key, value := range defaults {
		viper.SetDefault(key, value)
	}
}

func createDefaultConfig(path string) {
	viper.Set("Web.Enable", viper.GetBool("Web.Enable"))
	viper.Set("Web.ListeningAddr", viper.GetString("Web.ListeningAddr"))
	viper.Set("Kubernetes.Enable", viper.GetBool("Kubernetes.Enable"))
	viper.Set("Kubernetes.ConfigPath", viper.GetString("Kubernetes.ConfigPath"))
	viper.Set("Database.Type", viper.GetString("Database.Type"))
	viper.Set("Database.Path", viper.GetString("Database.Path"))
	viper.Set("Database.Addr", viper.GetString("Database.Addr"))
	viper.Set("Database.User", viper.GetString("Database.User"))
	viper.Set("Database.Password", viper.GetString("Database.Password"))
	viper.Set("Database.Name", viper.GetString("Database.Name"))
	viper.Set("Common.LogDir", viper.GetString("Common.LogDir"))
	viper.Set("Common.TermEnable", viper.GetBool("Common.TermEnable"))
	viper.Set("Auth.User", viper.GetString("Auth.User"))
	viper.Set("Auth.Pass", viper.GetString("Auth.Pass"))

	if err := viper.WriteConfigAs(path); err != nil {
		log.Println("Error writing default configuration file:", err)
	} else {
		log.Println("Default configuration file created successfully.")
	}
}

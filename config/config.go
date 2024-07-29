package config

import (
	"GoToKube/logger"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type WebConfig struct {
	Enable        bool   `mapstructure:"Enable"`
	ListeningAddr string `mapstructure:"ListeningAddr"`
}

type KubeConfig struct {
	Enable     bool   `mapstructure:"Enable"`
	ConfigPath string `mapstructure:"ConfigPath"`
}

type DatabaseConfig struct {
	Type     string `mapstructure:"Type"`
	Path     string `mapstructure:"Path"`
	Addr     string `mapstructure:"Addr"`
	User     string `mapstructure:"User"`
	Password string `mapstructure:"Password"`
	Name     string `mapstructure:"Name"`
}

type LogConfig struct {
	Dir string `mapstructure:"Dir"`
}

type AuthConfig struct {
	Pass string `mapstructure:"Pass"`
}

type Config struct {
	Web      WebConfig      `mapstructure:"Web"`
	Kube     KubeConfig     `mapstructure:"Kube"`
	Database DatabaseConfig `mapstructure:"Database"`
	Log      LogConfig      `mapstructure:"Log"`
	Auth     AuthConfig     `mapstructure:"Auth"`
}

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
		logger.GlobalLogger.Error("Configuration file does not exist, automatically create a default configuration file")
		createDefaultConfig(configPath)
	}

	if err := viper.ReadInConfig(); err != nil {
		logger.GlobalLogger.Error("Error reading configuration file" + err.Error())
	} else {
		logger.GlobalLogger.Info("Read configuration file")
	}

	if err := viper.Unmarshal(&Data); err != nil {
		logger.GlobalLogger.Error("Unable to decode into struct" + err.Error())
	}
	fmt.Println(Data)
}

func setDefaults() {
	viper.SetDefault("Web.Enable", false)
	viper.SetDefault("Web.ListeningAddr", ":8080")
	viper.SetDefault("Term.Enable", false)
	viper.SetDefault("Kube.Enable", false)
	viper.SetDefault("Kube.ConfigPath", "")
	viper.SetDefault("Database.Type", "sqlite")
	viper.SetDefault("Database.Path", "data.db")
	viper.SetDefault("Database.Addr", "")
	viper.SetDefault("Database.User", "")
	viper.SetDefault("Database.Password", "")
	viper.SetDefault("Database.Name", "")
	viper.SetDefault("Log.Dir", "")
	viper.SetDefault("Auth.Pass", "gotokube")
}

func createDefaultConfig(path string) {
	config := Config{
		Web: WebConfig{
			Enable:        viper.GetBool("Web.Enable"),
			ListeningAddr: viper.GetString("Web.ListeningAddr"),
		},
		Kube: KubeConfig{
			Enable:     viper.GetBool("Kube.Enable"),
			ConfigPath: viper.GetString("Kube.ConfigPath"),
		},
		Database: DatabaseConfig{
			Type:     viper.GetString("Database.Type"),
			Path:     viper.GetString("Database.Path"),
			Addr:     viper.GetString("Database.Addr"),
			User:     viper.GetString("Database.User"),
			Password: viper.GetString("Database.Password"),
			Name:     viper.GetString("Database.Name"),
		},
		Log: LogConfig{
			Dir: viper.GetString("Log.Dir"),
		},
		Auth: AuthConfig{
			Pass: viper.GetString("Auth.Pass"),
		},
	}

	viper.Set("Web.Enable", config.Web.Enable)
	viper.Set("Web.ListeningAddr", config.Web.ListeningAddr)
	viper.Set("Kube.Enable", config.Kube.Enable)
	viper.Set("Kube.ConfigPath", config.Kube.ConfigPath)
	viper.Set("Database.Type", config.Database.Type)
	viper.Set("Database.Path", config.Database.Path)
	viper.Set("Database.Addr", config.Database.Addr)
	viper.Set("Database.User", config.Database.User)
	viper.Set("Database.Password", config.Database.Password)
	viper.Set("Database.Name", config.Database.Name)
	viper.Set("Log.Dir", config.Log.Dir)
	viper.Set("Auth.Pass", config.Auth.Pass)

	err := viper.WriteConfigAs(path)
	if err != nil {
		logger.GlobalLogger.Error("Error writing default configuration file" + err.Error())
	} else {
		logger.GlobalLogger.Info("Default configuration file created successfully")
	}
}

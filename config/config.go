package config

import (
	"GoToKube/logger"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	WebEnable      bool   `toml:"WebEnable"`
	ListeningAddr  string `toml:"ListeningPort"`
	TermEnable     bool   `toml:"TermEnable"`
	KubeEnable     bool   `toml:"KubeEnable"`
	KubeConfigPath string `toml:"KubeConfigPath"`
	DBType         string `toml:"DBType"`
	DBPath         string `toml:"DBPath"`
	DBAddr         string `toml:"DBAddr"`
	DBUser         string `toml:"DBUser"`
	DBPass         string `toml:"DBPassword"`
	DBName         string `toml:"DBName"`
}

var (
	Data Config
)

func InitConfig() {
	configPath := "config.toml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logger.GlobalLogger.Error("Configuration file does not exist, automatically create a default configuration file")
		createDefaultConfig(configPath)
	}
	logger.GlobalLogger.Info("Read configuration file")
	Data = readConfig(configPath)
}

// 默认配置文件
func createDefaultConfig(path string) {
	config := Config{
		WebEnable:      false,
		ListeningAddr:  ":8080",
		TermEnable:     false,
		KubeEnable:     false,
		KubeConfigPath: "",
		DBType:         "sqlite",
		DBPath:         "data.db",
		DBAddr:         "",
		DBUser:         "",
		DBPass:         "",
		DBName:         "",
	}
	// 写入配置
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	// 编码为 toml 格式
	if err := toml.NewEncoder(file).Encode(config); err != nil {
		panic(err)
	} else {
		logger.GlobalLogger.Info("Configuration file created successfully")
	}

}

// 读取配置文件
func readConfig(path string) Config {
	var config Config
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	// 解码为 Config 结构体类型
	if err := toml.NewDecoder(file).Decode(&config); err != nil {
		panic(err)
	} else {
		logger.GlobalLogger.Info("Configuration file read successfully")
	}
	return config
}

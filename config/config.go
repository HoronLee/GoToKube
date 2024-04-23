package config

import (
	"VDController/logger"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	WebEnable bool `toml:"WebEnable"`
	ListeningAddr string `toml:"ListeningPort"`
}

var (
	ConfigData Config
	cLogger    *logger.Logger
)

func init() {
	cLogger = logger.NewLogger(logger.INFO)
	configPath := "config.toml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		cLogger.Log(logger.INFO, "配置文件不存在，自动创建默认配置文件")
		createDefaultConfig(configPath)
	}
	cLogger.Log(logger.INFO, "读取配置文件")
	ConfigData = readConfig(configPath)
}

func createDefaultConfig(path string) {
	// 默认配置文件
	config := Config{
		WebEnable: false,
		ListeningAddr: "127.0.0.0:8080",

	}
	// 写入配置
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 编码为 toml 格式
	if err := toml.NewEncoder(file).Encode(config); err != nil {
		panic(err)
	} else {
		cLogger.Log(logger.INFO, "配置文件创建成功")
	}

}

func readConfig(path string) Config {
	// 读取配置文件
	var config Config
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 解码为 Config 结构体类型
	if err := toml.NewDecoder(file).Decode(&config); err != nil {
		panic(err)
	} else {
		cLogger.Log(logger.INFO, "配置文件读取成功")
	}
	return config
}

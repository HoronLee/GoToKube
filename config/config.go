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
	LogDir         string `toml:"LogDir"`
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
		WebEnable:      getEnvBool("WEB_ENABLE", false),
		ListeningAddr:  getEnv("LISTENING_PORT", ":8080"),
		TermEnable:     getEnvBool("TERM_ENABLE", false),
		KubeEnable:     getEnvBool("KUBE_ENABLE", false),
		KubeConfigPath: getEnv("KUBE_CONFIG_PATH", ""),
		DBType:         getEnv("DB_TYPE", "sqlite"),
		DBPath:         getEnv("DB_PATH", "data.db"),
		DBAddr:         getEnv("DB_ADDR", ""),
		DBUser:         getEnv("DB_USER", ""),
		DBPass:         getEnv("DB_PASSWORD", ""),
		DBName:         getEnv("DB_NAME", ""),
		LogDir:         getEnv("LOG_DIR", ""),
	}
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	// 编码为 toml 格式
	if err := toml.NewEncoder(file).Encode(config); err != nil {
		panic(err)
	} else {
		logger.GlobalLogger.Info("Configuration file created successfully")
	}
	file.Close()
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

// 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// 获取布尔环境变量，如果不存在则返回默认值
func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value == "true"
}

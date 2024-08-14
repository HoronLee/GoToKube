package config

type Config struct {
	Web        `mapstructure:"Web"`
	Kubernetes `mapstructure:"Kubernetes"`
	Database   `mapstructure:"Database"`
	Common     `mapstructure:"Common"`
	Auth       `mapstructure:"Auth"`
}

type Web struct {
	Enable        bool   `mapstructure:"Enable"`
	ListeningAddr string `mapstructure:"ListeningAddr"`
}

type Kubernetes struct {
	Enable     bool   `mapstructure:"Enable"`
	ConfigPath string `mapstructure:"ConfigPath"`
}

type Database struct {
	Type     string `mapstructure:"Type"`
	Path     string `mapstructure:"Path"`
	Addr     string `mapstructure:"Addr"`
	User     string `mapstructure:"User"`
	Password string `mapstructure:"Password"`
	Name     string `mapstructure:"Name"`
}

type Common struct {
	LogDir     string `mapstructure:"LogDir"`
	TermEnable bool   `mapstructure:"TermEnable"`
}

type Auth struct {
	User string `mapstructure:"User"`
	Pass string `mapstructure:"Pass"`
}

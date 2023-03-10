package music_player

import (
	"errors"
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DB     DBConfig     `yaml:"db"`
	Server ServerConfig `yaml:"server"`
}

type DBConfig struct {
	DBName   string `yaml:"dbname" env:"DBNAME"`
	Username string `yaml:"username" env:"DBUSERNAME"`
	Password string `yaml:"password" env:"DBPASSWORD"`
	Host     string `yaml:"host" env:"DBHOST"`
	Port     string `yaml:"port" env:"DBPORT"`
	SSLMode  string `yaml:"sslmode" env:"SSLMODE"`
}

type ServerConfig struct {
	Port string `yaml:"port" env:"PORT"`
}

func (c *Config) Validate() error {
	if c.Server.Port == "" {
		return errors.New("empty server config")
	}

	if c.DB.Host == "" && c.DB.Port == "" && c.DB.DBName == "" && c.DB.Username == "" && c.DB.Password == "" {
		return errors.New("empty database config")
	}
	return nil
}

func LoadConfig(configPath string, configName string) (*Config, error) {

	if configPath == "" || configName == "" {
		return nil, errors.New("config path or name is empty")
	}

	var cfg Config

	// envs always prioritized

	if err := env.Parse(&cfg.DB); err != nil {
		logrus.Warnf("Failed to parse DB envs: %s", err.Error())
		return nil, err
	}

	if err := env.Parse(&cfg.Server); err != nil {
		logrus.Warnf("Failed to parse Server envs: %s", err.Error())
		return nil, err
	}

	if err := cfg.Validate(); err == nil {
		return &cfg, nil
	}

	if err := readConfig(configPath, configName); err != nil {
		return nil, err
	}

	if err := viper.UnmarshalKey("db", &cfg.DB); err != nil {
		return nil, err
	}
	if err := viper.UnmarshalKey("server", &cfg.Server); err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, errors.New("invalid data in config/envs")
	}
	return &cfg, nil
}

func readConfig(configPath, configName string) error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	return viper.ReadInConfig()
}

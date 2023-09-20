package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type Config struct {
	Server   Server   `toml:"server"`
	Database Database `toml:"database"`
}

type Server struct {
	Host string `toml:"host" default:"0.0.0.0"`
	Port int    `toml:"port" default:"8080"`
}

type Database struct {
	Host         string `toml:"host" default:"127.0.0.1"`
	Port         int    `toml:"port" default:"5432"`
	Database     string `toml:"database" default:"postgres"`
	Scheme       string `toml:"scheme" default:"public"`
	Username     string `toml:"username" default:"postgres"`
	Password     string `toml:"password" default:""`
	SslMode      bool   `toml:"sslMode" default:"false"`
	MaxIdleConns int    `toml:"maxIdleConns" default:"5"`
	MaxOpenConns int    `toml:"maxOpenConns" default:"10"`
}

func FromString(data string) (*Config, error) {
	var config Config
	_, err := toml.Decode(os.ExpandEnv(data), &config)
	return &config, err
}

func FromFile(path string) (*Config, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	_, err = toml.Decode(os.ExpandEnv(string(content)), &config)
	return &config, err
}

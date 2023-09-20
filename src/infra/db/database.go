package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var Client DBClient

//goland:noinspection GoNameStartsWithPackageName
type DBClient struct {
	*gorm.DB
}

//goland:noinspection SpellCheckingInspection
type Config struct {
	Host            string
	Port            int
	Database        string
	Scheme          string
	Username        string
	Password        string
	SslMode         bool
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

type DatabaseOption = func(*Config)

func NewConfig(host string, port int, username, password string, opts ...DatabaseOption) *Config {
	connMaxLifetime, _ := time.ParseDuration("1h")
	connMaxIdleTime, _ := time.ParseDuration("30m")
	config := &Config{
		Host:            host,
		Port:            port,
		Database:        "postgres",
		Scheme:          "public",
		Username:        username,
		Password:        password,
		MaxIdleConns:    5,
		MaxOpenConns:    10,
		ConnMaxLifetime: connMaxLifetime,
		ConnMaxIdleTime: connMaxIdleTime,
	}
	for _, opt := range opts {
		if opt != nil {
			opt(config)
		}
	}
	return config
}

func WithDatabase(database string) DatabaseOption {
	return func(config *Config) {
		config.Database = database
	}
}

func WithScheme(scheme string) DatabaseOption {
	return func(config *Config) {
		config.Scheme = scheme
	}
}

func WithSslMode(enable bool) DatabaseOption {
	return func(config *Config) {
		config.SslMode = enable
	}
}

func WithMaxIdleConns(conn int) DatabaseOption {
	return func(config *Config) {
		config.MaxIdleConns = conn
	}
}

func WithMaxOpenConns(conn int) DatabaseOption {
	return func(config *Config) {
		config.MaxOpenConns = conn
	}
}

func WithConnMaxLifetime(time time.Duration) DatabaseOption {
	return func(config *Config) {
		config.ConnMaxLifetime = time
	}
}

func WithConnMaxIdleTime(time time.Duration) DatabaseOption {
	return func(config *Config) {
		config.ConnMaxIdleTime = time
	}
}

func InitialDatabase(config *Config) (*DBClient, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s",
		config.Host, config.Port, config.Database, config.Username, config.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Scheme + ".",
			SingularTable: true,
		},
	})
	pool, err := db.DB()
	if err != nil {
		return nil, err
	}

	pool.SetConnMaxLifetime(config.ConnMaxLifetime)
	pool.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	pool.SetMaxOpenConns(config.MaxOpenConns)
	pool.SetMaxIdleConns(config.MaxIdleConns)

	client := &DBClient{
		DB: db,
	}
	Client = *client
	return client, nil
}

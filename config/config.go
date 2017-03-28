package config

import (
	"fmt"
	"os"
	"strconv"
)

// PostgresConfig ...
type PostgresConfig struct {
	Username string
	Host     string
	Password string
	Params   string
	Database string
	Port     int
}

// S3Config ...
type S3Config struct {
	Bucket string
	Region string
}

// Config ...
type Config struct {
	PgConf            PostgresConfig
	S3Conf            S3Config
	JWTSecret         string
	FirebaseServerKey string
	MigrationsPath    string
}

// PostgresConfigToConnectionString ...
func (cfg PostgresConfig) PostgresConfigToConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Params)
}

var defaultConf = Config{
	PgConf: PostgresConfig{
		Username: "mat",
		Database: "tutee",
		Password: "",
		Host:     "localhost",
		Params:   "?sslmode=disable",
		Port:     5432,
	},
	S3Conf: S3Config{
		Bucket: "tuteepics",
		Region: "eu-central-1",
	},
	JWTSecret:         "secret",
	FirebaseServerKey: "",
	MigrationsPath:    "./migrations",
}

// ParseTuteeConfig ...
func ParseTuteeConfig() Config {
	cfg := defaultConf

	jwt := os.Getenv("JWT_SECRET")
	firebase := os.Getenv("FIREBASE_SERVER_KEY")
	username := os.Getenv("POSTGRES_USER")
	port := os.Getenv("POSTGRES_PORT")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	params := os.Getenv("POSTGRES_PARAMS")
	database := os.Getenv("POSTGRES_DB")
	path := os.Getenv("MIGRATIONS_PATH")

	if jwt != "" {
		cfg.JWTSecret = jwt
	}

	if firebase != "" {
		cfg.FirebaseServerKey = firebase
	}

	if path != "" {
		cfg.MigrationsPath = path
	}

	if username != "" {
		cfg.PgConf.Username = username
	}

	if password != "" {
		cfg.PgConf.Password = password
	}

	if host != "" {
		cfg.PgConf.Host = host
	}

	if database != "" {
		cfg.PgConf.Database = database
	}

	if params != "" {
		cfg.PgConf.Params = params
	}

	if port != "" {
		portInt, _ := strconv.Atoi(port)
		cfg.PgConf.Port = portInt
	}

	return cfg
}

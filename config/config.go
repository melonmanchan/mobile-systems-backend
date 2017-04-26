package config

import "os"

// S3Config ...
type S3Config struct {
	Bucket string
	Region string
}

// Config ...
type Config struct {
	S3Conf            S3Config
	DatabaseURL       string
	JWTSecret         string
	FirebaseServerKey string
	MigrationsPath    string
}

var defaultConf = Config{
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
	path := os.Getenv("MIGRATIONS_PATH")

	db_path := os.Getenv("DATABASE_URL")

	cfg.DatabaseURL = db_path

	if jwt != "" {
		cfg.JWTSecret = jwt
	}

	if firebase != "" {
		cfg.FirebaseServerKey = firebase
	}

	if path != "" {
		cfg.MigrationsPath = path
	}

	return cfg
}

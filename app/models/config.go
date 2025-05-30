package models

type Config struct {
	Env        string `json:"env" env:"ENV" yaml:"env"`
	AppPort    string `json:"app_port" env:"APP_PORT" yaml:"app_port"`
	DBHost     string `json:"db_host" env:"DB_HOST" yaml:"db_host"`
	DBUser     string `json:"db_user" env:"DB_USER" yaml:"db_user"`
	DBPassword string `json:"db_password" env:"DB_PASSWORD" yaml:"db_password"`
	DBName     string `json:"db_name" env:"DB_NAME" yaml:"db_name"`
	DBPort     string `json:"db_port" env:"DB_PORT" yaml:"db_port"`
}

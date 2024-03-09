package config

type Database struct {
	DSN         string `env:"EARTH_POSTGRES_URL"`
	AutoMigrate bool   `env:"EARTH_DATABASE_AUTO_MIGRATE" envDefault:"true"`
}

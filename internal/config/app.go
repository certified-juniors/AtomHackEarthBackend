package config

type App struct {
	ErrorLevel string `env:"EARTH_ERROR_LEVEL" envDefault:"info"`

	API      API
	Database Database
	Minio    Minio
}

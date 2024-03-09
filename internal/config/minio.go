package config

type Minio struct {
	Endpoint      string `env:"EARTH_MINIO_URL"`
	MinioHost     string `env:"EARTH_MINIO_HOST"`
	MinioPort     string `env:"EARTH_MINIO_PORT"`
	MinioUser     string `env:"EARTH_MINIO_ROOT_USER"`
	MinioPassword string `env:"EARTH_MINIO_ROOT_PASSWORD"`
	MinioBucket   string `env:"EARTH_MINIO_BUCKET"`
}
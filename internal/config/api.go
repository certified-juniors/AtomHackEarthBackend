package config

type API struct {
	ServiceHost string `env:"EARTH_SERVICE_HOST" envDefault:"0.0.0.0"`
	ServicePort int    `env:"EARTH_SERVICE_PORT" envDefault:"8081"`
}

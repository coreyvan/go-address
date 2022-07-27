package app

type Config struct {
	Host string `conf:"default:localhost"`
	Port int    `conf:"default:3000"`
}

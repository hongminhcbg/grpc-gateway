package conf

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HttpPort string `envconfig:"HTTP_PORT" default:":9000"`
	GrpcPort string `envconfig:"HTTP_PORT" default:":10000"`
	Cors     *Cors
}

type Cors struct {
	AccessControlAllowOrigin  string `envconfig:"CORS_ACCESS_CONTROL_ALLOW_ORIGIN" default:"*"`
	AccessCOntrolAllowMethods string `envconfig:"CORS_ACCESS_CONTROL_ALLOW_METHODS" default:"POST, GET, OPTIONS, PUT, DELETE"`
	AccessCOntrolAllowHeaders string `envconfig:"CORS_ACCESS_CONTROL_ALLOW_HEADERS" default:"Accept,Content-Type,Content-Length,Accept-Encoding,Authorization,X-CSRF-Token,Device-Id"`
}

func GetConfig() (*Config, error) {
	var conf Config
	err := envconfig.Process("", &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

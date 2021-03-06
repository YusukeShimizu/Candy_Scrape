package env

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	// Config of the app.
	Config struct {
		Env      string `required:"true" envconfig:"ENV"`
		Pace     string `required:"true" envconfig:"PACE"`
		Secret   string `required:"true" envconfig:"SECRET"`
		Token    string `required:"true" envconfig:"TOKEN"`
		ID       string `required:"true" envconfig:"ID"`
		PUBLICID string `required:"true" envconfig:"ID"`
		PORT     string `required:"true" envconfig:"PORT"`
	}
)

// Process environment variables and create Config.
func Process() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

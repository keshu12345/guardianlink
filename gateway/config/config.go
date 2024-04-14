package config

import (
	"os"
	"path"

	"github.com/keshu12345/guardianlink/toolkit"
	logger "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

const (
	serverYML = "server.yml"
)

// NewFxModule returns the fx.Option that builds the *Configuration struct
// that could be later used by other fx modules.
func NewFxModule(configDirPath string, overridePath string) fx.Option {
	return fx.Provide(
		func() (*Configuration, error) {
			var conf Configuration
			if len(configDirPath) == 0 {
				logger.Info("trying env config path ")
				configDirPath = os.Getenv("CONFIG_PATH")
			}
			logger.Info("Using config path ", path.Join(configDirPath, serverYML))
			err := toolkit.NewConfig(&conf, path.Join(configDirPath, serverYML), overridePath)
			return &conf, err
		},
	)
}

type Configuration struct {
	EnvironmentName string
	Server          Server  `mapstructure:"server"`
	Swagger         Swagger `mapstructure:"swagger"`
	SQLite          DB      `mapstructure:"db"`
}
type Swagger struct {
	Host string
}

type Server struct {
	Port         int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

type DB struct {
	Filepath     string
	Databasename string
}

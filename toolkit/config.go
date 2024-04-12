package toolkit

import (
	"flag"
	"path/filepath"
	"strings"

	logger "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func NewConfig(conf interface{}, configPath, overridePath string, envVars ...map[string]string) error {

	flag.Parse()

	v := viper.New()

	//load configuration from file
	configDir := filepath.Dir(configPath)
	fileName := filepath.Base(configPath)
	file := strings.Split(fileName, ".")
	// load configurations from file in relative and absolute config path
	v.AddConfigPath(configDir)
	v.AutomaticEnv()

	v.SetConfigName(file[0])
	v.SetConfigType(file[1])

	for _, mp := range envVars {
		for ymlField, environmentVar := range mp {
			errV := v.BindEnv(ymlField, environmentVar)
			if errV != nil {
				logger.WithFields(logger.Fields{
					"config key": ymlField,
					"env var":    environmentVar,
					"error":      errV,
				}).Error("Error in binding env var")
			}
		}
	}

	var err error = nil
	err = v.ReadInConfig()
	if err != nil {
		return err
	}

	// Check if config should be overridden.
	if overridePath != "" {
		fileName = filepath.Base(overridePath)
		file = strings.Split(fileName, ".")
		v.AddConfigPath(filepath.Dir(overridePath))
		v.SetConfigName(file[0])
		v.SetConfigType(file[1])
		err = v.MergeInConfig()
		if err != nil {
			return err
		}
	}

	err = v.Unmarshal(conf)
	return err
}

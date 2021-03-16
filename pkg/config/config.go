package config

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Unmarshal parses and unmarshals configuration files and env variables into provided interface.
func Unmarshal(v interface{}, options ...ConfigOption) error {
	config, err := parseConfigOptions(options...)
	if err != nil {
		return err
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if config.RawConfig != nil {
		err := viper.ReadConfig(bytes.NewBuffer(config.RawConfig))
		if err != nil {
			return fmt.Errorf("failed to provided config: %v", err)
		}
	} else {
		err := viper.ReadInConfig()
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return fmt.Errorf("failed to read default config: %v", err)
			}
		}
	}

	if err := traverseIface(&ifaceProcess{}, v); err != nil {
		return fmt.Errorf("failed to traverse config: %v", err)
	}

	if err := viper.Unmarshal(v); err != nil {
		return fmt.Errorf("failed to unmarshal Viper config file: %v", err)
	}

	return nil
}

func parseConfigOptions(options ...ConfigOption) (*Config, error) {
	config := &Config{}

	for _, op := range options {
		err := op(config)
		if err != nil {
			return nil, fmt.Errorf("failed to parse config's configuration")
		}
	}

	return config, nil
}

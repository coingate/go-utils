package config

// ConfigOption base type for config options.
type ConfigOption func(*Config) error

// Config data structure for config options
type Config struct {
	RawConfig  []byte
	ConfigName string
}

// RawConfigOption sets raw config option
func RawConfigOption(cfg []byte) ConfigOption {
	return func(co *Config) error {
		co.RawConfig = cfg

		return nil
	}
}

// ConfigNameOption sets custom config name
func ConfigNameOption(name string) ConfigOption {
	return func(co *Config) error {
		co.ConfigName = name

		return nil
	}
}

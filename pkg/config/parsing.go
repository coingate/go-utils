package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type ifaceProcess struct{}

func (i *ifaceProcess) process(mapKey string) error {
	if err := viper.BindEnv(mapKey); err != nil {
		return fmt.Errorf("failed to bind env '%s' var: %v", mapKey, err)
	}

	return nil
}

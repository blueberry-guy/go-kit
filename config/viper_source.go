package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type viperConfigSource struct {
	v *viper.Viper
}

func newViperConfigSourceWithParams(configName string, configType string, configPaths ...string) ConfigSource {
	v := viper.New()
	v.SetConfigName(configName)
	v.SetConfigType(configType)
	if len(configPaths) > 0 {
		for _, cp := range configPaths {
			v.AddConfigPath(cp) // path to look for the config file
		}
	}
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	vcs := viperConfigSource{
		v,
	}
	return &vcs
}

func newViperConfigSource() ConfigSource {
	v := viper.New()
	v.AutomaticEnv()
	vcs := viperConfigSource{
		v,
	}
	return &vcs
}

func (vcs *viperConfigSource) GetInt(key string) int {
	return vcs.v.GetInt(key)
}

func (vcs *viperConfigSource) GetIntWithDefaultValue(key string, fallback int) int {
	valueInt := vcs.v.GetInt(key)
	if valueInt == 0 {
		return fallback
	}
	return valueInt
}

func (vcs *viperConfigSource) GetInt32(key string) int32 {
	return vcs.v.GetInt32(key)
}

func (vcs *viperConfigSource) GetInt64(key string) int64 {
	return vcs.v.GetInt64(key)
}

func (vcs *viperConfigSource) GetFloat64(key string) float64 {
	return vcs.v.GetFloat64(key)
}

func (vcs *viperConfigSource) GetFloat64WithDefaultValue(key string, fallback float64) float64 {
	valueFloat := vcs.v.GetFloat64(key)
	if valueFloat == 0 {
		return fallback
	}
	return valueFloat
}

func (vcs *viperConfigSource) GetString(key string) string {
	return vcs.v.GetString(key)
}

func (vcs *viperConfigSource) GetStringWithDefaultValue(key string, fallback string) string {
	value := vcs.v.GetString(key)
	if len(value) == 0 {
		return fallback
	}
	return value

}

func (vcs *viperConfigSource) GetBool(key string) bool {
	return vcs.v.GetBool(key)
}

func (vcs *viperConfigSource) GetStringSlice(key string) []string {
	return vcs.v.GetStringSlice(key)
}

func (vcs *viperConfigSource) Get(key string) interface{} {
	return vcs.v.Get(key)
}

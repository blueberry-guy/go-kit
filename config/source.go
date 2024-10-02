package config

type ConfigSource interface {
	GetInt(key string) int
	GetIntWithDefaultValue(key string, fallback int) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetFloat64(key string) float64
	GetFloat64WithDefaultValue(key string, fallback float64) float64
	GetString(key string) string
	GetBool(key string) bool
	GetStringSlice(key string) []string
	Get(key string) interface{}
	GetStringWithDefaultValue(key string, fallback string) string
}

func NewConfigSource(configName string, configType string, configPaths ...string) ConfigSource {
	return newViperConfigSourceWithParams(configName, configType, configPaths...)
}

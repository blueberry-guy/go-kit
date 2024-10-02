package config

import "github.com/blueberry-guy/go-kit/util"

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

type SourceParam struct {
	FileName    string
	FileExt     string
	Directories []string
}

func NewConfigSource(param *SourceParam) ConfigSource {
	if param == nil || (util.IsBlank(param.FileName) || util.IsBlank(param.FileExt) || len(param.Directories) == 0) {
		return newViperConfigSource()
	} else {
		return newViperConfigSourceWithParams(param.FileName, param.FileExt, param.Directories...)
	}
}

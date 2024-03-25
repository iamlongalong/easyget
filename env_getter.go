package easyget

import (
	"os"
	"strings"
)

// read string from env
type EnvGetter struct {
	prefix  string
	filters []IFilter
}

func NewEnvGetter(prefix string) *EnvGetter {
	return &EnvGetter{
		prefix: prefix,
	}
}

func (sg *EnvGetter) Get(key string) (string, bool) {
	fullKey := sg.prefix + key

	value, ok := os.LookupEnv(fullKey)
	if !ok {
		return "", false
	}

	return value, true
}

func (sg *EnvGetter) Gets() []KVPair {
	var pairs []KVPair
	for _, env := range os.Environ() {
		index := strings.Index(env, "=")
		envKey, envValue := env[:index], env[index+1:]
		if strings.HasPrefix(envKey, sg.prefix) {
			pair := KVPair{
				Key:   strings.TrimPrefix(envKey, sg.prefix),
				Value: envValue,
			}
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

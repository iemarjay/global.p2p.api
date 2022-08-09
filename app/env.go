package app

import "github.com/joho/godotenv"

type Env struct {
	value map[string]string
}

func (env Env) Set(value map[string]string) {
	env.value = value
}

func (env Env) GetOrDefault(key string, fail string) string {
	if value, exists := env.value[key]; exists {
		return value
	}

	return fail
}

func (env Env) Get(key string) string {
	if value, exists := env.value[key]; exists {
		return value
	}

	return ""
}

func NewEnv(envPath string) (*Env, error) {
	value, err := godotenv.Read(envPath)
	if err != nil {
		return nil, err
	}

	var env = &Env{
		value: value,
	}

	return env, err
}

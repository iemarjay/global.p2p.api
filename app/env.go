package app

import "github.com/joho/godotenv"

type Env struct {
	value map[string]string
}

const APP_PORT = "APP_PORT"
const APP_BASE_URL = "APP_BASE_URL"

const PUBLIC_ROOT_DIR = "PUBLIC_ROOT_DIR"
const PUBLIC_ROOT_DIR_VALUE = "resources/public"

const PUBLIC_PATH_PREFIX = "PUBLIC_PATH_PREFIX"
const PUBLIC_PATH_PREFIX_VALUE = "public"

const DATABASE_URL = "DATABASE_URL"
const DATABASE_NAME = "DATABASE_NAME"



func (env *Env) Set(value map[string]string) {
	env.value = value
}

func (env *Env) GetOrDefault(key string, fail string) string {
	if value, exists := env.value[key]; exists {
		return value
	}

	return fail
}

func (env *Env) Get(key string) string {
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

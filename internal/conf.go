package internal

import (
	"github.com/pkg/errors"
	"os"
)

type Config struct {
	Cmd string
	Identifier string
}

func ConfigFromEnvironment() (Config, error) {
	c := Config{}
	var ok bool

	c.Cmd, ok = os.LookupEnv("DOCKER_EXECUTOR_CMD")
	if !ok {
		return c, errors.New("DOCKER_EXECUTOR_CMD must be defined in the environment")
	}

	c.Identifier, ok = os.LookupEnv("DOCKER_EXECUTOR_IDENTIFIER")
	if !ok {
		return c, errors.New("DOCKER_EXECUTOR_IDENTIFIER must be defined in the environment")
	}

	return c, nil
}

package internal

import (
	"github.com/pkg/errors"
	"os"
	"strconv"
)

type Config struct {
	Cmd string
	Identifier string
	Port int
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
	port, ok := os.LookupEnv("DOCKER_EXECUTOR_PORT")
	if !ok {
		port = "8080"
	}

	var err error
	c.Port, err = strconv.Atoi(port)
	if err != nil {
		return c, errors.Wrap(err, "Could not parse port")
	}

	return c, nil
}

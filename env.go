package chainlib

import (
	"fmt"
	"regexp"
	"strings"
)

var safeEnvRegex = regexp.MustCompile("[^a-zA-Z0-9_]")

type Environment map[string]string

func NewEnvironment(vars ...string) Environment {
	env := Environment{}
	for _, envvar := range vars {
		parts := strings.SplitN(envvar, "=", 2)
		env[parts[0]] = parts[1]
	}
	return env
}

func (e Environment) Environ() []string {
	env := make([]string, 0)
	for k, v := range e {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	return env
}

func (e Environment) Merge(that Environment) {
	for k, v := range that {
		e[k] = v
	}
}

func SafeEnv(key string) string {
	return safeEnvRegex.ReplaceAllString(key, "_")
}

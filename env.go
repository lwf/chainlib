package chainlib

import "regexp"

var safeEnvRegex = regexp.MustCompile("[^a-zA-Z0-9_]")

func SafeEnv(key string) string {
	return safeEnvRegex.ReplaceAllString(key, "_")
}

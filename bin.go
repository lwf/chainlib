package chainlib

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func assert(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Convenience function for writing utilities which use Bernstein chaining to
// pass environment variables. Function f should perform logic to retrieve data
// to be serialized to environment variables and return them as a map of strings,
// or error if something went wrong.
func Main(prefix string, f func() (map[string]string, error)) {
	var upcase bool
	var rprefix string
	flag.StringVar(&rprefix, "prefix", prefix, "Prefix")
	flag.BoolVar(&upcase, "upcase", true, "Upcase variables")
	flag.Parse()

	env, err := f()
	assert(err)
	var menv []string
	for k, v := range env {
		key := SafeEnv(k)
		if upcase {
			key = strings.ToUpper(key)
		}
		menv = append(menv, fmt.Sprintf("%s_%s=%s", rprefix, key, v))
	}
	assert(Exec(flag.Args(), menv))
}

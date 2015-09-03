package chainlib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// Takes a string and returns path to binary if found on PATH.
func FindExec(s string) (string, error) {
	if s[0] == '/' {
		return s, nil
	}

	paths := strings.Split(os.Getenv("PATH"), ":")
	for _, path := range paths {
		file := filepath.Join(path, s)
		stat, err := os.Stat(file)
		if err != nil {
			continue
		}
		mode := stat.Mode()
		if mode&0100 == 0100 || mode&0001 == 0001 {
			return file, nil
		}
	}
	return "", fmt.Errorf("could not find executable %s on PATH", s)
}

// Execs into the program specified by args and passes environment
// variables in env.
func Exec(args []string, env []string) error {
	if len(args) > 0 {
		bin, err := FindExec(args[0])
		if err != nil {
			return err
		}
		menv := os.Environ()
		if env != nil {
			menv = append(menv, env...)
		}
		return syscall.Exec(bin, args, menv)
	} else {
		for _, v := range env {
			fmt.Println(v)
		}
	}
	return nil
}

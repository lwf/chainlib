package chainlib

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// Execs into the program specified by args and passes environment
// variables in env.
func Exec(args []string, env Environment) error {
	if len(args) > 0 {
		bin, err := exec.LookPath(args[0])
		if err != nil {
			return err
		}
		menv := NewEnvironment(os.Environ()...)
		menv.Merge(env)
		return syscall.Exec(bin, args, menv.Environ())
	} else {
		for _, v := range env.Environ() {
			fmt.Println(v)
		}
	}
	return nil
}

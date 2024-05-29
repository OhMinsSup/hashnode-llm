package cmd

import (
	"os"
)

// getCommandLineExecutable returns the name of the executable that was invoked
func getCommandLineExecutable() string {
	return os.Args[0]
}

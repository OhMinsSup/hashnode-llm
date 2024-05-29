package cmd

import (
	"context"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	NewRootCmd()
}

// NewRootCmd - create the root command
func NewRootCmd() *cobra.Command {
	RootCmd := &cobra.Command{
		Use:   getCommandLineExecutable(),
		Short: "Hashnode LLM",
		Long:  "Hashnode LLM is a CLI tool for Hashnode's Learning Lab Manager. ",
	}

	RootCmd.AddCommand(newServeCmd())

	return RootCmd
}

func Execute() {
	RootCmd := NewRootCmd()
	RootCmd.SetContext(context.Background())
	RootCmd.SetOutput(os.Stdout)
	if err := RootCmd.Execute(); err != nil {
		msg := err.Error()
		if len(msg) > 0 {
			// add newline if needed
			if !strings.HasSuffix(msg, "\n") {
				msg += "\n"
			}
			RootCmd.Print(msg)
		}
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"log"

	"github.com/OhMinsSup/tavoli/config"
	"github.com/OhMinsSup/tavoli/server"
	"github.com/spf13/cobra"
)

func NewServeConfig() (*config.Configuration, error) {
	config, err := config.ReadConfigFile("./config.json")
	if err != nil {
		return nil, fmt.Errorf("failed to load server config: %v", err)
	}
	return config, nil
}

func newServeCmd() *cobra.Command {
	serveConfig, err := NewServeConfig()
	if err != nil {
		log.Fatalf("failed to create serve options: %v", err)
	}

	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the hashnode llm server.",
		Long:  "Start the hashnode llm server.",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := serve(cmd, serveConfig)
			if err != nil {
				return fmt.Errorf("failed to start server: %v", err)
			}
			return nil
		},
	}

	return serveCmd
}

func serve(cmd *cobra.Command, config *config.Configuration) error {
	_, err := server.NewServer(config)
	if err != nil {
		return err
	}
	return nil
}

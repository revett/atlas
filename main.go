package main

import (
	"context"
	_ "embed"
	"os"

	"github.com/revett/atlas/internal/cmd"
	"github.com/revett/atlas/internal/cmdv2"
	"github.com/revett/atlas/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

//go:embed VERSION
var version string

func main() {
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out: os.Stdout,
		},
	)

	cfg, err := config.NewConfig(version)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	root := cmd.Root()

	// Commands from cmd package (v1).
	root.AddCommand(
		cmd.Doctor(),
	)

	// Commands from cmdv2 package.
	commands := []*cobra.Command{}
	for _, command := range cmdv2.RegisteredCommands() {
		commands = append(commands, command.Command())
	}
	root.AddCommand(commands...)

	ctx := context.WithValue(
		context.Background(), config.ContextConfigKey, *cfg,
	)

	if err := root.ExecuteContext(ctx); err != nil {
		log.Fatal().Err(err).Send()
	}
}

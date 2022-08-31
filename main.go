package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/revett/sepias/internal/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	// BuildFlag allows for the output from the version command to have a prefix
	// based on what is passed via ldflags when the CLI is built.
	BuildFlag = "" //nolint:gochecknoglobals

	//go:embed VERSION
	version string
)

func main() {
	rand.Seed(time.Now().UnixNano())

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out: os.Stdout,
		},
	)

	root := cmd.Root()
	root.AddCommand(
		cmd.Completion(),
		cmd.Doctor(),
		cmd.Version(
			generateVersion(),
		),
	)

	if err := root.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func generateVersion() string {
	if BuildFlag == "" {
		return version
	}

	return fmt.Sprintf("%s-%s", BuildFlag, version)
}

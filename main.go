package main

import (
	_ "embed"
	"math/rand"
	"os"
	"time"

	"github.com/revett/sepias/internal/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//go:embed VERSION
var version string

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
		cmd.Version(version),
	)

	if err := root.Execute(); err != nil {
		log.Fatal().Err(err).Send()
	}
}

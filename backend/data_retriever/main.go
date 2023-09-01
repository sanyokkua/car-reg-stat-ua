package main

import (
	"data_retriever/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("App Started")

	err := app.GetData(false)
	if err != nil {
		panic(err)
	}

	log.Info().Msg("App Finished")
}

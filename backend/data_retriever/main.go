package main

import (
    "data_retriever/app"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "os"
)

const CLEAN_RUN = "clean"

func main() {
    log.Info().Msg("App Started")
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

    args := os.Args[1:]
    var cleanRun bool
    for _, arg := range args {
        if arg == CLEAN_RUN {
            cleanRun = true
            break
        }
    }

    err := app.ProcessRegistrationData(cleanRun)
    if err != nil {
        log.Error().Err(err).Msg("Application failed to process Registration Data")
        panic(err)
    }

    log.Info().Msg("App Finished")
}

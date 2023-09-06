package app

import (
    koatuuApp "data_retriever/data_sources/koatuu/app"
    registrationApp "data_retriever/data_sources/registrations/app"
    "github.com/rs/zerolog/log"
)

func ProcessApplicationData(cleanRun bool) error {
    defer func() {
        // Global handler for panic
        if r := recover(); r != nil {
            log.Error().Msgf("Application failed to process data, %s", r)
        }
    }()
    log.Debug().Msg("App Started")

    err := koatuuApp.ProcessKoatuuData()
    if err != nil {
        return err
    }

    err = registrationApp.ProcessRegistrationData(cleanRun)
    if err != nil {
        return err
    }

    log.Debug().Msg("App Finished")

    return nil
}

package utils

import (
    "github.com/rs/zerolog/log"
    "io"
)

func CloseFunc(closableItem io.Closer) {
    // Helper function for defer keyword with logging of error if any happens
    err := closableItem.Close()
    if err != nil {
        log.Error().Err(err).Msg("Failed to close closableItem")
    }
}

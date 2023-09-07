package app

import (
	commonCsv "data_retriever/common/utils/csv"
	"data_retriever/data_sources/koatuu/constants"
	"data_retriever/data_sources/koatuu/csv"
	"github.com/rs/zerolog/log"
	"os"
	"path"
)

func ProcessKoatuuData() error {
	defer func() {
		// Global handler for panic
		if r := recover(); r != nil {
			log.Error().Msgf("Application failed to process data, %s", r)
		}
	}()

	getwd, err := os.Getwd()
	if err != nil {
		return err
	}

	csvPath := path.Join(getwd, constants.KoatuuFileName)
	records, err := commonCsv.ParseRecords(csvPath, csv.ToKOATUUJsonRecord)
	if err != nil {
		return err
	}
	log.Debug().Msgf("Num of records: %d", len(records))

	return nil
}

package csv

import (
	"data_retriever/common/utils"
	"data_retriever/common/utils/files"
	"encoding/csv"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

type Mapper[IN, OUT any] func(value IN) *OUT

func ParseRecords[T any](csvFilePath string, mapper Mapper[map[string]string, T]) ([]T, error) {
	// Check preconditions before opening file
	validationErr := files.CheckForValidCsvFile(csvFilePath)
	if validationErr != nil {
		return nil, validationErr
	}

	// Open CSV file for further processing
	openedCsvFile, openFileErr := os.Open(csvFilePath)
	if openFileErr != nil {
		log.Error().Err(openFileErr).Msgf("Failed to open %s", csvFilePath)
		return nil, openFileErr
	}
	defer utils.CloseFunc(openedCsvFile)

	// Create a new CSV csvFileReader and configure delimiter of csvFileRecords
	csvFileReader := csv.NewReader(openedCsvFile)
	csvFileReader.Comma = ';'

	// Read all the csvFileRecords to memory
	csvFileRecords, readErr := csvFileReader.ReadAll()
	if readErr != nil {
		log.Error().Err(readErr).Msg("Failed to read CSV file")
		return nil, readErr
	}

	// Prepare containers for keeping records
	csvRecordMap := make(map[string]string)
	mappedRecords := make([]T, 0, len(csvFileRecords)-1) // -1 because first line is headers line
	headerLine := csvFileRecords[0]
	for i := range headerLine {
		// Headers can be lowercase in old csv files, so we need to make them all in one format
		headerLine[i] = strings.ToUpper(headerLine[i])
	}

	// Map all records from CSV file to models.CsvRecord
	for i, record := range csvFileRecords {
		if i == 0 {
			// Skip first header line
			continue
		}

		for index, value := range record {
			csvRecordMap[headerLine[index]] = value
		}

		mappedRecord := mapper(csvRecordMap)
		mappedRecords = append(mappedRecords, *mappedRecord)
	}

	log.Debug().Msgf("Number of csvFileRecords in the %s: %d", csvFilePath, len(mappedRecords))

	return mappedRecords, nil
}

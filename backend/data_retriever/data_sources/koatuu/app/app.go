package app

import (
    "data_retriever/common/utils"
    "data_retriever/common/utils/files"
    "data_retriever/data_sources/koatuu/constants"
    "data_retriever/data_sources/koatuu/models"
    "encoding/csv"
    "github.com/rs/zerolog/log"
    "os"
    "path"
    "strings"
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

    csvPath := path.Join(getwd, constants.KOATUU_FILE_NAME)
    records, err := ParseRecords(csvPath)
    if err != nil {
        return err
    }
    log.Debug().Msgf("Num of records: %d", len(records))

    return nil
}

func ParseRecords(csvFilePath string) ([]models.KOATUUJsonRecord, error) {
    // Check preconditions before opening file
    validationErr := files.IsExistingCsvFile(csvFilePath)
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
    mappedRecords := make([]models.KOATUUJsonRecord, 0, len(csvFileRecords)-1) // -1 because first line is headers line
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

        mappedRecord := convertMapToStruct(csvRecordMap)
        mappedRecords = append(mappedRecords, *mappedRecord)
    }

    log.Debug().Msgf("Number of csvFileRecords in the %s: %d", csvFilePath, len(mappedRecords))

    return mappedRecords, nil
}

func convertMapToStruct(recordMap map[string]string) *models.KOATUUJsonRecord {
    lev1 := recordMap[constants.KEY_KOATUU_LEVEL_1]
    lev2 := recordMap[constants.KEY_KOATUU_LEVEL_2]
    lev3 := recordMap[constants.KEY_KOATUU_LEVEL_3]
    lev4 := recordMap[constants.KEY_KOATUU_LEVEL_4]
    cat := recordMap[constants.KEY_KOATUU_CATEGORY]
    name := recordMap[constants.KEY_KOATUU_NAME]

    return &models.KOATUUJsonRecord{
        Level1:   lev1,
        Level2:   lev2,
        Level3:   lev3,
        Level4:   lev4,
        Category: cat,
        Name:     name,
    }
}

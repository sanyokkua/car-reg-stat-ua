package csv

import (
    "data_retriever/common/utils"
    "data_retriever/common/utils/files"
    "data_retriever/data_sources/registrations/models"
    "encoding/csv"
    "encoding/json"
    "errors"
    "github.com/rs/zerolog/log"
    "os"
    "strings"
)

func ParseDataPackageJsonFile(path string) (*models.DataPackageJson, error) {
    log.Debug().Msgf("ParseDataPackageJsonFile Path: %s", path)

    if !files.IsFileExist(path) {
        return nil, errors.New("file can't be opened because it doesn't exist")
    }

    // Read the file content
    data, readErr := os.ReadFile(path)
    if readErr != nil {
        log.Error().Err(readErr).Msg("Error happened during reading file content")
        return nil, readErr
    }

    // Declare a variable of type DataPackageJson for parsing json
    var dataPackage models.DataPackageJson

    // Unmarshal the JSON data into the dataPackage variable
    unmarshalErr := json.Unmarshal(data, &dataPackage)
    if unmarshalErr != nil {
        log.Error().Err(unmarshalErr).Msg("Error happened during parsing json file")
        return nil, unmarshalErr
    }

    log.Debug().Msgf("Parsed DataPackageJson json")

    return &dataPackage, nil
}

func FindUrlsOfCSVFilesInDataPackageJson(dataPackage *models.DataPackageJson) ([]string, error) {
    if dataPackage == nil || dataPackage.Resources == nil || len(dataPackage.Resources) == 0 {
        return nil, errors.New("dataPackage resources are empty")
    }

    urls := make([]string, 0, len(dataPackage.Resources))

    for i, resource := range dataPackage.Resources {
        log.Debug().Msgf("Link %d -- %s", i, resource.Path)
        urls = append(urls, resource.Path)
    }

    return urls, nil
}

func ParseRegistrationsCsvToRecordsArray(csvFilePath string) ([]models.CsvRecord, error) {
    // Check preconditions before opening file
    validationErr := validatePath(csvFilePath)
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
    mappedRecords := make([]models.CsvRecord, 0, len(csvFileRecords)-1) // -1 because first line is headers line
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

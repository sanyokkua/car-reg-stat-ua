package utils

import (
	"data_retriever/constants"
	"data_retriever/models"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"path"
	"strconv"
	"strings"
)

func validatePath(csvFilePath string) error {
	if csvFilePath == "" {
		return errors.New("file path is blank")
	}

	ext := path.Ext(csvFilePath)
	if ext != constants.CSV_FILE_EXTENSION {
		errMsg := fmt.Sprintf("extension of the file is not correct. Expected %s, Actual: %s", constants.CSV_FILE_EXTENSION, ext)
		return errors.New(errMsg)
	}

	if !IsFileExist(csvFilePath) {
		return errors.New("file is not exist")
	}

	return nil
}

func stringToInt(value string) int {
	// Some csv files has lowercase values, to simplify processing is better to make everything in Uppercase
	value = strings.ToUpper(value)
	if value == "" || value == "NULL" || value == "NONE" {
		// Possible expected values, better to check than parse
		return 0
	}

	if strings.Contains(value, ".") {
		// Some digits are float/double, and they should be interpreted as int values
		splitValue := strings.Split(value, ".")
		value = strings.TrimSpace(splitValue[0])
	}

	if strings.Contains(value, ",") {
		// Some digits are float/double, and they should be interpreted as int values
		splitValue := strings.Split(value, ",")
		value = strings.TrimSpace(splitValue[0])
	}

	parsedInt, err := strconv.Atoi(value)

	if err != nil {
		// It is ok on error to have 0 as default value
		log.Warn().Err(err).Msg("There were error with converting string to int")
	}

	return parsedInt
}

func fixString(value string) string {
	newValue := strings.ToUpper(value)
	return strings.TrimSpace(newValue)
}

func fixOperationName(opName string, opCode string) string {
	if strings.Contains(opName, opCode) {
		opName = strings.Replace(opName, opCode, "", 1)
		opName = strings.TrimSpace(opName)
		if opName[0] == '-' {
			opName = strings.Replace(opName, "-", "", 1)
		}
		opName = strings.TrimSpace(opName)
	}
	return fixString(opName)
}

func fixNameThatHasCode(name string, code string) string {
	if strings.Contains(name, code) {
		name = strings.Replace(name, code, "", 1)
		name = strings.TrimSpace(name)
	}

	return fixString(name)
}

func convertMapToStruct(dataObject map[string]string) models.CsvRecord {
	person := fixString(dataObject[constants.PERSON])                                                         // String
	koatuu := fixString(dataObject[constants.REGISTRATION_ADDRESS_KOATUU])                                    // Nullable
	opCode := stringToInt(dataObject[constants.OPERATION_CODE])                                               // Number
	opName := fixOperationName(dataObject[constants.OPERATION_NAME], dataObject[constants.OPERATION_CODE])    // Can contain opCode "code - XXXXX"
	dateRegistration := fixString(dataObject[constants.DATE_REGISTRATION])                                    // String, Date
	departmentCode := fixString(dataObject[constants.DEPARTMENT_CODE])                                        // Should be string, there are numbers that starts from 0XXX
	department := fixNameThatHasCode(dataObject[constants.DEPARTMENT], dataObject[constants.DEPARTMENT_CODE]) // Can contain departmentCode "XXXX code"
	brand := fixNameThatHasCode(dataObject[constants.BRAND], dataObject[constants.MODEL])                     // Can contain mode "Brand MODEL"
	model := fixString(dataObject[constants.MODEL])                                                           // String
	year := stringToInt(dataObject[constants.MAKE_YEAR])                                                      // Number
	color := fixString(dataObject[constants.COLOR])                                                           // String
	kind := fixString(dataObject[constants.KIND])                                                             // String
	body := fixString(dataObject[constants.BODY])                                                             // String
	purpose := fixString(dataObject[constants.PURPOSE])                                                       // String
	fuel := fixString(dataObject[constants.FUEL])                                                             // String, Nullable
	capacity := stringToInt(dataObject[constants.CAPACITY])                                                   // Number, Nullable
	weight := stringToInt(dataObject[constants.OWN_WEIGHT])                                                   // Number, Nullable
	totalWeight := stringToInt(dataObject[constants.TOTAL_WEIGHT])                                            // Number, Nullable
	registrationNew := fixString(dataObject[constants.NUMBER_REGISTRATION_NEW])                               // String, Nullable
	vin := fixString(dataObject[constants.VIN])                                                               // String, Nullable

	operation := models.Operation{
		Name: opName,
		Code: opCode,
	}

	departmentObj := models.Department{
		Name: department,
		Code: departmentCode,
	}

	characteristic := models.Characteristic{
		MakeYear:    year,
		Color:       color,
		Body:        body,
		Fuel:        fuel,
		Capacity:    capacity,
		OwnWeight:   weight,
		TotalWeight: totalWeight,
	}

	vehicle := models.Vehicle{
		Brand:          brand,
		Model:          model,
		Characteristic: characteristic,
	}

	return models.CsvRecord{
		Person:                    person,
		RegistrationAddressKoatuu: koatuu,
		Operation:                 operation,
		DateRegistration:          dateRegistration,
		Department:                departmentObj,
		Vehicle:                   vehicle,
		Kind:                      kind,
		Purpose:                   purpose,
		NumberRegistrationNew:     registrationNew,
		Vin:                       vin,
	}
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
	defer CloseFunc(openedCsvFile)

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
		mappedRecords = append(mappedRecords, mappedRecord)
	}

	log.Debug().Msgf("Number of csvFileRecords in the %s: %d", csvFilePath, len(mappedRecords))

	return mappedRecords, nil
}

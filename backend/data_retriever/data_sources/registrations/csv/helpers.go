package csv

import (
    "data_retriever/common/utils/files"
    "data_retriever/data_sources/registrations/constants"
    "data_retriever/data_sources/registrations/models"
    "errors"
    "fmt"
    "github.com/rs/zerolog/log"
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

    if !files.IsFileExist(csvFilePath) {
        return errors.New("file is not exist")
    }

    return nil
}

func stringToInt(value string) int {
    // Some csv files has lowercase values, to simplify processing is better to make everything in Uppercase
    value = fixNullableString(value)

    if len(value) == 0 || value == "NULL" || value == "NONE" {
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

    if strings.HasPrefix(value, "-") {
        // By specifics of processed data there can't be any numbers below zero
        value = strings.ReplaceAll(value, "-", "")
        value = strings.TrimSpace(value)
    }

    parsedInt, err := strconv.Atoi(value)

    if err != nil {
        // It is ok on error to have 0 as default value
        log.Warn().Err(err).Msg("There were error with converting string to int")
    }

    return parsedInt
}

func fixString(value string) string {
    if strings.Contains(value, "/") {
        value = strings.ReplaceAll(value, "/", "")
    }
    if strings.HasPrefix(value, "-") {
        // By specifics of processed data there can't be strings that starts from -
        value = strings.Replace(value, "-", "", 1)
        value = strings.TrimSpace(value)
    }
    if strings.HasPrefix(value, ".") {
        // By specifics of processed data there can't be strings that starts from '.'
        value = strings.Replace(value, ".", "", 1)
        value = strings.TrimSpace(value)
    }
    if strings.HasPrefix(value, ",") {
        // By specifics of processed data there can't be strings that starts from ','
        value = strings.Replace(value, ",", "", 1)
        value = strings.TrimSpace(value)
    }
    newValue := strings.ToUpper(value)
    return strings.TrimSpace(newValue)
}

func fixNullableString(value string) string {
    newValue := fixString(value)
    if newValue == "" || newValue == "NONE" || newValue == "NULL" || newValue == "." {
        newValue = ""
    }
    return newValue
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

func convertMapToStruct(dataObject map[string]string) *models.CsvRecord {
    person := fixString(dataObject[constants.PERSON])                                                         // String
    koatuu := fixNullableString(dataObject[constants.REGISTRATION_ADDRESS_KOATUU])                            // Nullable
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

    fuelValue := dataObject[constants.FUEL]
    fixedFuelValue := fixNullableString(fuelValue)
    fuel := constants.GetFuelMapping(fixedFuelValue)

    capacity := stringToInt(dataObject[constants.CAPACITY])                             // Number, Nullable
    weight := stringToInt(dataObject[constants.OWN_WEIGHT])                             // Number, Nullable
    totalWeight := stringToInt(dataObject[constants.TOTAL_WEIGHT])                      // Number, Nullable
    registrationNew := fixNullableString(dataObject[constants.NUMBER_REGISTRATION_NEW]) // String, Nullable
    vin := fixNullableString(dataObject[constants.VIN])                                 // String, Nullable

    return &models.CsvRecord{
        DepartmentName:            department,
        DepartmentCode:            departmentCode,
        OperationName:             opName,
        OperationCode:             opCode,
        Brand:                     brand,
        Model:                     model,
        MakeYear:                  year,
        Color:                     color,
        Body:                      body,
        Fuel:                      fuel,
        Capacity:                  capacity,
        OwnWeight:                 weight,
        TotalWeight:               totalWeight,
        Person:                    person,
        RegistrationAddressKoatuu: koatuu,
        DateRegistration:          dateRegistration,
        Kind:                      kind,
        Purpose:                   purpose,
        NumberRegistrationNew:     registrationNew,
        Vin:                       vin,
    }
}

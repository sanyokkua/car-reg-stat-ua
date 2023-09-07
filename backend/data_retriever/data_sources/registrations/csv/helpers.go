package csv

import (
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
)

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

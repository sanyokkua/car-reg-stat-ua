package constants

import "strings"

const DATA_PACKAGE_JSON_URL = "https://data.gov.ua/dataset/06779371-308f-42d7-895e-5a39833375f0/datapackage"
const CSV_FILE_EXTENSION = ".csv"
const FOLDER_ARCHIVES = "archives"
const FOLDER_CSV = "csv"
const FOLDER_JSON = "json"
const PERSON = "PERSON"
const REGISTRATION_ADDRESS_KOATUU = "REG_ADDR_KOATUU"
const OPERATION_CODE = "OPER_CODE"
const OPERATION_NAME = "OPER_NAME"
const DATE_REGISTRATION = "D_REG"
const DEPARTMENT_CODE = "DEP_CODE"
const DEPARTMENT = "DEP"
const BRAND = "BRAND"
const MODEL = "MODEL"
const MAKE_YEAR = "MAKE_YEAR"
const COLOR = "COLOR"
const KIND = "KIND"
const BODY = "BODY"
const PURPOSE = "PURPOSE"
const FUEL = "FUEL"
const CAPACITY = "CAPACITY"
const OWN_WEIGHT = "OWN_WEIGHT"
const TOTAL_WEIGHT = "TOTAL_WEIGHT"
const NUMBER_REGISTRATION_NEW = "N_REG_NEW"
const VIN = "VIN"

func GetFuelMapping(value string) string {
    value = strings.ToUpper(value)
    value = strings.TrimSpace(value)

    mapping := map[string]string{
        "":             "НЕВИЗНАЧЕНИЙ",
        ".":            "НЕВИЗНАЧЕНИЙ",
        "НЕ ВИЗНАЧЕНО": "НЕВИЗНАЧЕНИЙ",
        "NULL":         "НЕВИЗНАЧЕНИЙ",
    }

    mappedValue, found := mapping[value]
    if found {
        return mappedValue
    }

    return value
}

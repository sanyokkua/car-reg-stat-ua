package constants

import "strings"

const DataPackageJsonUrl = "https://data.gov.ua/dataset/06779371-308f-42d7-895e-5a39833375f0/datapackage"
const FolderArchives = "archives"
const FolderCsv = "csv"
const FolderJson = "json"
const PERSON = "PERSON"
const RegistrationAddressKoatuu = "REG_ADDR_KOATUU"
const OperationCode = "OPER_CODE"
const OperationName = "OPER_NAME"
const DateRegistration = "D_REG"
const DepartmentCode = "DEP_CODE"
const DEPARTMENT = "DEP"
const BRAND = "BRAND"
const MODEL = "MODEL"
const MakeYear = "MAKE_YEAR"
const COLOR = "COLOR"
const KIND = "KIND"
const BODY = "BODY"
const PURPOSE = "PURPOSE"
const FUEL = "FUEL"
const CAPACITY = "CAPACITY"
const OwnWeight = "OWN_WEIGHT"
const TotalWeight = "TOTAL_WEIGHT"
const NumberRegistrationNew = "N_REG_NEW"
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

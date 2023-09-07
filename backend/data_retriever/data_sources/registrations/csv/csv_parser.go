package csv

import (
	"data_retriever/data_sources/registrations/constants"
	"data_retriever/data_sources/registrations/models"
)

func ToRegistrationCsvRecord(dataObject map[string]string) *models.RegistrationCsvRecord {
	person := fixString(dataObject[constants.PERSON])                                                        // String
	koatuu := fixNullableString(dataObject[constants.RegistrationAddressKoatuu])                             // Nullable
	opCode := stringToInt(dataObject[constants.OperationCode])                                               // Number
	opName := fixOperationName(dataObject[constants.OperationName], dataObject[constants.OperationCode])     // Can contain opCode "code - XXXXX"
	dateRegistration := fixString(dataObject[constants.DateRegistration])                                    // String, Date
	departmentCode := fixString(dataObject[constants.DepartmentCode])                                        // Should be string, there are numbers that starts from 0XXX
	department := fixNameThatHasCode(dataObject[constants.DEPARTMENT], dataObject[constants.DepartmentCode]) // Can contain departmentCode "XXXX code"
	brand := fixNameThatHasCode(dataObject[constants.BRAND], dataObject[constants.MODEL])                    // Can contain mode "Brand MODEL"
	model := fixString(dataObject[constants.MODEL])                                                          // String
	year := stringToInt(dataObject[constants.MakeYear])                                                      // Number
	color := fixString(dataObject[constants.COLOR])                                                          // String
	kind := fixString(dataObject[constants.KIND])                                                            // String
	body := fixString(dataObject[constants.BODY])                                                            // String
	purpose := fixString(dataObject[constants.PURPOSE])                                                      // String

	fuelValue := dataObject[constants.FUEL]
	fixedFuelValue := fixNullableString(fuelValue)
	fuel := constants.GetFuelMapping(fixedFuelValue)

	capacity := stringToInt(dataObject[constants.CAPACITY])                           // Number, Nullable
	weight := stringToInt(dataObject[constants.OwnWeight])                            // Number, Nullable
	totalWeight := stringToInt(dataObject[constants.TotalWeight])                     // Number, Nullable
	registrationNew := fixNullableString(dataObject[constants.NumberRegistrationNew]) // String, Nullable
	vin := fixNullableString(dataObject[constants.VIN])                               // String, Nullable

	return &models.RegistrationCsvRecord{
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

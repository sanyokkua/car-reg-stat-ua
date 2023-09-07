package constants

import "strings"

// Information about KOATUU can be found here https://www.ukrstat.gov.ua/klasf/st_kls/op_koatuu_2016.htm
// or https://uk.wikipedia.org/wiki/Класифікатор_об%27єктів_адміністративно-територіального_устрою_України
// From 2020.07.17 this standard is outdated and shouldn't receive any changes, but still is used in many places
// Valid data on 2023.09.06

const KoatuuFileName = "KOATUU.csv"
const KeyKoatuuLevel1 = "ПЕРШИЙ РІВЕНЬ"
const KeyKoatuuLevel2 = "ДРУГИЙ РІВЕНЬ"
const KeyKoatuuLevel3 = "ТРЕТІЙ РІВЕНЬ"
const KeyKoatuuLevel4 = "ЧЕТВЕРТИЙ РІВЕНЬ"
const KeyKoatuuCategory = "КАТЕГОРІЯ"
const KeyKoatuuName = "НАЗВА ОБ'ЄКТА УКРАЇНСЬКОЮ МОВОЮ"

func GetAdministrativeUnitType(value string) string {
	mapping := map[string]string{
		"М": "Місто",
		"Р": "Район",
		"С": "Село",
		"Т": "Селище міського типу",
		"Щ": "Селище",
		"":  "",
	}
	trimmedValue := strings.TrimSpace(value)
	upperCase := strings.ToUpper(trimmedValue)
	newMappedValue, found := mapping[upperCase]
	if found {
		return newMappedValue
	}
	return ""
}

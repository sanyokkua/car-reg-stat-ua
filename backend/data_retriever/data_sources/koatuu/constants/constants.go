package constants

import "strings"

// Information about KOATUU can be found here https://www.ukrstat.gov.ua/klasf/st_kls/op_koatuu_2016.htm
// or https://uk.wikipedia.org/wiki/Класифікатор_об%27єктів_адміністративно-територіального_устрою_України
// From 2020.07.17 this standard is outdated and shouldn't receive any changes, but still is used in many places
// Valid data on 2023.09.06

const KOATUU_FILE_NAME = "KOATUU.csv"
const KEY_KOATUU_LEVEL_1 = "ПЕРШИЙ РІВЕНЬ"
const KEY_KOATUU_LEVEL_2 = "ДРУГИЙ РІВЕНЬ"
const KEY_KOATUU_LEVEL_3 = "ТРЕТІЙ РІВЕНЬ"
const KEY_KOATUU_LEVEL_4 = "ЧЕТВЕРТИЙ РІВЕНЬ"
const KEY_KOATUU_CATEGORY = "КАТЕГОРІЯ"
const KEY_KOATUU_NAME = "НАЗВА ОБ'ЄКТА УКРАЇНСЬКОЮ МОВОЮ"

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

package models

import (
	"data_retriever/data_sources/koatuu/constants"
	"encoding/json"
)

type KOATUUJsonRecord struct {
	Level1   string `json:"Перший рівень"`
	Level2   string `json:"Другий рівень"`
	Level3   string `json:"Третій рівень"`
	Level4   string `json:"Четвертий рівень"`
	Category string `json:"Категорія"`
	Name     string `json:"Назва об'єкта українською мовою"`
}

func (d *KOATUUJsonRecord) UnmarshalJSON(data []byte) error {
	var objMap map[string]*json.RawMessage

	err := json.Unmarshal(data, &objMap)
	if err != nil {
		return err
	}

	lev1, _ := getStringValue(objMap, constants.KeyKoatuuLevel1)
	lev2, _ := getStringValue(objMap, constants.KeyKoatuuLevel2)
	lev3, _ := getStringValue(objMap, constants.KeyKoatuuLevel3)
	lev4, _ := getStringValue(objMap, constants.KeyKoatuuLevel4)
	cat, _ := getStringValue(objMap, constants.KeyKoatuuCategory)
	name, _ := getStringValue(objMap, constants.KeyKoatuuName)

	d.Level1 = lev1
	d.Level2 = lev2
	d.Level3 = lev3
	d.Level4 = lev4
	d.Category = constants.GetAdministrativeUnitType(cat)
	d.Name = name

	return nil
}

func getStringValue(objMap map[string]*json.RawMessage, key string) (string, error) {
	var value string
	if rawMsg, ok := objMap[key]; ok {
		err := json.Unmarshal(*rawMsg, &value)
		if err != nil {
			return "", err
		}
	}
	return value, nil
}

type KOATUUJson struct {
	Records []KOATUUJsonRecord
}

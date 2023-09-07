package csv

import (
	"data_retriever/data_sources/koatuu/constants"
	"data_retriever/data_sources/koatuu/models"
)

func ToKOATUUJsonRecord(recordMap map[string]string) *models.KOATUUJsonRecord {
	lev1 := recordMap[constants.KeyKoatuuLevel1]
	lev2 := recordMap[constants.KeyKoatuuLevel2]
	lev3 := recordMap[constants.KeyKoatuuLevel3]
	lev4 := recordMap[constants.KeyKoatuuLevel4]
	cat := recordMap[constants.KeyKoatuuCategory]
	name := recordMap[constants.KeyKoatuuName]

	return &models.KOATUUJsonRecord{
		Level1:   lev1,
		Level2:   lev2,
		Level3:   lev3,
		Level4:   lev4,
		Category: cat,
		Name:     name,
	}
}

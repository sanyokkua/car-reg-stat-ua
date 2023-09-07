package json

import (
	"data_retriever/common/utils/files"
	"data_retriever/data_sources/registrations/models"
	"encoding/json"
	"errors"
	"github.com/rs/zerolog/log"
	"os"
)

func ParseDataPackageJsonFile(path string) (*models.DataPackageJson, error) {
	log.Debug().Msgf("ParseDataPackageJsonFile Path: %s", path)

	if !files.FileExist(path) {
		return nil, errors.New("file can't be opened because it doesn't exist")
	}

	// Read the file content
	data, readErr := os.ReadFile(path)
	if readErr != nil {
		log.Error().Err(readErr).Msg("Error happened during reading file content")
		return nil, readErr
	}

	// Declare a variable of type DataPackageJson for parsing json
	var dataPackage models.DataPackageJson

	// Unmarshal the JSON data into the dataPackage variable
	unmarshalErr := json.Unmarshal(data, &dataPackage)
	if unmarshalErr != nil {
		log.Error().Err(unmarshalErr).Msg("Error happened during parsing json file")
		return nil, unmarshalErr
	}

	log.Debug().Msgf("Parsed DataPackageJson json")

	return &dataPackage, nil
}

func FindUrlsOfCSVFilesInDataPackageJson(dataPackage *models.DataPackageJson) ([]string, error) {
	if dataPackage == nil || dataPackage.Resources == nil || len(dataPackage.Resources) == 0 {
		return nil, errors.New("dataPackage resources are empty")
	}

	urls := make([]string, 0, len(dataPackage.Resources))

	for i, resource := range dataPackage.Resources {
		log.Debug().Msgf("Link %d -- %s", i, resource.Path)
		urls = append(urls, resource.Path)
	}

	return urls, nil
}

package app

import (
	"data_retriever/constants"
	"data_retriever/models"
	"data_retriever/utils"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

func GetData(cleanRun bool) error {
	defer func() {
		// Global handler for panic
		if r := recover(); r != nil {
			log.Error().Msgf("Application failed to process data, %s", r)
		}
	}()

	osTempDirPath := os.TempDir()
	folderTemp := filepath.Join(osTempDirPath, constants.FOLDER_APP_TEMP)
	folderJson := filepath.Join(folderTemp, constants.FOLDER_JSON)
	folderArch := filepath.Join(folderTemp, constants.FOLDER_ARCHIVES)
	folderCsv := filepath.Join(folderTemp, constants.FOLDER_CSV)

	step1PrepareTemporaryFolder(cleanRun, folderTemp)
	step2CreateTemporaryFolders(folderTemp, folderJson, folderArch, folderCsv)
	dataPackageJsonFilePath := step3DownloadDataPackageJson(folderJson)
	parsedDataPackageJson := step4ParseDataPackageJson(dataPackageJsonFilePath)
	urlOfCsvArchives := step5RetrieveDownloadUrlsFromJsonDataPackage(parsedDataPackageJson)
	downloadedArchivesPaths := step6DownloadCsvArchives(folderArch, urlOfCsvArchives)
	extractedCsvFiles := step7ExtractAllDownloadedArchives(downloadedArchivesPaths, folderCsv)
	allRegistrationRecords := step8ParseAllCsvFilesAndGetRecords(extractedCsvFiles)

	for _, record := range allRegistrationRecords {
		log.Debug().Msg(fmt.Sprintf("%+v", record))
	}
	//err = utils.DeleteFolder(folderTemp)
	//if err != nil {
	//    return err
	//}

	log.Debug().Msg("Finished")

	return nil
}

func step1PrepareTemporaryFolder(cleanRun bool, folderTemp string) {
	log.Debug().Msgf("Step 1. If cleanRun is true (%b), temp folder (%s) will be deleted", cleanRun, folderTemp)
	if cleanRun {
		err := utils.DeleteFolder(folderTemp)
		if err != nil {
			panic(err)
		}
	}
}

func step2CreateTemporaryFolders(folderTemp string, folderJson string, folderArch string, folderCsv string) {
	log.Debug().Msgf("Step 2. Initial folders structure will be created. %s, %s, %s, %s", folderTemp, folderJson, folderArch, folderCsv)
	err := createFoldersIfNotExist(folderTemp, folderJson, folderArch, folderCsv)
	if err != nil {
		panic(err)
	}
}

func createFoldersIfNotExist(paths ...string) error {
	for _, folderPath := range paths {
		err := utils.CreateFolder(folderPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func step3DownloadDataPackageJson(folderJson string) string {
	log.Debug().Msgf("Step 3. Data Package JSON will be downloaded to: %s, from URL: %s", folderJson, constants.DATA_PACKAGE_JSON_URL)
	dataJsonPath, err := utils.Download(folderJson, constants.DATA_PACKAGE_JSON_URL)
	if err != nil {
		panic(err)
	}
	return dataJsonPath
}

func step4ParseDataPackageJson(dataJsonPath string) *models.DataPackage {
	log.Debug().Msgf("Step 4. Data Package JSON (%s) will be parsed from JSON to models.DataPackage", dataJsonPath)
	jsonFile, err := utils.ParseJsonFile(dataJsonPath)
	if err != nil {
		panic(err)
	}
	return jsonFile
}

func step5RetrieveDownloadUrlsFromJsonDataPackage(jsonFile *models.DataPackage) []string {
	log.Debug().Msgf("Step 5. URLs of CSV Archives will be retrieved from Data Package JSON %s", fmt.Sprintf("%+v", jsonFile))
	csvFiles, err := utils.FindUrlsOfCSVFiles(jsonFile)
	if err != nil {
		panic(err)
	}
	return csvFiles
}

func step6DownloadCsvArchives(folderArch string, urlOfCsvArchives []string) []string {
	log.Debug().Msgf("Step 6. CSV Archives will be downloaded to %s", folderArch)
	archives, err := utils.DownloadFilesToFolder(folderArch, urlOfCsvArchives)
	if err != nil {
		panic(err)
	}
	return archives
}

func step7ExtractAllDownloadedArchives(downloadedArchives []string, folderCsv string) []string {
	log.Debug().Msgf("Step 7. CSV Archives will be extracted to %s", folderCsv)
	extractedCsvFiles := make([]string, 0, len(downloadedArchives))

	for _, file := range downloadedArchives {
		files, err := utils.ExtractFiles(file, folderCsv)
		if err != nil {
			panic(err)
		}

		for _, filePath := range files {
			ext := filepath.Ext(filePath)
			if ext == ".csv" {
				extractedCsvFiles = append(extractedCsvFiles, filePath)
			}
		}
	}
	return extractedCsvFiles
}

func step8ParseAllCsvFilesAndGetRecords(extractedCsvFiles []string) []models.CsvRecord {
	log.Debug().Msgf("Step 8. Registration records will be parsed from all extracted CSV files")
	allRegistrationRecords := make([]models.CsvRecord, 0, 1_000_000)
	for _, filePath := range extractedCsvFiles {
		records, err := utils.ParseRegistrationsCsvToRecordsArray(filePath)
		if err != nil {
			panic(err)
		}
		log.Debug().Msgf("Number of records: %d", len(records))

		for _, record := range records {
			allRegistrationRecords = append(allRegistrationRecords, record)
		}
	}
	return allRegistrationRecords
}

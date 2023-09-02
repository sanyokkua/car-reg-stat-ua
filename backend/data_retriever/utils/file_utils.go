package utils

import (
    "archive/zip"
    "data_retriever/constants"
    "data_retriever/models"
    "encoding/json"
    "errors"
    "github.com/rs/zerolog/log"
    "io"
    "os"
    "path"
    "strings"
)

func IsFileExist(path string) bool {
    // We need to check that file exists, and currently it is the way how we can do that
    if path == "" {
        return false
    }
    info, err := os.Stat(path)
    if err == nil && info.Mode().IsRegular() {
        return true
    }
    return false
}

func CloseFunc(closableItem io.Closer) {
    // Helper function for defer keyword with logging of error if any happens
    err := closableItem.Close()
    if err != nil {
        log.Error().Err(err).Msg("Failed to close closableItem")
    }
}
func CreateFolder(path string) error {
    log.Debug().Msgf("Folder by path: %s will be created", path)

    if path == "" {
        return errors.New("path is blank, folder can't be created")
    }

    _, err := os.Stat(path)

    if os.IsNotExist(err) {
        return os.Mkdir(path, 0777)
    }

    return err
}

func DeleteFolder(path string) error {
    log.Debug().Msgf("Passed folder for deletion: %s", path)

    // We shouldn't even try to delete anything by blank path
    if path == "" {
        log.Error().Msg("Error happened. Passed empty path")
        return errors.New("path is empty")
    }

    // Get the file information for the path.
    fileInfo, err := os.Stat(path)
    if err != nil {
        log.Error().Err(err).Msgf("Error happened during getting info about folder: %s", path)
        return err
    }

    // Check if the file is a directory.
    if !fileInfo.IsDir() {
        log.Error().Msgf("Passed path %s is not a folder", path)
        return errors.New("passed path is not a folder")
    }

    // It is required to delete folder and all files in it
    remErr := os.RemoveAll(path)
    if remErr != nil {
        log.Error().Err(remErr).Msg("Error happened during removing folder")
        return remErr
    }

    return nil
}

func ParseJsonFile(path string) (*models.DataPackage, error) {
    log.Debug().Msgf("ParseJsonFile Path: %s", path)

    if !IsFileExist(path) {
        return nil, errors.New("file can't be opened because it doesn't exist")
    }

    // Read the file content
    data, readErr := os.ReadFile(path)
    if readErr != nil {
        log.Error().Err(readErr).Msg("Error happened during reading file content")
        return nil, readErr
    }

    // Declare a variable of type DataPackage for parsing json
    var dataPackage models.DataPackage

    // Unmarshal the JSON data into the dataPackage variable
    unmarshalErr := json.Unmarshal(data, &dataPackage)
    if unmarshalErr != nil {
        log.Error().Err(unmarshalErr).Msg("Error happened during parsing json file")
        return nil, unmarshalErr
    }

    log.Debug().Msgf("Parsed DataPackage json")

    return &dataPackage, nil
}

func FindUrlsOfCSVFiles(dataPackage *models.DataPackage) ([]string, error) {
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

func ExtractFiles(archivePath string, destinationDirPath string) ([]string, error) {
    log.Debug().Msgf("Archive path: %s, destinationFolder path: %s", archivePath, destinationDirPath)

    if archivePath == "" || destinationDirPath == "" {
        return nil, errors.New("file and destination folder arguments can't be blank")
    }

    if !IsFileExist(archivePath) {
        log.Error().Msgf("Archive file doesn't exist (%s)", archivePath)
        return nil, errors.New("archive file doesn't exist")
    }

    // Here we need to open ZIP archive for further processing (extracting data)
    archiveFileReader, readErr := zip.OpenReader(archivePath)
    if readErr != nil {
        log.Error().Err(readErr).Msgf("Failed to open reader for archive file: %s", archivePath)
        return nil, readErr
    }
    defer CloseFunc(archiveFileReader)

    // Now all files from archive should be extracted to the destination folder
    filePaths := make([]string, 0, len(archiveFileReader.File))
    for _, fileFromArchive := range archiveFileReader.File {
        filePath, extractionErr := extractFile(destinationDirPath, fileFromArchive)
        if extractionErr != nil {
            return nil, extractionErr
        }
        filePaths = append(filePaths, filePath)
    }

    log.Debug().Msgf("Archive %s was unpacked to %s", archivePath, destinationDirPath)
    return filePaths, nil
}

func extractFile(destinationDirPath string, fileFromArchive *zip.File) (string, error) {
    if destinationDirPath == "" {
        return "", errors.New("destination path is blank")
    }
    if fileFromArchive == nil {
        return "", errors.New("fileFromArchive is nil")
    }

    ext := strings.ToLower(path.Ext(fileFromArchive.Name))
    name := fixCsvFileName(ext, fileFromArchive.Name)
    log.Debug().Msgf("File '%s' will be extracted to folder: %s", name, destinationDirPath)

    extractionFilePath := path.Join(destinationDirPath, name)

    // Check if fileFromArchive was already created, do not extract to existing fileFromArchive
    if IsFileExist(extractionFilePath) {
        return extractionFilePath, nil
    }

    // Open file from archive for processing
    fileFromArchiveReader, openErr := fileFromArchive.Open()
    if openErr != nil {
        log.Error().Err(openErr).Msg("Error during opening fileFromArchive")
        return "", openErr
    }
    defer CloseFunc(fileFromArchiveReader)

    // Create destination file where data will be extracted
    extractedFile, extractedFileCreationErr := os.Create(extractionFilePath)
    if extractedFileCreationErr != nil {
        log.Error().Err(extractedFileCreationErr).Msg("Failed to create fileFromArchive")
        return "", extractedFileCreationErr
    }
    defer CloseFunc(extractedFile)

    // Cope data from file in zip archive to destination file in filesystem
    _, copyErr := io.Copy(extractedFile, fileFromArchiveReader)
    if copyErr != nil {
        log.Error().Err(copyErr).Msg("Failed to copy content from archive to created fileFromArchive")
        return "", copyErr
    }

    return extractionFilePath, nil
}

func fixCsvFileName(ext string, name string) string {
    // Sometime happens that packed CSV file has incorrect symbols in extension and here is required to replace it
    hasFirst2Symbols := strings.Contains(ext, "cs")
    hasLast2Symbols := strings.Contains(ext, "sv")
    hasFirstAndLastSymbols := strings.Contains(ext, "s") && strings.Contains(ext, "v")

    if hasFirst2Symbols || hasLast2Symbols || hasFirstAndLastSymbols {
        name = strings.ToValidUTF8(name, "")
        splitted := strings.Split(name, ".")[0]
        name, _ = strings.CutSuffix(splitted, ".")
        name = name + constants.CSV_FILE_EXTENSION
    }
    return name
}

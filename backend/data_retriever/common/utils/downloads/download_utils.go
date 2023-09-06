package downloads

import (
    "data_retriever/common/utils"
    "data_retriever/common/utils/files"
    "errors"
    "github.com/rs/zerolog/log"
    "io"
    "net/http"
    "os"
    "path"
)

func DownloadFile(targetDirectory string, fileUrl string) (string, error) {
    // Prepare temporary folder and targetFile for downloading
    log.Debug().Msgf("DownloadFile with params: targetDirectory=%s, fileUrl=%s", targetDirectory, fileUrl)

    fileName := path.Base(fileUrl)
    log.Debug().Msgf("FileName is: %s", fileName)

    filePath := path.Join(targetDirectory, fileName)
    log.Debug().Msgf("FilePath is: %s", filePath)

    // Check if file already exist to not spam the server
    if files.IsFileExist(filePath) {
        log.Debug().Msgf("File '%s' exists, returning path to this targetFile", filePath)
        return filePath, nil
    }

    // Create target file that will be used to save downloaded content
    targetFile, creationErr := os.Create(filePath)
    if creationErr != nil {
        log.Error().Err(creationErr).Msg("Error happened during creation a targetFile")
        return "", creationErr
    }
    defer utils.CloseFunc(targetFile)

    // DownloadFile the targetFile
    resp, downloadErr := http.Get(fileUrl)
    if downloadErr != nil {
        log.Error().Err(downloadErr).Msgf("Error happened during downloading %s", fileUrl)
        return "", downloadErr
    }
    defer utils.CloseFunc(resp.Body)

    // Copy the downloaded data to the local targetFile
    numberOfBytes, copyErr := io.Copy(targetFile, resp.Body)
    if copyErr != nil {
        log.Error().Err(copyErr).Msg("Error happened during copy of the downloaded content to destination targetFile")
        return "", copyErr
    }

    log.Debug().Msgf("Downloaded and copied %d bytes", numberOfBytes)
    log.Debug().Msgf("File was downloaded and content copied to destination targetFile: %s", filePath)

    return filePath, nil
}

func DownloadFiles(targetDirectory string, urls []string) ([]string, error) {
    if urls == nil || len(urls) == 0 {
        return nil, errors.New("no urls for downloading")
    }

    log.Debug().Msgf("Will be downloaded files to folder: %s", targetDirectory)

    downloadedFiles := make([]string, 0, len(urls))
    for _, url := range urls {
        filePath, err := DownloadFile(targetDirectory, url)
        if err != nil {
            return nil, err
        }

        downloadedFiles = append(downloadedFiles, filePath)
    }

    return downloadedFiles, nil
}

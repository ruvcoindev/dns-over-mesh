package utils

import (
    "log"
    "os"
)

var logger *log.Logger

func InitLogger(filePath string) error {
    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    logger = log.New(file, "DNS-over-MESH: ", log.LstdFlags|log.Lshortfile)
    return nil
}

func LogInfo(message string) {
    if logger != nil {
        logger.Println("INFO: " + message)
    }
}

func LogError(err error) {
    if logger != nil {
        logger.Println("ERROR: ", err)
    }
}


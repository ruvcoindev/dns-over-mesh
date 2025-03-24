package utils

import (
    "encoding/json"
    "io"
    "os"
)

func LoadConfig(filename string, config interface{}) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    return json.NewDecoder(file).Decode(config)
}

func SaveConfig(filename string, config interface{}) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    return json.NewEncoder(file).Encode(config)
}


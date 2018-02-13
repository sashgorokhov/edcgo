package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"io/ioutil"
)

type Config struct {
	Token string
}

func getConfigPath(fl *FlightLog) string {
	return filepath.Join(fl.path, "edcgo.json")
}

func loadConfig(fl *FlightLog) (Config, bool) {
	config := Config{}
	path := getConfigPath(fl)
	logger := Logger.WithField("filename", path)
	file, err := os.Open(path)
	if err != nil {
		logger.Warning("Error loading config")
		config = generateConfig()
		saveConfig(fl, config)
		logger.Infoln("Created new config")
		return config, true
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		logger.Warning("Error decoding config")
		config = generateConfig()
		saveConfig(fl, config)
		logger.Infoln("Created new config")
		return config, true
	}
	return config, false
}

func generateConfig() Config {
	return Config{Token:randToken()}
}

func saveConfig(fl *FlightLog, config Config) {
	path := getConfigPath(fl)
	logger := Logger.WithField("filename", path)
	config_json, _ := json.Marshal(config)
	err := ioutil.WriteFile(path, config_json, 0644)
	if err != nil {
		logger.Warning("Error saving config:", err)
	}
}
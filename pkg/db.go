/*
Copyright © 2025 Nevin Manoj Nevin-Manoj@ibm.com
*/
package pkg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Credentials struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type Config map[string]Credentials

const (
	dirName    = ".kbox"
	dbFileName = "kboxdb.json"
)

func SaveConfig(c Config) error {

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	//get config file path
	configFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	//update file
	return os.WriteFile(configFile, data, 0600)
}

func SaveCredentials(user, password, server string) error {

	// Load existing config if it exists
	config, err := GetConfig()
	if err != nil {
		return err
	}

	//check if we are overwriting
	_, exists := config[server]
	if exists {
		var input string
		fmt.Println("server name exists, do you wish to overwrite(y/n)")
		fmt.Scan(&input)

		if input != "y" && input != "n" {
			fmt.Println("input valid options 'y' or 'n'")
			fmt.Scan(&input)
		}

		if input != "y" && input != "n" {
			fmt.Println("Invalid option. Existing...")
			os.Exit(1)
		}

		if input != "y" {
			os.Exit(0)
		}
	}

	// Update config with new credentials
	config[server] = Credentials{
		User:     user,
		Password: password,
	}
	return SaveConfig(config)

}

func GetCredentials(server string) (*Credentials, error) {

	config, err := GetConfig()
	if err != nil {
		return nil, err
	}
	// Lookup credentials for the server
	creds, found := config[server]
	if !found {
		return nil, fmt.Errorf("no credentials found for server: %s", server)
	}

	return &creds, nil
}

func GetConfig() (Config, error) {
	//gets config if exists

	configFile, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	config := Config{}
	if data, err := os.ReadFile(configFile); err == nil {
		json.Unmarshal(data, &config)
	}

	return config, nil
}

func GetServerNames() ([]string, error) {
	config, err := GetConfig()
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(config))
	for k := range config {
		keys = append(keys, k)
	}
	return keys, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(homeDir, dirName)

	if err := os.MkdirAll(configDir, 0700); err != nil {
		return "", err
	}
	configFile := filepath.Join(configDir, dbFileName)

	return configFile, nil
}

func RemoveServer(server string) error {
	config, err := GetConfig()
	if err != nil {
		return err
	}
	delete(config, server)
	return SaveConfig(config)
}

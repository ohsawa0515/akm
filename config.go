package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type AkmConfig struct {
	Current string `toml:"current"`
}

func getAkmConfigPath() string {
	if file, ok := os.LookupEnv("AKM_FILE"); ok {
		return file
	}

	return filepath.Join(getHomeDir(), ".akm.toml")
}

func createAkmConfig() error {
	akmConfigPath := getAkmConfigPath()

	if _, err := os.Stat(akmConfigPath); err == nil {
		return fmt.Errorf("%s already exists", akmConfigPath)
	}

	f, err := os.OpenFile(akmConfigPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	f.WriteString("current = \"default\"\n")
	defer f.Close()

	return nil
}

func NewAkmConfig() (*AkmConfig, error) {
	akmConfig := &AkmConfig{}
	akmConfigPath := getAkmConfigPath()

	if err := createAkmConfig(); err != nil {
		return nil, err
	}

	if _, err := toml.DecodeFile(akmConfigPath, akmConfig); err != nil {
		return nil, err
	}

	return akmConfig, nil
}

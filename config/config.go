package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/ohsawa0515/akm/utils"
)

type AkmConfig struct {
	Current string `toml:"current"`
}

func getAkmHomeDir() string {
	if dir, ok := os.LookupEnv("AKM_CONFIG_DIR"); ok {
		return dir
	}

	return filepath.Join(utils.GetHomeDir(), ".akm")
}

func getAkmConfigPath() string {
	if file, ok := os.LookupEnv("AKM_CONFIG_FILE"); ok {
		return file
	}

	return filepath.Join(getAkmHomeDir(), "config")
}

func CreateAkmConfig() error {
	akmConfigPath := getAkmConfigPath()

	if _, err := os.Stat(akmConfigPath); err == nil {
		return fmt.Errorf("%s already exists", akmConfigPath)
	}

	if err := os.Mkdir(getAkmHomeDir(), 0755); err != nil {
		return err
	}

	f, err := os.OpenFile(akmConfigPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

func NewAkmConfig() (*AkmConfig, error) {
	akmConfig := &AkmConfig{}
	akmConfigPath := getAkmConfigPath()

	if _, err := toml.DecodeFile(akmConfigPath, akmConfig); err != nil {
		return nil, err
	}

	return akmConfig, nil
}

func (akmConfig *AkmConfig) Save() error {
	var buf bytes.Buffer

	if err := toml.NewEncoder(&buf).Encode(akmConfig); err != nil {
		return err
	}

	if err := ioutil.WriteFile(getAkmConfigPath(), buf.Bytes(), 0); err != nil {
		return err
	}
	defer buf.Reset()

	return nil
}

func (akmConfig *AkmConfig) Delete() error {
	akmConfig.Current = ""
	if err := akmConfig.Save(); err != nil {
		return err
	}

	return nil
}

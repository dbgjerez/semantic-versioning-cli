package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type ConfigStore struct {
	Path string
}

func NewConfigStore(path string) ConfigStore {
	return ConfigStore{Path: path}
}

func (store *ConfigStore) Exists() bool {
	if _, err := os.Stat(store.Path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func (store *ConfigStore) ReadConfig() (Config, error) {
	f, err := ioutil.ReadFile("./" + store.Path)
	if err != nil {
		return Config{}, errors.New(fmt.Sprintf("Error reading the file %s: %v", store.Path, err))
	}

	var config Config
	err = json.Unmarshal(f, &config)
	if err != nil {
		return Config{}, errors.New(fmt.Sprintf("Error unmarshal the %s content: %v", store.Path, err))
	}
	return config, nil
}

func (store *ConfigStore) SaveConfig(config Config) error {
	file, _ := json.MarshalIndent(config, "", "  ")
	err := ioutil.WriteFile(store.Path, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

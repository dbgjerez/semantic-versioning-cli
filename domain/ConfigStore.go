package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type ConfigStore struct {
	Path string
}

func NewConfigStore(path string) ConfigStore {
	return ConfigStore{Path: path}
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

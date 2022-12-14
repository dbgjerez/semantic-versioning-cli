package domain

import (
	"testing"
)

const (
	PATH string = "test"
)

func TestNewConfigStore(t *testing.T) {
	want := ConfigStore{Path: "./" + PATH}
	got := NewConfigStore(PATH)

	if want != got {
		t.Errorf("Expected path %s but got %s", PATH, got.Path)
	}
}

func TestExistsFalse(t *testing.T) {
	configStore := ConfigStore{Path: "./" + PATH}
	want := false
	got := configStore.Exists()

	if want != got {
		t.Errorf("Expected %t but got %t ", want, got)
	}
}

func TestReadConfigError(t *testing.T) {
	configStore := ConfigStore{Path: "./" + PATH}
	_, err := configStore.ReadConfig()
	if err == nil {
		t.Errorf("Expected error")
	}
}

func TestSaveConfigError(t *testing.T) {
	configStore := ConfigStore{}
	config := Store{}
	err := configStore.SaveConfig(config)
	if err == nil {
		t.Errorf("Expected error")
	}
}

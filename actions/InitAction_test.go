package actions

import (
	"semver/domain"
	"testing"
)

func TestNewConfig(t *testing.T) {
	want := domain.Config{
		Data: domain.DataConfig{
			ArtifactName: ArtifactName,
			Version: domain.VersionConfig{
				Major: Major,
				Minor: Minor,
				Patch: Patch,
			},
		},
	}
	init := InitAction{
		ArtifactName: ArtifactName,
		Major:        Major,
		Minor:        Minor,
		Patch:        Patch,
	}
	got, err := init.NewConfig()
	if err != nil {
		t.Errorf("Don't expected fails!")
	}
	if want != got {
		t.Errorf("The config gotten %v is different than the desired config: %v", got, want)
	}
}

package actions

import (
	"semver/domain"
)

type InitAction struct {
	ArtifactName string
	Major        int
	Minor        int
	Patch        int
}

func (init *InitAction) NewConfig() (domain.Config, error) {
	return domain.Config{
		Data: domain.DataConfig{
			ArtifactName: init.ArtifactName,
			Version: domain.VersionConfig{
				Major: init.Major,
				Minor: init.Minor,
				Patch: init.Patch,
			},
		},
	}, nil
}

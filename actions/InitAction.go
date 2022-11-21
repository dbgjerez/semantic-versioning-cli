package actions

import (
	"semver/domain"
)

type InitAction struct {
	ArtifactName    string
	Major           int
	Minor           int
	Patch           int
	SnapshotsEnable bool
	SnapshotsKey    string
}

const (
	INIT_MAJOR_VERSION     = 0
	INIT_MINOR_VERSION     = 0
	INIT_PATCH_VERSION     = 0
	INIT_SNAPSHOTS_ENABLED = true
	INIT_SNAPSHOTS_KEY     = "SNAPSHOT"
)

func (init *InitAction) NewConfig() (domain.Config, error) {
	return domain.Config{
		Config: domain.SemverConfig{
			Versions: domain.SemverConfigVersions{
				Snapshot: domain.SemverConfigSnapshots{
					Enabled: init.SnapshotsEnable,
					Key:     init.SnapshotsKey,
				},
			},
		},
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

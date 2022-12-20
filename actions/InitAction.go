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
	INIT_MAJOR_VERSION     int    = 0
	INIT_MINOR_VERSION     int    = 0
	INIT_PATCH_VERSION     int    = 0
	INIT_SNAPSHOTS_ENABLED bool   = true
	INIT_SNAPSHOTS_KEY     string = "SNAPSHOT"
)

func (init *InitAction) NewConfig() (domain.Store, error) {
	if init.SnapshotsKey == "" {
		init.SnapshotsKey = INIT_SNAPSHOTS_KEY
	}
	store := domain.Store{
		Config: domain.SemverConfig{
			Snapshots: domain.SemverSubType{
				Key:     init.SnapshotsKey,
				Enabled: init.SnapshotsEnable,
			},
		},
		Data: domain.DataStore{
			ArtifactName: init.ArtifactName,
			Version: domain.VersionConfig{
				Major: init.Major,
				Minor: init.Minor,
				Patch: init.Patch,
			},
		},
	}
	if store.IsSnapshotEnabled() {
		store.Data.Version.Snapshot = true
	}
	return store, nil
}

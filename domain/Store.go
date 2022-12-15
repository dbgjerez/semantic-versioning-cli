package domain

type Store struct {
	Data   DataStore    `json:"data"`
	Config SemverConfig `json:"config"`
}

type SemverConfig struct {
	Snapshots         SemverSubType       `json:"snapshots"`
	ReleaseCandidates SemverSubType       `json:"release-candidates"`
	GitFlow           SemverGitflowConfig `json:"gitflow"`
}

type SemverSubType struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"key"`
}

type SemverGitflowConfig struct {
	Enabled  bool            `json:"enabled"`
	Branches GitFlowBranches `json:"branches"`
}

type GitFlowBranches struct {
	Snapshots         []string `json:"snapshots"`
	ReleaseCandidates []string `json:"release-candidates"`
}

type DataStore struct {
	ArtifactName string        `json:"artifactName"`
	Version      VersionConfig `json:"version"`
}

type VersionConfig struct {
	Major    int  `json:"major"`
	Minor    int  `json:"minor"`
	Patch    int  `json:"patch"`
	RC       int  `json:"rc"`
	Snapshot bool `json:"snapshot"`
}

func (store Store) IsSnapshotEnabled() bool {
	return store.Config.Snapshots.Enabled
}

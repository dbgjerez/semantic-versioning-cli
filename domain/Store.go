package domain

type Store struct {
	Data   DataStore    `json:"data"`
	Config SemverConfig `json:"config"`
}

type SemverConfig struct {
	Versions SemverConfigVersions `json:"versions"`
}

type SemverConfigVersions struct {
	Snapshot SemverConfigSnapshots `json:"snapshots"`
}

type SemverConfigSnapshots struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"key"`
}

type DataStore struct {
	ArtifactName string        `json:"artifactName"`
	Version      VersionConfig `json:"version"`
}

type VersionConfig struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

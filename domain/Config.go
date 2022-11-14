package domain

type Config struct {
	Data DataConfig `json:"data"`
}

type DataConfig struct {
	ArtifactName string        `json:"artifactName"`
	Version      VersionConfig `json:"version"`
}

type VersionConfig struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Patch int `json:"patch"`
}

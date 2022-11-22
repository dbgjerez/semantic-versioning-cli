package actions

import (
	"fmt"
	"semver/domain"
)

type InfoAction struct {
	c *domain.Store
}

func NewInfoAction(c *domain.Store) InfoAction {
	return InfoAction{c: c}
}

func (info *InfoAction) CompleteInfo() string {
	return fmt.Sprintf("Artifact name: %s\n"+
		"Version: %s\n", info.ArtifactName(), info.ArtifactVersion())
}

func (info *InfoAction) ArtifactName() string {
	return info.c.Data.ArtifactName
}

func (info *InfoAction) ArtifactVersion() string {
	version := fmt.Sprintf("%d.%d", info.c.Data.Version.Major, info.c.Data.Version.Minor)
	if info.c.Data.Version.Patch != 0 {
		version += fmt.Sprintf(".%d", info.c.Data.Version.Patch)
	}
	if info.c.Config.Versions.Snapshot.Enabled {
		version += fmt.Sprintf("-%s", info.c.Config.Versions.Snapshot.Key)
	}
	return version
}

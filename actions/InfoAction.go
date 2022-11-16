package actions

import (
	"fmt"
	"semver/domain"
)

type InfoAction struct {
	c *domain.Config
}

func NewInfoAction(c *domain.Config) InfoAction {
	return InfoAction{c: c}
}

func (info *InfoAction) ArtifactName() string {
	return info.c.Data.ArtifactName
}

func (info *InfoAction) ArtifactVersion() string {
	version := fmt.Sprintf("%d.%d", info.c.Data.Version.Major, info.c.Data.Version.Minor)
	if info.c.Data.Version.Patch != 0 {
		version += fmt.Sprintf(".%d", info.c.Data.Version.Patch)
	}
	return version
}

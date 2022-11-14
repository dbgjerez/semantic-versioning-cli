package actions

import (
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

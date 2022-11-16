package actions

import (
	"errors"
	"fmt"
	"semver/domain"
)

type ReleaseAction struct {
	Config *domain.Config
}

func NewReleaseAction(config domain.Config) ReleaseAction {
	return ReleaseAction{Config: &config}
}

func (action *ReleaseAction) CreateRelease(majorVersion int) (domain.Config, error) {
	if majorVersion < -1 {
		return *action.Config, errors.New(
			fmt.Sprintf("Invalid version %d", majorVersion))
	} else if majorVersion == -1 {
		action.Config.Data.Version.Major += 1
	} else if majorVersion == action.Config.Data.Version.Major {
		return *action.Config, errors.New(
			fmt.Sprintf("The actual major version is already %d", majorVersion))
	} else {
		action.Config.Data.Version.Major = majorVersion
	}
	action.Config.Data.Version.Minor = 0
	action.Config.Data.Version.Patch = 0
	return *action.Config, nil
}

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

func (action *ReleaseAction) CreateMajor(majorVersion int) (domain.Config, error) {
	if majorVersion == -1 {
		action.Config.Data.Version.Major += 1
	} else if majorVersion < action.Config.Data.Version.Major {
		return *action.Config, errors.New(
			fmt.Sprintf("You cannot down the version from %d to %d",
				action.Config.Data.Version.Major, majorVersion))
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

func (action *ReleaseAction) CreateFeature(minorVersion int) (domain.Config, error) {
	if minorVersion < -1 {
		return *action.Config, errors.New(
			fmt.Sprintf("Invalid version %d", minorVersion))
	} else if minorVersion == -1 {
		action.Config.Data.Version.Minor += 1
	} else if minorVersion == action.Config.Data.Version.Minor {
		return *action.Config, errors.New(
			fmt.Sprintf("The actual minor version is already %d", minorVersion))
	} else {
		action.Config.Data.Version.Minor = minorVersion
	}
	action.Config.Data.Version.Patch = 0
	return *action.Config, nil
}

func (action *ReleaseAction) CreatePatch(patchVersion int) (domain.Config, error) {
	if patchVersion < -1 {
		return *action.Config, errors.New(
			fmt.Sprintf("Invalid version %d", patchVersion))
	} else if patchVersion == -1 {
		action.Config.Data.Version.Patch += 1
	} else if patchVersion == action.Config.Data.Version.Patch {
		return *action.Config, errors.New(
			fmt.Sprintf("The actual patch version is already %d", patchVersion))
	} else {
		action.Config.Data.Version.Patch = patchVersion
	}
	return *action.Config, nil
}

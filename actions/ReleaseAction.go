package actions

import (
	"errors"
	"fmt"
	"semver/domain"
)

type ReleaseAction struct {
	Store *domain.Store
}

func NewReleaseAction(store *domain.Store) ReleaseAction {
	return ReleaseAction{Store: store}
}

func (action *ReleaseAction) CreateMajor(majorVersion int) (domain.Store, error) {
	if majorVersion == -1 {
		action.Store.Data.Version.Major += 1
	} else if majorVersion < action.Store.Data.Version.Major {
		return *action.Store, errors.New(
			fmt.Sprintf("You cannot down the version from %d to %d",
				action.Store.Data.Version.Major, majorVersion))
	} else if majorVersion == action.Store.Data.Version.Major {
		return *action.Store, errors.New(
			fmt.Sprintf("The actual major version is already %d", majorVersion))
	} else {
		action.Store.Data.Version.Major = majorVersion
	}
	action.Store.Data.Version.Minor = 0
	action.Store.Data.Version.Patch = 0
	if action.Store.IsSnapshotEnabled() {
		action.Store.Data.Version.Snapshot = true
	}
	return *action.Store, nil
}

func (action *ReleaseAction) CreateFeature(minorVersion int) (domain.Store, error) {
	if minorVersion == -1 {
		action.Store.Data.Version.Minor += 1
	} else if minorVersion == action.Store.Data.Version.Minor {
		return *action.Store, errors.New(
			fmt.Sprintf("The actual minor version is already %d", minorVersion))
	} else if minorVersion < action.Store.Data.Version.Minor {
		return *action.Store, errors.New(
			fmt.Sprintf("You cannot down the version %d for %d",
				action.Store.Data.Version.Minor, minorVersion))
	} else {
		action.Store.Data.Version.Minor = minorVersion
	}
	action.Store.Data.Version.Patch = 0
	return *action.Store, nil
}

func (action *ReleaseAction) CreatePatch(patchVersion int) (domain.Store, error) {
	if patchVersion == -1 {
		action.Store.Data.Version.Patch += 1
	} else if patchVersion < action.Store.Data.Version.Patch {
		return *action.Store, errors.New(
			fmt.Sprintf("The actual patch versoin (%d) is grather than (%d)",
				action.Store.Data.Version.Patch, patchVersion))
	} else if patchVersion == action.Store.Data.Version.Patch {
		return *action.Store, errors.New(
			fmt.Sprintf("The actual patch version is already %d", patchVersion))
	} else {
		action.Store.Data.Version.Patch = patchVersion
	}
	return *action.Store, nil
}

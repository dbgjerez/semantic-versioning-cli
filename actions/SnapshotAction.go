package actions

import (
	"errors"
	"semver/domain"
)

var ErrSnapshotsNotEnable = errors.New("Snapshots are not enabled")

type SnapshotAction struct {
	C           *domain.Store
	Force       bool // Indicates if the value is forced
	ForcedValue bool // Indicates the value
}

func (action *SnapshotAction) IsSnapshot() bool {
	if action.C.Config.Snapshots.Enabled &&
		action.C.Data.Version.Snapshot {
		return true
	}
	return false
}

func (action *SnapshotAction) ChangeStatus() error {
	if !action.C.Config.Snapshots.Enabled {
		return ErrSnapshotsNotEnable
	} else {
		if action.Force {
			action.C.Data.Version.Snapshot = action.ForcedValue
		} else {
			if action.IsSnapshot() {
				action.C.Data.Version.Snapshot = false
			} else {
				action.C.Data.Version.Snapshot = true
			}
		}
	}
	return nil
}

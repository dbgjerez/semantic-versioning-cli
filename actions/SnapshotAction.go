package actions

import (
	"errors"
	"semver/domain"
)

var ErrSnapshotsNotEnable = errors.New("Snapshots are not enabled")

type SnapshotAction struct {
	c *domain.Store
}

func (action *SnapshotAction) IsSnapshot() bool {
	if action.c.Config.Snapshots.Enabled &&
		action.c.Data.Version.Snapshot {
		return true
	}
	return false
}

func (action *SnapshotAction) ChangeStatus() error {
	if !action.c.Config.Snapshots.Enabled {
		return ErrSnapshotsNotEnable
	} else {
		if action.IsSnapshot() {
			action.c.Data.Version.Snapshot = false
		} else {
			action.c.Data.Version.Snapshot = true
		}
	}
	return nil
}

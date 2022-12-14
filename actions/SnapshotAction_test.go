package actions

import (
	"errors"
	"reflect"
	"semver/domain"
	"testing"
)

type Tests struct {
	name          string
	store         domain.Store
	expectedError error
	force         bool
}

func TestSnapshotAction(t *testing.T) {
	tests := []Tests{
		{
			name: "snapshots-disables",
			store: domain.Store{
				Config: domain.SemverConfig{
					Snapshots: domain.SemverSubType{
						Enabled: false,
					},
				},
			},
			expectedError: ErrSnapshotsNotEnable,
		},
		{
			name: "snapshots-true",
			store: domain.Store{
				Config: domain.SemverConfig{
					Snapshots: domain.SemverSubType{
						Enabled: true,
					},
				},
				Data: domain.DataStore{
					Version: domain.VersionConfig{
						Snapshot: true,
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "snapshots-false",
			store: domain.Store{
				Config: domain.SemverConfig{
					Snapshots: domain.SemverSubType{
						Enabled: true,
					},
				},
				Data: domain.DataStore{
					Version: domain.VersionConfig{
						Snapshot: false,
					},
				},
			},
			expectedError: nil,
		},
		{
			name: "snapshots-force-true",
			store: domain.Store{
				Config: domain.SemverConfig{
					Snapshots: domain.SemverSubType{
						Enabled: true,
					},
				},
				Data: domain.DataStore{
					Version: domain.VersionConfig{
						Snapshot: false,
					},
				},
			},
			force:         true,
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			action := SnapshotAction{C: &test.store, Force: test.force, ForcedValue: true}
			previous := test.store.Data.Version.Snapshot
			err := action.ChangeStatus()

			if test.expectedError != nil {
				if !errors.Is(err, ErrSnapshotsNotEnable) {
					t.Errorf("Expected error FAILED: expected [%v] got [%v]", test.expectedError, err)
				}
			} else {
				if reflect.DeepEqual(action.C.Data.Version.Snapshot, previous) {
					t.Errorf("Expected valuer FAILED: expected [%t] got [%t]",
						action.C.Data.Version.Snapshot, previous)
				}
			}
		})
	}
}

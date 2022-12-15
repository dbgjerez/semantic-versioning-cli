package actions

import (
	"reflect"
	"semver/domain"
	"testing"
)

func TestNewReleaseAction(t *testing.T) {
	c := NewConfigMock(true)
	want := ReleaseAction{Store: c}
	got := NewReleaseAction(c)

	if want != got {
		t.Errorf("Expected ReleaseAction is different")
	}
}

type StoreHelper struct {
	SnapshotEnabled bool
	Major           int
	Minor           int
	Patch           int
	Snapshot        bool
}

func (storeHelper StoreHelper) NewStoreConfig() domain.Store {
	return domain.Store{
		Data: domain.DataStore{
			Version: domain.VersionConfig{
				Major:    storeHelper.Major,
				Minor:    storeHelper.Minor,
				Patch:    storeHelper.Patch,
				Snapshot: storeHelper.Snapshot,
			},
		},
		Config: domain.SemverConfig{
			Snapshots: domain.SemverSubType{
				Enabled: storeHelper.SnapshotEnabled,
				Key:     INIT_SNAPSHOTS_KEY,
			},
		},
	}
}

type Test struct {
	name           string
	config         domain.Store
	param          int
	expectedError  bool
	expectedResult string
}

func TestCreateMajor(t *testing.T) {
	// Received -1 => Default case (sum 1)
	// Received == actual version ==> Error
	// Received > actual version ==> Force version
	// Received < actual version ==> Error
	tests := []Test{
		{
			name: "update major without snapshot",
			config: StoreHelper{
				SnapshotEnabled: false,
				Major:           1,
				Minor:           1,
				Patch:           0,
			}.NewStoreConfig(),
			param:          -1,
			expectedError:  false,
			expectedResult: "2.0",
		},
		{
			name: "update major with same version",
			config: StoreHelper{
				SnapshotEnabled: false,
				Major:           1,
				Minor:           1,
				Patch:           1,
			}.NewStoreConfig(),
			param:         1,
			expectedError: true,
		},
		{
			name: "update major with one more version",
			config: StoreHelper{
				SnapshotEnabled: false,
				Major:           1,
				Minor:           1,
				Patch:           1,
			}.NewStoreConfig(),
			param:          2,
			expectedError:  false,
			expectedResult: "2.0",
		},
		{
			name: "update major forcing major version without snapshot",
			config: StoreHelper{
				SnapshotEnabled: false,
				Major:           1,
				Minor:           1,
				Patch:           1,
			}.NewStoreConfig(),
			param:          4,
			expectedError:  false,
			expectedResult: "4.0",
		},
		{
			name: "update major forcing minus version value without snapshot",
			config: StoreHelper{
				SnapshotEnabled: false,
				Major:           4,
				Minor:           1,
				Patch:           1,
			}.NewStoreConfig(),
			param:         2,
			expectedError: true,
		},
		{
			name: "update major with snapshot enabled",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           1,
				Minor:           1,
				Patch:           0,
				Snapshot:        false,
			}.NewStoreConfig(),
			param:          2,
			expectedError:  false,
			expectedResult: "2.0-SNAPSHOT",
		},
		{
			name: "update major with snapshot enabled and snapshot",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           1,
				Minor:           1,
				Patch:           0,
				Snapshot:        true,
			}.NewStoreConfig(),
			param:          2,
			expectedError:  false,
			expectedResult: "2.0-SNAPSHOT",
		},
		{
			name: "update major with snapshot enabled and same version",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           1,
				Minor:           1,
				Patch:           0,
				Snapshot:        true,
			}.NewStoreConfig(),
			param:         1,
			expectedError: true,
		},
		{
			name: "update major with correct version and snapshot enabled",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           1,
				Minor:           1,
				Patch:           1,
				Snapshot:        true,
			}.NewStoreConfig(),
			param:          2,
			expectedError:  false,
			expectedResult: "2.0-SNAPSHOT",
		},
		{
			name: "update major with correct version and snapshot enabled without snapshot",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           1,
				Minor:           1,
				Patch:           1,
				Snapshot:        false,
			}.NewStoreConfig(),
			param:          2,
			expectedError:  false,
			expectedResult: "2.0-SNAPSHOT",
		},
		{
			name: "update major with major version and snapshot enabled without snapshot",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           1,
				Minor:           1,
				Patch:           1,
				Snapshot:        false,
			}.NewStoreConfig(),
			param:          4,
			expectedError:  false,
			expectedResult: "4.0-SNAPSHOT",
		},
		{
			name: "update major with major version and snapshot enabled with snapshot",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           1,
				Minor:           1,
				Patch:           1,
				Snapshot:        true,
			}.NewStoreConfig(),
			param:          4,
			expectedError:  false,
			expectedResult: "4.0-SNAPSHOT",
		},
		{
			name: "update major with invalid version and snapshot enabled with snapshot",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           4,
				Minor:           1,
				Patch:           1,
				Snapshot:        true,
			}.NewStoreConfig(),
			param:         2,
			expectedError: true,
		},
		{
			name: "update major with major version and snapshot enabled without snapshot",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           4,
				Minor:           1,
				Patch:           1,
				Snapshot:        false,
			}.NewStoreConfig(),
			param:         2,
			expectedError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			action := NewReleaseAction(&test.config)
			infoAction := NewInfoAction(&test.config)
			config, e := action.CreateMajor(test.param)
			resInfoAction := NewInfoAction(&config)
			if e == nil &&
				test.expectedError {
				t.Errorf("Expected error with version %s and param %d, but got a version %s",
					infoAction.ArtifactVersion(),
					test.param,
					resInfoAction.ArtifactVersion())
			} else if e != nil &&
				!test.expectedError {
				t.Errorf("Unexpected error %v with version %s and param %d",
					e,
					infoAction.ArtifactVersion(),
					test.param)
			} else if !test.expectedError &&
				!reflect.DeepEqual(resInfoAction.ArtifactVersion(), test.expectedResult) {
				t.Errorf("Expected version %s with param %d and got version %s",
					test.expectedResult,
					test.param,
					resInfoAction.ArtifactVersion())
			}
		})
	}
}

func TestCreateFeature(t *testing.T) {
	// Received -1 => Default case (sum 1)
	// Received == actual version ==> Error
	// Received > actual version ==> Force version
	// Received < actual version ==> Error
	tests := []Test{
		{
			name: "[1.1 - 1.2] ==> normal case",
			config: StoreHelper{
				SnapshotEnabled: false,
				Major:           1,
				Minor:           1,
				Patch:           0,
			}.NewStoreConfig(),
			param:          -1,
			expectedError:  false,
			expectedResult: "1.2",
		},
		{
			name: "[1.1.1 - 1.2] ==> normal case",
			config: StoreHelper{
				SnapshotEnabled: false,
				Major:           1,
				Minor:           1,
				Patch:           1,
			}.NewStoreConfig(),
			param:          -1,
			expectedResult: "1.2",
		},
		{
			name: "[1.1.1 - 1.1] ==> error with forced 1",
			config: StoreHelper{
				SnapshotEnabled: false,
				Major:           1,
				Minor:           1,
				Patch:           1,
			}.NewStoreConfig(),
			param:         1,
			expectedError: true,
		},
		{
			name: "[1.1.1 - 1.2] ==> forced 2",
			config: StoreHelper{
				SnapshotEnabled: false,
				Major:           1,
				Minor:           1,
				Patch:           1,
			}.NewStoreConfig(),
			param:          2,
			expectedError:  false,
			expectedResult: "1.2",
		},
		{
			name: "[1.1.1 - 1.4] ==> forced 4",
			config: StoreHelper{
				SnapshotEnabled: false,
				Major:           1,
				Minor:           1,
				Patch:           1,
			}.NewStoreConfig(),
			param:          4,
			expectedResult: "1.4",
		},
		{
			name: "[1.4.1 - 1.2] ==> error forced 2",
			config: StoreHelper{
				SnapshotEnabled: false,
				Major:           1,
				Minor:           4,
				Patch:           1,
			}.NewStoreConfig(),
			param:         2,
			expectedError: true,
		},
		{
			name: "[1.1 - 1.2-SNAPSHOT] ==> normal case with snapshot",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           1,
				Minor:           1,
				Patch:           0,
				Snapshot:        false,
			}.NewStoreConfig(),
			param:          -1,
			expectedError:  false,
			expectedResult: "1.2-SNAPSHOT",
		},
		{
			name: "[1.1.1-SNAPSHOT - 1.2-SNAPSHOT] ==> normal case with snapshot",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           1,
				Minor:           1,
				Patch:           1,
				Snapshot:        true,
			}.NewStoreConfig(),
			param:          -1,
			expectedError:  false,
			expectedResult: "1.2-SNAPSHOT",
		},
		{
			name: "[1.1.1 - 1.2-SNAPSHOT] ==> normal case with snapshot",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           1,
				Minor:           1,
				Patch:           1,
				Snapshot:        false,
			}.NewStoreConfig(),
			param:          -1,
			expectedError:  false,
			expectedResult: "1.2-SNAPSHOT",
		},
		{
			name: "[1.1.1-SNAPSHOT - 1.1] ==> error with same version",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           1,
				Minor:           1,
				Patch:           1,
				Snapshot:        true,
			}.NewStoreConfig(),
			param:         1,
			expectedError: true,
		},
		{
			name: "[1.1.1 - 1.4-SNAPSHOT] ==> forced 1.4-SNAPSHOT",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           1,
				Minor:           1,
				Patch:           1,
				Snapshot:        false,
			}.NewStoreConfig(),
			param:          4,
			expectedError:  false,
			expectedResult: "1.4-SNAPSHOT",
		},
		{
			name: "[1.4.1 - 1.2] ==> error forced feature 2",
			config: StoreHelper{
				SnapshotEnabled: true,
				Major:           1,
				Minor:           4,
				Patch:           1,
				Snapshot:        false,
			}.NewStoreConfig(),
			param:         2,
			expectedError: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			action := NewReleaseAction(&test.config)
			infoAction := NewInfoAction(&test.config)
			initV := infoAction.ArtifactVersion()
			config, e := action.CreateFeature(test.param)
			resInfoAction := NewInfoAction(&config)
			if e == nil && test.expectedError {
				t.Errorf("Expected error with version %s and param %d, but got a version %s",
					initV,
					test.param,
					resInfoAction.ArtifactVersion())
			} else if e != nil && !test.expectedError {
				t.Errorf("Unexpected error %v with version %s and param %d",
					e,
					initV,
					test.param)
			} else if !test.expectedError && resInfoAction.ArtifactVersion() != test.expectedResult {
				t.Errorf("Expected version %s with param %d and got version %s",
					test.expectedResult,
					test.param,
					resInfoAction.ArtifactVersion())
			}
		})
	}
}

func TestCreatePatch(t *testing.T) {
	// Received -1 => Default case (sum 1)
	// Received == actual version ==> Error
	// Received > actual version ==> Force version
	// Received < actual version ==> Error
	type WantTest struct {
		err     bool
		version string
	}
	type VersionTest struct {
		v     domain.VersionConfig
		param int
		want  WantTest
	}
	versions := []VersionTest{
		{
			v:     NewVersionMock(1, 1, 0, false),
			param: -1,
			want:  WantTest{err: false, version: "1.1.1"},
		},
		{
			v:     NewVersionMock(1, 1, 1, false),
			param: -1,
			want:  WantTest{err: false, version: "1.1.2"},
		},
		{
			v:     NewVersionMock(1, 1, 1, false),
			param: 1,
			want:  WantTest{err: true},
		},
		{
			v:     NewVersionMock(1, 1, 1, false),
			param: 2,
			want:  WantTest{err: false, version: "1.1.2"},
		},
		{
			v:     NewVersionMock(1, 1, 1, false),
			param: 4,
			want:  WantTest{err: false, version: "1.1.4"},
		},
		{
			v:     NewVersionMock(1, 1, 4, true),
			param: 2,
			want:  WantTest{err: true},
		},
		{
			v:     NewVersionMock(1, 1, 0, true),
			param: -1,
			want:  WantTest{err: false, version: "1.1.1-SNAPSHOT"},
		},
		{
			v:     NewVersionMock(1, 1, 1, true),
			param: -1,
			want:  WantTest{err: false, version: "1.1.2-SNAPSHOT"},
		},
		{
			v:     NewVersionMock(1, 1, 1, true),
			param: 1,
			want:  WantTest{err: true},
		},
		{
			v:     NewVersionMock(1, 1, 1, true),
			param: 2,
			want:  WantTest{err: false, version: "1.1.2-SNAPSHOT"},
		},
		{
			v:     NewVersionMock(1, 1, 1, true),
			param: 4,
			want:  WantTest{err: false, version: "1.1.4-SNAPSHOT"},
		},
		{
			v:     NewVersionMock(1, 1, 4, true),
			param: 2,
			want:  WantTest{err: true},
		},
	}
	for _, v := range versions {
		c := NewConfigMockWithVersion(v.v)
		action := NewReleaseAction(c)
		infoAction := NewInfoAction(c)
		initV := infoAction.ArtifactVersion()
		config, e := action.CreatePatch(v.param)
		resInfoAction := NewInfoAction(&config)
		if e == nil && v.want.err {
			t.Errorf("Expected error with version %s and param %d, but got a version %s",
				initV,
				v.param,
				resInfoAction.ArtifactVersion())
		} else if e != nil && !v.want.err {
			t.Errorf("Unexpected error %v with version %s and param %d",
				e,
				initV,
				v.param)
		} else if !v.want.err && resInfoAction.ArtifactVersion() != v.want.version {
			t.Errorf("Expected version %s with param %d and got version %s",
				v.want.version,
				v.param,
				resInfoAction.ArtifactVersion())
		}
	}
}

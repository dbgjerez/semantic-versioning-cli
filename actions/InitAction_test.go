package actions

import (
	"testing"
)

func TestNewConfig(t *testing.T) {
	useCases := []InitAction{
		{
			ArtifactName:    ArtifactName,
			Major:           Major,
			Minor:           Minor,
			Patch:           Patch,
			SnapshotsEnable: true,
			SnapshotsKey:    INIT_SNAPSHOTS_KEY,
		},
		{
			ArtifactName:    ArtifactName,
			Major:           Major,
			Minor:           Minor,
			Patch:           Patch,
			SnapshotsEnable: true,
			SnapshotsKey:    "",
		},
		{
			ArtifactName: ArtifactName,
			Major:        Major,
			Minor:        Minor,
			Patch:        Patch,
			SnapshotsKey: INIT_SNAPSHOTS_KEY,
		},
	}
	for _, uc := range useCases {
		got, err := uc.NewConfig()
		if err != nil {
			t.Errorf("Don't expected fails!")
		}
		if uc.SnapshotsEnable != got.Config.Snapshots.Enabled {
			t.Errorf("Want SnapshotsEnable %t and got %t", uc.SnapshotsEnable, got.Config.Snapshots.Enabled)
		}
		if uc.SnapshotsKey != got.Config.Snapshots.Key {
			t.Errorf("Want SnapshotsKey %s and got %s", uc.SnapshotsKey, got.Config.Snapshots.Key)
		}
		if uc.ArtifactName != got.Data.ArtifactName {
			t.Errorf("Want ArtifactName %s and got %s", ArtifactName, got.Data.ArtifactName)
		}
		if uc.Major != got.Data.Version.Major {
			t.Errorf("Want major version %d and got %d", Major, got.Data.Version.Major)
		}
		if uc.Minor != got.Data.Version.Minor {
			t.Errorf("Want minor version %d and got %d", Minor, got.Data.Version.Minor)
		}
		if uc.Patch != got.Data.Version.Patch {
			t.Errorf("Want patch version %d and got %d", Patch, got.Data.Version.Patch)
		}
	}
}

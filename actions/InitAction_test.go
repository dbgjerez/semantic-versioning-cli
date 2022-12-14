package actions

import (
	"semver/domain"
	"testing"
)

func TestNewConfig(t *testing.T) {
	want := domain.Store{
		Config: domain.SemverConfig{
			Snapshots: domain.SemverSubType{
				Key:     INIT_SNAPSHOTS_KEY,
				Enabled: INIT_SNAPSHOTS_ENABLED,
			},
		},
		Data: domain.DataStore{
			ArtifactName: ArtifactName,
			Version: domain.VersionConfig{
				Major: Major,
				Minor: Minor,
				Patch: Patch,
			},
		},
	}
	useCases := []InitAction{
		{
			ArtifactName:    ArtifactName,
			Major:           Major,
			Minor:           Minor,
			Patch:           Patch,
			SnapshotsEnable: INIT_SNAPSHOTS_ENABLED,
			SnapshotsKey:    INIT_SNAPSHOTS_KEY,
		},
		{
			ArtifactName:    ArtifactName,
			Major:           Major,
			Minor:           Minor,
			Patch:           Patch,
			SnapshotsEnable: INIT_SNAPSHOTS_ENABLED,
			SnapshotsKey:    "",
		},
	}
	for _, uc := range useCases {
		got, err := uc.NewConfig()
		if err != nil {
			t.Errorf("Don't expected fails!")
		}
		if want.Config.Snapshots.Key != got.Config.Snapshots.Key ||
			want.Config.Snapshots.Key != INIT_SNAPSHOTS_KEY ||
			got.Config.Snapshots.Key != INIT_SNAPSHOTS_KEY {
			t.Errorf("Want ArtifactName %s and got %s", INIT_SNAPSHOTS_KEY, got.Config.Snapshots.Key)
		}
		if want.Data.ArtifactName != got.Data.ArtifactName ||
			want.Data.ArtifactName != ArtifactName ||
			got.Data.ArtifactName != ArtifactName {
			t.Errorf("Want ArtifactName %s and got %s", ArtifactName, got.Data.ArtifactName)
		}
		if want.Data.Version.Major != got.Data.Version.Major ||
			want.Data.Version.Major != Major ||
			got.Data.Version.Major != Major {
			t.Errorf("Want major version %d and got %d", Major, got.Data.Version.Major)
		}
		if want.Data.Version.Minor != got.Data.Version.Minor ||
			want.Data.Version.Minor != Minor ||
			got.Data.Version.Minor != Minor {
			t.Errorf("Want minor version %d and got %d", Minor, got.Data.Version.Minor)
		}
		if want.Data.Version.Patch != got.Data.Version.Patch ||
			want.Data.Version.Patch != Patch ||
			got.Data.Version.Patch != Patch {
			t.Errorf("Want patch version %d and got %d", Patch, got.Data.Version.Patch)
		}
	}
}

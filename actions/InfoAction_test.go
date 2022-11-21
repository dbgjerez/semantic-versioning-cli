package actions

import (
	"semver/domain"
	"testing"
)

const (
	ArtifactName string = "semver"
	Major        int    = 2
	Minor        int    = 1
	Patch        int    = 1
	SnapshotKey  string = "SNAPSHOT"
)

func NewConfigMockWithVersion(enableSnapshots bool, version domain.VersionConfig) *domain.Store {
	return &domain.Store{
		Data:   NewConfigMockVersion(version),
		Config: NewConfigConfigMock(enableSnapshots),
	}
}

func NewConfigMock(enableSnapshots bool) *domain.Store {
	version := NewVersionMock(Major, Minor, Patch)
	return NewConfigMockWithVersion(enableSnapshots, version)
}

func NewConfigConfigMock(enableSnapshots bool) domain.SemverConfig {
	return domain.SemverConfig{
		Versions: domain.SemverConfigVersions{
			Snapshot: domain.SemverConfigSnapshots{
				Enabled: enableSnapshots,
				Key:     SnapshotKey,
			},
		},
	}
}

func NewConfigMockVersion(v domain.VersionConfig) domain.DataStore {
	data := domain.DataStore{
		ArtifactName: ArtifactName,
		Version:      v,
	}
	return data
}

func NewInfoActionMock() InfoAction {
	c := NewConfigMock(true)
	return InfoAction{c: c}
}

func NewVersionMock(major int, minor int, patch int) domain.VersionConfig {
	return domain.VersionConfig{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func TestNewInfoAction(t *testing.T) {
	c := NewConfigMock(true)
	want := InfoAction{c: c}
	got := NewInfoAction(c)

	if want != got {
		t.Errorf("got different action")
	}
}

func TestCompleteInfo(t *testing.T) {
	action := NewInfoActionMock()
	want := "Artifact name: semver\n" +
		"Version: 2.1.1\n"
	got := action.CompleteInfo()

	if want != got {
		t.Errorf("got different action")
	}
}

func TestArtifactName(t *testing.T) {
	action := NewInfoActionMock()
	want := ArtifactName
	got := action.ArtifactName()

	if want != got {
		t.Errorf("Expected artifactName %s and got %s", want, got)
	}
}

func TestArtifactVersion(t *testing.T) {
	type VersionTest struct {
		v    domain.VersionConfig
		want string
	}
	versions := []VersionTest{
		{
			v:    NewVersionMock(1, 1, 0),
			want: "1.1",
		},
		{
			v:    NewVersionMock(1, 1, 1),
			want: "1.1.1",
		},
	}
	for _, v := range versions {
		c := NewConfigMockWithVersion(true, v.v)
		action := NewInfoAction(c)
		got := action.ArtifactVersion()
		if v.want != got {
			t.Errorf("Expected version %s and got %s", v.want, got)
		}
	}
}

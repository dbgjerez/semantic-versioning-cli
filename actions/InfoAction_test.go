package actions

import (
	"semver/domain"
	"testing"
)

var artifactName string = "semver"
var major int = 2
var minor int = 1
var patch int = 1

func NewConfigMock() domain.Config {
	return NewConfigMockVersion(domain.VersionConfig{
		Major: major,
		Minor: minor,
		Patch: patch,
	})
}

func NewConfigMockVersion(v domain.VersionConfig) domain.Config {
	return domain.Config{
		Data: domain.DataConfig{
			ArtifactName: artifactName,
			Version:      v,
		},
	}
}

func NewInfoActionMock() InfoAction {
	c := NewConfigMock()
	return InfoAction{c: &c}
}

func TestNewInfoAction(t *testing.T) {
	c := NewConfigMock()
	want := InfoAction{c: &c}
	got := NewInfoAction(&c)

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
	want := artifactName
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
			v: domain.VersionConfig{
				Major: 1,
				Minor: 1,
				Patch: 0,
			},
			want: "1.1",
		},
		{
			v: domain.VersionConfig{
				Major: 1,
				Minor: 1,
				Patch: 1,
			},
			want: "1.1.1",
		},
	}
	for _, v := range versions {
		c := NewConfigMockVersion(v.v)
		action := NewInfoAction(&c)
		got := action.ArtifactVersion()
		if v.want != got {
			t.Errorf("Expected version %s and got %s", v.want, got)
		}
	}
}

package actions

import (
	"fmt"
	"semver/domain"
	"testing"
)

var artifactName string = "semver"
var major int = 2
var minor int = 1
var patch int = 1

func NewConfigMock() domain.Config {
	return domain.Config{
		Data: domain.DataConfig{
			ArtifactName: artifactName,
			Version: domain.VersionConfig{
				Major: major,
				Minor: minor,
				Patch: patch,
			},
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

func (info *InfoAction) TestArtifactName() string {
	return info.c.Data.ArtifactName
}

func (info *InfoAction) TestArtifactVersion() string {
	version := fmt.Sprintf("%d.%d", info.c.Data.Version.Major, info.c.Data.Version.Minor)
	if info.c.Data.Version.Patch != 0 {
		version += fmt.Sprintf(".%d", info.c.Data.Version.Patch)
	}
	return version
}

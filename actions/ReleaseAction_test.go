package actions

import (
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

func TestCreateMajor(t *testing.T) {
	// Received -1 => Default case (sum 1)
	// Received == actual version ==> Error
	// Received > actual version ==> Force version
	// Received < actual version ==> Error
	type WantTest struct {
		err     bool
		version string
	}
	type VersionTest struct {
		v        domain.VersionConfig
		snapshot bool
		param    int
		want     WantTest
	}
	versions := []VersionTest{
		{
			v:        NewVersionMock(1, 1, 0),
			snapshot: false,
			param:    -1,
			want:     WantTest{err: false, version: "2.0"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: false,
			param:    1,
			want:     WantTest{err: true},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: false,
			param:    2,
			want:     WantTest{err: false, version: "2.0"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: false,
			param:    4,
			want:     WantTest{err: false, version: "4.0"},
		},
		{
			v:        NewVersionMock(4, 1, 1),
			snapshot: false,
			param:    2,
			want:     WantTest{err: true},
		},
		{
			v:        NewVersionMock(1, 1, 0),
			snapshot: true,
			param:    -1,
			want:     WantTest{err: false, version: "2.0-SNAPSHOT"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: true,
			param:    1,
			want:     WantTest{err: true},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: true,
			param:    2,
			want:     WantTest{err: false, version: "2.0-SNAPSHOT"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: true,
			param:    4,
			want:     WantTest{err: false, version: "4.0-SNAPSHOT"},
		},
		{
			v:        NewVersionMock(4, 1, 1),
			snapshot: true,
			param:    2,
			want:     WantTest{err: true},
		},
	}
	for _, v := range versions {
		c := NewConfigMockWithVersion(v.snapshot, v.v)
		action := NewReleaseAction(c)
		infoAction := NewInfoAction(c)
		config, e := action.CreateMajor(v.param)
		resInfoAction := NewInfoAction(&config)
		if e == nil && v.want.err {
			t.Errorf("Expected error with version %s and param %d, but got a version %s",
				infoAction.ArtifactVersion(),
				v.param,
				resInfoAction.ArtifactVersion())
		} else if e != nil && !v.want.err {
			t.Errorf("Unexpected error %v with version %s and param %d",
				e,
				infoAction.ArtifactVersion(),
				v.param)
		} else if !v.want.err && resInfoAction.ArtifactVersion() != v.want.version {
			t.Errorf("Expected version %s with param %d and got version %s",
				v.want.version,
				v.param,
				resInfoAction.ArtifactVersion())
		}
	}
}

func TestCreateFeature(t *testing.T) {
	// Received -1 => Default case (sum 1)
	// Received == actual version ==> Error
	// Received > actual version ==> Force version
	// Received < actual version ==> Error
	type WantTest struct {
		err     bool
		version string
	}
	type VersionTest struct {
		v        domain.VersionConfig
		snapshot bool
		param    int
		want     WantTest
	}
	versions := []VersionTest{
		{
			v:        NewVersionMock(1, 1, 0),
			snapshot: false,
			param:    -1,
			want:     WantTest{err: false, version: "1.2"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: false,
			param:    -1,
			want:     WantTest{err: false, version: "1.2"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: false,
			param:    1,
			want:     WantTest{err: true},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: false,
			param:    2,
			want:     WantTest{err: false, version: "1.2"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: false,
			param:    4,
			want:     WantTest{err: false, version: "1.4"},
		},
		{
			v:        NewVersionMock(1, 4, 1),
			snapshot: false,
			param:    2,
			want:     WantTest{err: true},
		},
		{
			v:        NewVersionMock(1, 1, 0),
			snapshot: true,
			param:    -1,
			want:     WantTest{err: false, version: "1.2-SNAPSHOT"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: true,
			param:    -1,
			want:     WantTest{err: false, version: "1.2-SNAPSHOT"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: true,
			param:    1,
			want:     WantTest{err: true},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: true,
			param:    2,
			want:     WantTest{err: false, version: "1.2-SNAPSHOT"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: true,
			param:    4,
			want:     WantTest{err: false, version: "1.4-SNAPSHOT"},
		},
		{
			v:        NewVersionMock(1, 4, 1),
			snapshot: true,
			param:    2,
			want:     WantTest{err: true},
		},
	}
	for _, v := range versions {
		c := NewConfigMockWithVersion(v.snapshot, v.v)
		action := NewReleaseAction(c)
		infoAction := NewInfoAction(c)
		initV := infoAction.ArtifactVersion()
		config, e := action.CreateFeature(v.param)
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
		v        domain.VersionConfig
		snapshot bool
		param    int
		want     WantTest
	}
	versions := []VersionTest{
		{
			v:        NewVersionMock(1, 1, 0),
			snapshot: false,
			param:    -1,
			want:     WantTest{err: false, version: "1.1.1"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: false,
			param:    -1,
			want:     WantTest{err: false, version: "1.1.2"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: false,
			param:    1,
			want:     WantTest{err: true},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: false,
			param:    2,
			want:     WantTest{err: false, version: "1.1.2"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: false,
			param:    4,
			want:     WantTest{err: false, version: "1.1.4"},
		},
		{
			v:        NewVersionMock(1, 1, 4),
			snapshot: true,
			param:    2,
			want:     WantTest{err: true},
		},
		{
			v:        NewVersionMock(1, 1, 0),
			snapshot: true,
			param:    -1,
			want:     WantTest{err: false, version: "1.1.1-SNAPSHOT"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: true,
			param:    -1,
			want:     WantTest{err: false, version: "1.1.2-SNAPSHOT"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: true,
			param:    1,
			want:     WantTest{err: true},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: true,
			param:    2,
			want:     WantTest{err: false, version: "1.1.2-SNAPSHOT"},
		},
		{
			v:        NewVersionMock(1, 1, 1),
			snapshot: true,
			param:    4,
			want:     WantTest{err: false, version: "1.1.4-SNAPSHOT"},
		},
		{
			v:        NewVersionMock(1, 1, 4),
			snapshot: true,
			param:    2,
			want:     WantTest{err: true},
		},
	}
	for _, v := range versions {
		c := NewConfigMockWithVersion(v.snapshot, v.v)
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

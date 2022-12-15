package domain

import (
	"encoding/json"
	"reflect"
	"testing"
)

const (
	COMPLETE_JSON = `{
		"data": {
		  "artifactName": "semantic-versioning-cli",
		  "version": {
			"major": 0,
			"minor": 0,
			"patch": 0,
			"rc": 0,
			"snapshot": true
		  }
		},
		"config": {
		  "snapshots": {
			"enabled": true,
			"key": "SNAPSHOT"
		  },
		  "release-candidates": {
			"enabled": true,
			"key": "RC"
		  },
		  "gitflow": {
			"enabled": true,
			"branches": {
			  "snapshots": [
				"develop",
				"feature/*"
			  ],
			  "releases-candidates": [
				"release/*"
			  ]
			}
		  }
		}
	  }`
)

func TestCompleteModel(t *testing.T) {
	var config Store
	err := json.Unmarshal([]byte(COMPLETE_JSON), &config)
	if err != nil {
		t.Errorf("Invalid JSON")
	}

	if !config.Data.Version.Snapshot ||
		!config.Config.ReleaseCandidates.Enabled ||
		!config.Config.Snapshots.Enabled ||
		!config.Config.GitFlow.Enabled {
		t.Errorf("Expected a correct JSON config load")
	}
}

type Test struct {
	name           string
	config         Store
	expectedResult bool
}

func TestIsSnapshot(t *testing.T) {
	tests := []Test{
		{
			name: "snapshot-enabled",
			config: Store{
				Config: SemverConfig{
					Snapshots: SemverSubType{
						Enabled: true,
					},
				},
			},
			expectedResult: true,
		},
		{
			name: "no-snapshot-enabled",
			config: Store{
				Config: SemverConfig{
					Snapshots: SemverSubType{
						Enabled: false,
					},
				},
			},
			expectedResult: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if !reflect.DeepEqual(test.config.IsSnapshotEnabled(), test.expectedResult) {
				t.Errorf("Expected %t and got %t",
					test.config.IsSnapshotEnabled(),
					test.expectedResult)
			}
		})
	}
}

package domain

import (
	"encoding/json"
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

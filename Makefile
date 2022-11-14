#!make

.DEFAULT_GOAL := help
VAR_FILE := .semver
MAJOR := $(shell grep MAJOR $(VAR_FILE) | cut -d':' -f2)
MINOR := $(shell grep MINOR $(VAR_FILE) | cut -d':' -f2)
PATCH := $(shell grep PATCH $(VAR_FILE) | cut -d':' -f2)
ARTIFACT_NAME := $(shell grep ARTIFACT_NAME $(VAR_FILE) | cut -d':' -f2)

.PHONY: help 
help: ## Show options and short description
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: version-major
version-major: ## Upgrade the major version
	@sed -i 's/MAJOR: '$$(( $(MAJOR)+0 ))'/MAJOR: '$$(( $(MAJOR)+1 ))'/g' $(VAR_FILE)
	@sed -i 's/MINOR: '$$(( $(MINOR)+0 ))'/MINOR: 0/g' $(VAR_FILE)
	@sed -i 's/PATCH: '$$(( $(PATCH)+0 ))'/PATCH: 0/g' $(VAR_FILE)

.PHONY: version-minor
version-minor: ## Upgrade the minor version
	@sed -i 's/MINOR: '$$(( $(MINOR)+0 ))'/MINOR: '$$(( $(MINOR)+1 ))'/g' $(VAR_FILE)
	@sed -i 's/PATCH: '$$(( $(PATCH)+0 ))'/PATCH: 0/g' $(VAR_FILE)

.PHONY: version-patch
version-patch: ## Upgrade the patch number
	@sed -i 's/PATCH: '$$(( $(PATCH)+0 ))'/PATCH: '$$(( $(PATCH)+1 ))'/g' $(VAR_FILE)

.PHONY: info
info: ## Show the artifact info
	@echo 'Name': $(ARTIFACT_NAME)
	@echo 'Version: '$$(( $(MAJOR)+0 )).$$(( $(MINOR)+0 )).$$(( $(PATCH)+0 ))

.PHONY: init
init: ## Init the process
	@echo ARTIFACT_NAME: $(artifactName) > $(VAR_FILE)
	@echo MAJOR: $(major) >> $(VAR_FILE)
	@echo MINOR: $(minor) >> $(VAR_FILE)
	@echo PATCH: $(patch) >> $(VAR_FILE)

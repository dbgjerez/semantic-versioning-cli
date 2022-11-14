# semantic-versioning-cli

## Help
```bash
$ make
help                           Show options and short description
info                           Show the artifact info
init                           Init the process
version-major                  Upgrade the major version
version-minor                  Upgrade the minor version
version-patch                  Upgrade the patch number
```

## Init a project
The following parameters are mandatory to initialize a new project.

|Param|Description|
|--|--|
|artifactName|Artifact name|
|major|Default major version|
|minor|Default minor version|
|patch|Default patch version|

This example, initialize a project with the version ```0.0.0```
```zsh
make init \
    artifactName=semantic-versioning-cli \
    major=0 \
    minor=0 \
    patch=0 
```

## Info
To show the version of the project and the name:

```zsh
$ make info
Name: semantic-versioning-cli
Version: 0.0.0
```

## Versioning
### New patch
The command ```version-patch``` increases the patch number:

```zsh
$ make version-patch
$ make info | grep Version | cut -d ' ' -f2
0.0.1
```
### New minor version
The command ```version-minor``` increases the minor version. It should be called when we have a new functionality finished. 

At the same time, the patch version will change to zero. 

```zsh
$ make version-minor
$ make info | grep Version | cut -d ' ' -f2
0.1.0
```

### Upgrade major version
When we have a big change in our architecture, we should increase the major number version:

```zsh
$ make version-major
$ make info | grep Version | cut -d ' ' -f2
1.0.0
```

## Roadmap
 * Change to Golang cli 
 * Change data sctructure to .json file
# semantic-versioning-cli
```semver``` is an Open Source project used to manage application versions in a decoupled way.

## Help
```bash
$ semver help  
NAME:
   semver - A new cli application

USAGE:
   semver [global options] command [command options] [arguments...]

COMMANDS:
   info, i     Show the artifact info
   release, r  Create a new release
   feature, f  Create a new feature
   patch, p    Create a new patch
   init        Init the versioning configuration file
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --file value, -f value  Config file (default: ".semver.yaml")
   --help, -h              show help (default: false)
```

## Usage
### Init a project
The following parameters are mandatory to initialize a new project.

|Param|Alias|Default value|Description|
|--|--|--|--|
|--name value|-n value||Artifact name|
|--major value|-ma value|0|Default major version|
|--minor value|-mi value|0|Default minor version|
|--patch value|-p value|0|Default patch version|

This example, initialize a project with the version ```0.0.0```
```bash
semver init \
    --name semantic-versioning-cli 
```

## Info
To show the version of the project and the name:

```zsh
$ semver info 
Artifact name: semantic-versioning-cli
Version: 1.0.1
```

## Versioning
### Major
This number should increment when you make an incompatible API change.

When you increment the major version, automatically the minor and path version changes to zero. 

```zsh
$ semver m
$ semver info
Artifact name: semantic-versioning-cli
Version: 2.0
```

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



## Roadmap
 * Change to Golang cli 
 * Change data sctructure to .json file
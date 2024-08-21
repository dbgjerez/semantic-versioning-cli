# semantic-versioning-cli
```semver``` is an Open Source project used to manage application versions in a decoupled way.

## Installation
```bash
wget https://github.com/dbgjerez/semantic-versioning-cli/releases/download/1.2/semver -O semver
chmod +x semver
sudo mv semver /usr/local/bin/
semver --version
```

## Help
```bash
$ semver help  
NAME:
   semver - A new cli application

USAGE:
   semver [global options] command [command options] [arguments...]

COMMANDS:
   info, i     Show the artifact info
   major, m    Create a new major version
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
|--snapshot|-s|false|Enable snapshots versions|

This example, initialize a project with the version ```0.0.0```
```bash
semver init \
    --name semantic-versioning-cli 
```

### Info
To show the version of the project and the name:

```zsh
$ semver info 
Artifact name: semantic-versioning-cli
Version: 1.0.1
```

In addition, you can use flags to retrieve only the desired information. For more information: ```semver info --help```

## Versioning
### Major
This number should increment when you make an incompatible API change.

When you increment the major version, automatically the minor and path version changes to zero. 

```bash
$ semver info v 
1.0.1

$ semver m
2.0
```

Besides, you can force the version using the ```-f``` option. 

For more information: ```semver major --help```

### Feature
The minor version upgrades after a new feature or functionality. 

At the same time, the patch version will change to zero. 

```bash
$ semver info
Artifact name: semantic-versioning-cli
Version: 1.0.1

$ semver feature
1.1                                                       
```

Like the major version, the minor version can be forced using the ```-f``` flag. 

For more information: ```semver feature --help```

### New patch
This number is used for bug-fix control and should be upgraded for each new release with bug fixes.

```zsh
$ semver info 
Artifact name: semantic-versioning-cli
Version: 1.0.1

$ semver patch
1.0.2
```

You can force this part, in the same way, that previous numbers. 

To amplify the information: ```semver patch --help```

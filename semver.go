package main

import (
	"fmt"
	"log"
	"os"
	"semver/actions"
	"semver/domain"

	"github.com/urfave/cli/v2"
)

var (
	version    string 
	appName    = "SemVer"
	repository = "https://github.com/dbgjerez/semantic-versioning-cli"
)

func main() {
	var file string

	cli.VersionPrinter = func(cCtx *cli.Context) {
		fmt.Printf("%s - %s\nVersion: %s\nRepository: %s\n", appName, cCtx.App.Usage, cCtx.App.Version, repository)
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the application name, repository, and version",
	}

	app := &cli.App{
		Name:    appName,
		Usage:   "A tool for managing semantic versioning",
		Version: version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Usage:       "Config file",
				Value:       ".semver.yaml",
				Destination: &file,
			},
		},
		Commands: []*cli.Command{
			newInfoCommand(&file),
			newMajorCommand(&file),
			newFeatureCommand(&file),
			newPatchCommand(&file),
			newSnapshotCommand(&file),
			newInitCommand(&file),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newInfoCommand(file *string) *cli.Command {
	return &cli.Command{
		Name:    "info",
		Aliases: []string{"i"},
		Usage:   "Show the artifact info",
		Action: func(*cli.Context) error {
			store := domain.NewConfigStore(*file)
			c, err := store.ReadConfig()
			if err != nil {
				return err
			}
			action := actions.NewInfoAction(&c)
			fmt.Printf(action.CompleteInfo())
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Artifact version",
				Action: func(*cli.Context) error {
					store := domain.NewConfigStore(*file)
					c, err := store.ReadConfig()
					if err != nil {
						return err
					}
					action := actions.NewInfoAction(&c)
					fmt.Printf(action.ArtifactVersion())
					return nil
				},
			},
		},
	}
}

func newMajorCommand(file *string) *cli.Command {
	return &cli.Command{
		Name:    "major",
		Aliases: []string{"m"},
		Usage:   "Create a new major version",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Force the version",
				Value:   -1,
			},
		},
		Action: func(ctx *cli.Context) error {
			r := ctx.Int("force")
			store := domain.NewConfigStore(*file)
			c, err := store.ReadConfig()
			if err != nil {
				return err
			}
			action := actions.NewReleaseAction(&c)
			config, err2 := action.CreateMajor(r)
			if err2 != nil {
				return err2
			}
			infoAction := actions.NewInfoAction(&config)
			fmt.Printf(infoAction.ArtifactVersion())
			return store.SaveConfig(config)
		},
	}
}

func newFeatureCommand(file *string) *cli.Command {
	return &cli.Command{
		Name:    "feature",
		Aliases: []string{"f"},
		Usage:   "Create a new feature",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Force the version",
				Value:   -1,
			},
		},
		Action: func(ctx *cli.Context) error {
			r := ctx.Int("force")
			store := domain.NewConfigStore(*file)
			c, err := store.ReadConfig()
			if err != nil {
				return err
			}
			action := actions.NewReleaseAction(&c)
			config, err2 := action.CreateFeature(r)
			if err2 != nil {
				return err2
			}
			infoAction := actions.NewInfoAction(&config)
			fmt.Printf(infoAction.ArtifactVersion())
			return store.SaveConfig(config)
		},
	}
}

func newPatchCommand(file *string) *cli.Command {
	return &cli.Command{
		Name:    "patch",
		Aliases: []string{"p"},
		Usage:   "Create a new patch",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Force the version",
				Value:   -1,
			},
		},
		Action: func(ctx *cli.Context) error {
			r := ctx.Int("force")
			store := domain.NewConfigStore(*file)
			c, err := store.ReadConfig()
			if err != nil {
				return err
			}
			action := actions.NewReleaseAction(&c)
			config, err2 := action.CreatePatch(r)
			if err2 != nil {
				return err2
			}
			infoAction := actions.NewInfoAction(&config)
			fmt.Printf(infoAction.ArtifactVersion())
			return store.SaveConfig(config)
		},
	}
}

func newSnapshotCommand(file *string) *cli.Command {
	return &cli.Command{
		Name:    "snapshot",
		Aliases: []string{"sn"},
		Usage:   "Modify the snapshot flag",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
				Usage:   "Force the snapshot value",
			},
		},
		Action: func(ctx *cli.Context) error {
			force := ctx.Bool("force")
			store := domain.NewConfigStore(*file)
			config, err := store.ReadConfig()
			if err != nil {
				return err
			}

			action := actions.SnapshotAction{
				C:           &config,
				Force:       true,
				ForcedValue: force,
			}

			if err := action.ChangeStatus(); err != nil {
				return err
			}

			infoAction := actions.NewInfoAction(&config)
			fmt.Printf(infoAction.ArtifactVersion())
			return store.SaveConfig(config)
		},
	}
}

func newInitCommand(file *string) *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: "Init the versioning configuration file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "Artifact name",
				Required: true,
			},
			&cli.IntFlag{
				Name:    "major",
				Aliases: []string{"ma"},
				Usage:   "Init major number",
				Value:   actions.INIT_MAJOR_VERSION,
			},
			&cli.IntFlag{
				Name:    "minor",
				Aliases: []string{"mi"},
				Usage:   "Init minor number",
				Value:   actions.INIT_MINOR_VERSION,
			},
			&cli.IntFlag{
				Name:    "patch",
				Aliases: []string{"p"},
				Usage:   "Init patch number",
				Value:   actions.INIT_PATCH_VERSION,
			},
			&cli.BoolFlag{
				Name:    "snapshot",
				Aliases: []string{"s"},
				Usage:   "Enable Snapshots",
			},
		},
		Action: func(ctx *cli.Context) error {
			store := domain.NewConfigStore(*file)
			if store.Exists() {
				return fmt.Errorf("Project already initialized!")
			}
			action := actions.InitAction{
				ArtifactName:    ctx.String("name"),
				Major:           ctx.Int("major"),
				Minor:           ctx.Int("minor"),
				Patch:           ctx.Int("patch"),
				SnapshotsEnable: ctx.Bool("snapshot"),
			}
			config, err := action.NewConfig()
			if err != nil {
				return err
			}
			return store.SaveConfig(config)
		},
	}
}

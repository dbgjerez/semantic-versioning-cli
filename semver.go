package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"semver/actions"
	"semver/domain"

	"github.com/urfave/cli/v2"
)

func main() {
	var file string
	var store domain.ConfigStore

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Usage:       "Config file",
				Value:       ".semver.yaml",
				Required:    false,
				Destination: &file,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "info",
				Aliases: []string{"i"},
				Usage:   "Show the artifact info",
				Action: func(*cli.Context) error {
					store := domain.NewConfigStore(file)
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
							store = domain.NewConfigStore(file)
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
			},
			{
				Name:    "major",
				Aliases: []string{"m"},
				Usage:   "Create a new major version",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "force",
						Aliases: []string{"f"},
						Usage:   "force the version",
						Value:   -1,
					},
				},
				Action: func(ctx *cli.Context) error {
					r := ctx.Int("force")
					store := domain.NewConfigStore(file)
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
			},
			{
				Name:    "feature",
				Aliases: []string{"f"},
				Usage:   "Create a new feature",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "force",
						Aliases: []string{"f"},
						Usage:   "force the version",
						Value:   -1,
					},
				},
				Action: func(ctx *cli.Context) error {
					r := ctx.Int("force")
					store := domain.NewConfigStore(file)
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
			},
			{
				Name:    "patch",
				Aliases: []string{"p"},
				Usage:   "Create a new patch",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "force",
						Aliases: []string{"f"},
						Usage:   "force the version",
						Value:   -1,
					},
				},
				Action: func(ctx *cli.Context) error {
					r := ctx.Int("force")
					store := domain.NewConfigStore(file)
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
			},
			{
				Name:    "snapshot",
				Aliases: []string{"sn"},
				Usage:   "Modify the snapshot flag",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "force",
						Aliases: []string{"f"},
						Usage:   "force the snapshot value",
					},
				},
				Action: func(ctx *cli.Context) error {
					force := false
					if ctx.Count("force") > 0 { // if > 0, user set value
						force = true
					}
					forcedValue := ctx.Bool("force")
					store := domain.NewConfigStore(file)
					config, err := store.ReadConfig()
					if err != nil {
						return err
					}

					action := actions.SnapshotAction{
						C:           &config,
						Force:       force,
						ForcedValue: forcedValue,
					}

					err = action.ChangeStatus()
					if err != nil {
						return err
					}

					infoAction := actions.NewInfoAction(&config)
					fmt.Printf(infoAction.ArtifactVersion())
					return store.SaveConfig(config)
				},
			},
			{
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
						Value:   actions.INIT_SNAPSHOTS_ENABLED,
					},
				},
				Action: func(ctx *cli.Context) error {
					store = domain.NewConfigStore(file)
					if store.Exists() {
						return errors.New("Project initialized yet!")
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
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

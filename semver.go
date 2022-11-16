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
					fmt.Printf("Name: %s", c.Data.ArtifactName)
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
				Name:  "release",
				Usage: "Create a new release",
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
					action := actions.NewReleaseAction(c)
					config, err2 := action.CreateRelease(r)
					if err2 != nil {
						return err2
					}
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
						Value:   0,
					},
					&cli.IntFlag{
						Name:    "minor",
						Aliases: []string{"mi"},
						Usage:   "Init minor number",
						Value:   0,
					},
					&cli.IntFlag{
						Name:    "patch",
						Aliases: []string{"p"},
						Usage:   "Init patch number",
						Value:   0,
					},
				},
				Action: func(ctx *cli.Context) error {
					store = domain.NewConfigStore(file)
					if store.Exists() {
						return errors.New("Project initialized yet!")
					}
					action := actions.InitAction{
						ArtifactName: ctx.String("name"),
						Major:        ctx.Int("major"),
						Minor:        ctx.Int("minor"),
						Patch:        ctx.Int("patch"),
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

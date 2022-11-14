package main

import (
	"errors"
	"fmt"
	"log"
	"os"
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
					store = domain.NewConfigStore(file)
					c, err := store.ReadConfig()
					if err != nil {
						return err
					}
					fmt.Printf("Name: %s", c.Data.ArtifactName)
					return nil
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
						Value:   1,
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
					config := domain.Config{
						Data: domain.DataConfig{
							ArtifactName: ctx.String("name"),
							Version: domain.VersionConfig{
								Major: ctx.Int("major"),
								Minor: ctx.Int("minor"),
								Patch: ctx.Int("patch"),
							},
						},
					}
					fmt.Printf("Name: %s\nMajor: %d\nMinor: %d\nPatch: %d\n",
						config.Data.ArtifactName,
						config.Data.Version.Major,
						config.Data.Version.Minor,
						config.Data.Version.Patch)
					return store.SaveConfig(config)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

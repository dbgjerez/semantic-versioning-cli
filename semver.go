package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "info",
				Aliases: []string{"i"},
				Usage:   "Show the artifact info",
				Action: func(*cli.Context) error {
					fmt.Println("Hello world")
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
					fmt.Printf("Name: %s\nMajor: %d\nMinor: %d\nPatch: %d\n",
						ctx.String("name"),
						ctx.Int("major"),
						ctx.Int("minor"),
						ctx.Int("patch"))
					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

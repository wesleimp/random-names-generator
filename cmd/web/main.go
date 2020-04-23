package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/wesleimp/random-names-generator/cmd/web/actions"
)

func main() {
	app := &cli.App{
		Name:  "Sprint name generator",
		Usage: "Generate funny randon name for your sprints or teams",
		Authors: []*cli.Author{
			{
				Email: "wesleimsr@gmail.com",
				Name:  "Weslei Juan Moser Pereira",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "addr",
				Usage:   "Application port address",
				Value:   ":8080",
				EnvVars: []string{"ADDR"},
			},
		},
		Action: actions.Run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Error starting application")
	}
}

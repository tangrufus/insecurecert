package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"insecurecert/command"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			(&command.Trust{}).Command(),
			(&command.Sponsor{}).Command(),
			(&command.Bug{}).Command(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
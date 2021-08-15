package main

import (
	"os"

	logg "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const usage = `docker is a simple container runtime implementation.`

func main() {
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		// Log as JSON instead of the default ASCII formatter.
		logg.SetFormatter(&logg.JSONFormatter{})

		logg.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logg.Fatal(err)
	}
}

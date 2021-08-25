package main

import (
	"fmt"

	"github.com/docker/cgroups/subsystems"
	"github.com/docker/container"
	logg "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name: "run",
	Usage: `Create a container with namespace and cgroups limit
			docker run -ti [command]`,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "ti",
			Usage: "enable tty",
		},
		cli.BoolFlag{
			Name:  "d",
			Usage: "detach container",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "container name",
		},
	},
	Action: func(context *cli.Context) error {
		if len(context.Args()) < 1 {
			return fmt.Errorf("missing container command")
		}
		var cmdArray []string
		for _, arg := range context.Args() {
			cmdArray = append(cmdArray, arg)
		}
		createTty := context.Bool("ti")
		detach := context.Bool("d")
		if createTty && detach {
			return fmt.Errorf("can not both provided")
		}

		resConf := &subsystems.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuSet:      context.String("cpuset"),
			CpuShare:    context.String("cpushare"),
		}
		logg.Infof("createTty %v", createTty)
		//Run(tty, cmdArray, resConf)

		contianerName := context.String("name")
		Run(createTty, cmdArray, resConf, contianerName)
		return nil
	},
}
var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process",
	Action: func(context *cli.Context) error {
		logg.Infof("init come on")
		err := container.RunContainerInitProcess()
		return err
	},
}

var listCommand = cli.Command{
	Name:  "ps",
	Usage: "list all the containers",
	Action: func(context *cli.Context) error {
		container.ListContainers()
		return nil
	},
}

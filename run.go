package main

import (
	"os"

	"github.com/docker/container"
	logg "github.com/sirupsen/logrus"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		logg.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}

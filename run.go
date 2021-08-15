package main

import (
	"os"
	"strings"

	"github.com/docker/cgroups"
	"github.com/docker/cgroups/subsystems"
	"github.com/docker/container"
	logg "github.com/sirupsen/logrus"
)

func Run(tty bool, comArray []string, res *subsystems.ResourceConfig) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		logg.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		logg.Error(err)
	}
	// use mydocker-cgroup as cgroup name
	cgroupManager := cgroups.NewCgroupManager("mydocker-cgroup")
	defer cgroupManager.Destroy()
	cgroupManager.Set(res)
	cgroupManager.Apply(parent.Process.Pid)

	sendInitCommand(comArray, writePipe)
	parent.Wait()
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	logg.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}

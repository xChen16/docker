package main

import (
	"os"
	"strings"

	"github.com/docker/cgroups"
	"github.com/docker/cgroups/subsystems"
	"github.com/docker/container"
	logg "github.com/sirupsen/logrus"
)

func Run(tty bool, comArray []string, res *subsystems.ResourceConfig,contianerName) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		logg.Errorf("New parent process error")
		return
	}
	if err := parent.Start(); err != nil {
		logg.Error(err)
	}
	// use mydocker-cgroup as cgroup name

	cgroupManager := cgroups.NewCgroupManager("docker-cgroup")
	defer cgroupManager.Destroy()
	cgroupManager.Set(res)
	cgroupManager.Apply(parent.Process.Pid)

	sendInitCommand(comArray, writePipe)
	if tty {
		parent.Wait()
	}
	/*
		mntURL := "/root/mnt/"
		rootURL := "/root/"
		container.DeleteWorkSpace(rootURL, mntURL)
		os.Exit(0)
	*/
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	logg.Infof("command all is %s", command)
	writePipe.WriteString(command)
	writePipe.Close()
}
func randStringBytes(n int) string {
	letterBytes := "1234567890"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func recordContainerInfo(){
	
}
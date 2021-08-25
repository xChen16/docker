package container

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"

	logg "github.com/sirupsen/logrus"
)

func ListContainers() {
	dirURL := fmt.Sprintf(DefaultInfoLocation, "")
	dirURL = dirURL[:len(dirURL)-1]
	files, err := ioutil.ReadDir(dirURL)
	if err != nil {
		logg.Errorf("Read dir %s error %v", dirURL, err)
		return
	}

	var containers []*ContainerInfo
	for _, file := range files {
		tmpContainer, err := getContainerInfo(file)
		if err != nil {
			logg.Errorf("Get container info error %v", err)
			continue
		}
		containers = append(containers, tmpContainer)
	}

	w := tabwriter.NewWriter(os.Stdout, 12, 1, 3, ' ', 0)
	fmt.Fprint(w, "ID\tNAME\tPID\tSTATUS\tCOMMAND\tCREATED\n")
	for _, item := range containers {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
			item.Id,
			item.Name,
			item.Pid,
			item.Status,
			item.Command,
			item.CreatedTime)
	}
	if err := w.Flush(); err != nil {
		logg.Errorf("Flush error %v", err)
		return
	}
}

func getContainerInfo(file os.FileInfo) (*ContainerInfo, error) {
	containerName := file.Name()
	configFileDir := fmt.Sprintf(DefaultInfoLocation, containerName)
	configFileDir = configFileDir + ConfigName
	content, err := ioutil.ReadFile(configFileDir)
	if err != nil {
		logg.Errorf("Read file %s error %v", configFileDir, err)
		return nil, err
	}
	var containerInfo ContainerInfo
	if err := json.Unmarshal(content, &containerInfo); err != nil {
		logg.Errorf("Json unmarshal error %v", err)
		return nil, err
	}

	return &containerInfo, nil
}

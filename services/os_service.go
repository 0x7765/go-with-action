package services

import (
	"fmt"

	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
)

// OsInfo print host info
func OsInfo() {
	info, err := host.Info()
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", info)
}

// IsDocker check is docker
func IsDocker() bool {
	list, err := docker.GetDockerIDList()
	return err == nil && len(list) > 0
}

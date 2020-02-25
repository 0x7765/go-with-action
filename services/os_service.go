/**
@author: wei.xuan
@time: 2020/2/25 22:17
@desc:
*/
package services

import (
	"fmt"

	"github.com/shirou/gopsutil/docker"
	"github.com/shirou/gopsutil/host"
)

func OsInfo() {
	info, err := host.Info()
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", info)
}

func IsDocker() bool {
	list, err := docker.GetDockerIDList()
	return err == nil && len(list) > 0
}

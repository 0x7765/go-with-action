/**
@author: wei.xuan
@time: 2020/2/25 22:22
@desc:
*/
package services

import (
	"testing"
)

func TestOsInfo(t *testing.T) {
	OsInfo()
}

func TestIsDocker(t *testing.T) {
	t.Logf("is docker :%v\n", IsDocker())
}

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

package main

import (
	"github.com/shenzh1990/TopList/onstart"
	"github.com/shenzh1990/TopList/pkg/settings"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(settings.CpuMaxProcess)
	onstart.Run()
}

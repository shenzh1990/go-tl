package onstart

import (
	"fmt"
	"github.com/gotoeasy/glang/cmn"
	"github.com/shenzh1990/TopList/pkg/settings"
	"os"
)

func init() {
	cmn.Info("Daemon init")
	httpPort := cmn.IntToString(settings.HTTPPort)
	// 端口冲突时退出
	if cmn.IsPortOpening(httpPort) {
		fmt.Printf("port %s conflict, startup failed.\n", httpPort)
		os.Exit(0)
	}
}

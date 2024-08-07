package onstart

import (
	"github.com/shenzh1990/TopList/redisutil"
	"github.com/shenzh1990/TopList/router"
)

func Run() {
	redisutil.Start()
	router.Start()
}

package cron

import (
	"fmt"
	"github.com/gotoeasy/glang/cmn"
	"github.com/jasonlvhit/gocron"
	"github.com/shenzh1990/TopList/redisutil"
	"github.com/shenzh1990/TopList/service"
	"github.com/shenzh1990/TopList/util"
	"reflect"
	"strconv"
	"sync"
	"time"
)

var group sync.WaitGroup

func Start() {
	cmn.Info("cron Init")
	go exec()
	select {}
}
func exec() {
	cron()
	gocron.Every(1).Day().At("7:00").Do(cron)
	<-gocron.Start()
}
func cron() {
	allData := []string{
		"V2EX",
		"ZhiHu",
		"TieBa",
		"DouBan",
		"TianYa",
		"WeiBo",
		"HuPu",
		"GitHub",
		"BaiDu",
		"36Kr",
		"QDaily",
		"GuoKr",
		"HuXiu",
		"ZHDaily",
		"Segmentfault",
		"WYNews",
		"WaterAndWood",
		"HacPai",
		"KD",
		"NGA",
		"WeiXin",
		"Mop",
		"Chiphell",
		"ChouTi",
		"ITHome",
		"JianDan",
	}
	cmn.Info("开始抓取" + strconv.Itoa(len(allData)) + "种数据类型")
	//group.Add(len(allData))
	for _, value := range allData {
		cmn.Info("开始抓取" + value)
		ExecGetData(value)
	}
	//group.Wait()
	cmn.Info("完成抓取")
}
func ExecGetData(dataTypeValue string) {
	reflectValue := reflect.ValueOf(&service.HotSearchService{})
	methodName := "Get" + dataTypeValue
	if reflectValue.MethodByName(methodName).IsValid() == false {
		cmn.Error(methodName + "方法不存在")
		return
	}
	dataType := reflectValue.MethodByName(methodName)
	data := dataType.Call(nil)
	originData := data[0].Interface().([]map[string]interface{})
	start := time.Now()
	//Common.MySql{}.GetConn().Where(map[string]string{"dataType": spider.DataType}).Update("hotData2", map[string]string{"str": SaveDataToJson(originData)})
	redisutil.RDS.Set(redisutil.DEFAULT_REDIS_PRE_KEY+dataTypeValue, util.JsonResponse(1, "成功", originData))
	//group.Done()
	seconds := time.Since(start).Seconds()
	cmn.Info(fmt.Sprintf("耗费 %.2fs 秒完成抓取%s", seconds, dataTypeValue))
}

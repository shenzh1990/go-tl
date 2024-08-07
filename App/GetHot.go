package main

import (
	"encoding/json"
	"fmt"
	"github.com/shenzh1990/TopList/cron"
	"github.com/shenzh1990/TopList/redisutil"
	"log"
	"reflect"
	"sync"
	"time"
)

type HotData struct {
	Code    int
	Message string
	Data    interface{}
}

type Spider struct {
	DataType string
}

func SaveDataToJson(data interface{}) string {
	Message := HotData{}
	Message.Code = 0
	Message.Message = "获取成功"
	Message.Data = data
	jsonStr, err := json.Marshal(Message)
	if err != nil {
		log.Fatal("序列化json错误")
	}
	return string(jsonStr)

}

/*
*
执行每个分类数据
*/
func ExecGetData(spider Spider) {
	reflectValue := reflect.ValueOf(spider)
	dataType := reflectValue.MethodByName("Get" + spider.DataType)
	data := dataType.Call(nil)
	_ = data[0].Interface().([]map[string]interface{})
	start := time.Now()
	//Common.MySql{}.GetConn().Where(map[string]string{"dataType": spider.DataType}).Update("hotData2", map[string]string{"str": SaveDataToJson(originData)})
	group.Done()
	seconds := time.Since(start).Seconds()
	fmt.Printf("耗费 %.2fs 秒完成抓取%s", seconds, spider.DataType)
	fmt.Println()

}

var group sync.WaitGroup

func main() {
	redisutil.Start()
	cron.Start()
	//hotSearchService := service.HotSearchService{}
	//hotSearchService.DataType = "ZhiHu"
	//cmn.Info(util.JsonResponse(1, "成功", hotSearchService.GetZhiHu()))
	//allData := []string{
	//	"V2EX",
	//	"ZhiHu",
	//	"TieBa",
	//	"DouBan",
	//	"TianYa",
	//	"HuPu",
	//	"GitHub",
	//	"BaiDu",
	//	"36Kr",
	//	"QDaily",
	//	"GuoKr",
	//	"HuXiu",
	//	"ZHDaily",
	//	"Segmentfault",
	//	"WYNews",
	//	"WaterAndWood",
	//	"HacPai",
	//	"KD",
	//	"NGA",
	//	"WeiXin",
	//	"Mop",
	//	"Chiphell",
	//	"JianDan",
	//	"ChouTi",
	//	"ITHome",
	//}
	//fmt.Println("开始抓取" + strconv.Itoa(len(allData)) + "种数据类型")
	//group.Add(len(allData))
	//var spider Spider
	//for _, value := range allData {
	//	fmt.Println("开始抓取" + value)
	//	spider = Spider{DataType: value}
	//	go ExecGetData(spider)
	//}
	//group.Wait()
	//fmt.Print("完成抓取")
}

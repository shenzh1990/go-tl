package model

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/shenzh1990/TopList/pkg/settings"
	log "github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

var Db *xorm.Engine

func init() {
	var err error
	//打开数据库
	//DSN数据源字符串：用户名:密码@协议(地址:端口)/数据库?参数=参数值
	Db, err = xorm.NewEngine(settings.CommonConfig.Db.DriverName, settings.CommonConfig.Db.DBUrl)
	if err != nil {
		log.Fatal(err)
	}
	Db.ShowSQL(false)
	//tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, settings.BitConfig.Db.Prefix)
	//Db.SetTableMapper(tbMapper)
	//Db.Sync(new(User), new(Community), new(Contact))
	//Db.CreateTables(BaseLottery{})
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	//log.Info("init database successful")
}

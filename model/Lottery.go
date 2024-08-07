package model

import "time"

type BaseLottery struct {
	Id           int64     `xorm:"pk autoincr comment('id') INT(10)"`
	Issue        string    `xorm:" notnull unique comment('期号') VARCHAR(100)"`
	FrontNumbers string    `xorm:"  comment('前区号码 |号隔开') VARCHAR(128)"`
	EndNumbers   string    `xorm:"  comment('后区号码 |号隔开') VARCHAR(128)"`
	BingoDate    time.Time `xorm:"comment('日期')"`
	CreatedAt    time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
}

func (c *BaseLottery) AddLottery() bool {
	Db.Insert(c)
	return true
}
func (c *BaseLottery) GetAllLottery(nums int) (baseLottery []BaseLottery) {
	if nums <= 0 {
		Db.OrderBy("issue desc").Find(&baseLottery)
	} else {
		Db.Limit(nums, 0).OrderBy("issue desc").Find(&baseLottery)
	}

	return
}
func (c *BaseLottery) GetLottery(maps interface{}) (baseLottery []BaseLottery) {
	e := Db.Where(maps).Find(&baseLottery)
	if e != nil {
		panic(e)
	}
	return
}

package model

import "time"

const (
	SexWomen   = "F"
	SexMan     = "M"
	SexUnknown = "U"
)

type User struct {
	Id        int64     `xorm:"pk autoincr bigint(64)" form:"id" json:"id"`  //唯一id
	Mobile    string    `xorm:"varchar(20)" form:"mobile" json:"mobile"`     //手机号码
	Password  string    `xorm:"varchar(40)" form:"password" json:"-"`        // 用户密码
	Avatar    string    `xorm:"varchar(150)" form:"avatar" json:"avatar"`    //用户头像地址
	Sex       string    `xorm:"varchar(2)" form:"sex" json:"sex"`            //性别
	Nickname  string    `xorm:"varchar(20)" form:"nickname" json:"nickname"` //昵称
	Salt      string    `xorm:"varchar(10)" form:"salt" json:"-"`
	Online    int       `xorm:"int(10)" form:"online" json:"online"`   //是否在线
	Token     string    `xorm:"varchar(40)" form:"token" json:"token"` //用户鉴权
	Memo      string    `xorm:"varchar(140)" form:"memo" json:"memo"`  //备注
	CreatedAt time.Time `xorm:"created" json:"-"`
	UpdatedAt time.Time `xorm:"updated" json:"-"`
}

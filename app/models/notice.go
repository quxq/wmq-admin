package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"time"
)

type Notice struct {
	Id int64 `orm:"auto;pk;column(notice_id);"`
	UserName string
	Message string
	CreateTime int64
}

//table name
func (notice *Notice) TableName() string {
	return TableName("notice");
}

//获取所有的公告
func GetNotices() ([]*Notice) {
	notices := make([]*Notice, 0)
	query := orm.NewOrm().QueryTable(TableName("notice"))
	query.All(&notices);
	return notices;
}

//插入一条用户信息
func InsertNotice(notice *Notice) (int64, error) {
	if(notice.Message == "") {
		return 0, fmt.Errorf("%s", "公告内容不能为空!");
	}
	if(notice.UserName == "") {
		return 0, fmt.Errorf("%s", "没有登录!");
	}

	notice.CreateTime = time.Now().Unix();

	return orm.NewOrm().Insert(notice);
}
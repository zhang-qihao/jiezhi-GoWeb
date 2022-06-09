package models

import "time"

// PersonPost 定义个人投递情况数据表，标号为5
//Status, 简历的四种状态，0：未查看， 1：拒绝，2：待定，3：发起面试
type PersonPost struct {
	Pid        string    `json:"pid"`
	Jid        string    `json:"Jid"`
	Cid        string    `json:"cid"`
	Status     int       `json:"status"`
	Title      string    `json:"title"`
	Address    string    `json:"address"`
	Experience string    `json:"experience"`
	Education  string    `json:"education"`
	Salary     string    `json:"salary"`
	Name       string    `json:"name"`
	Age        int       `json:"age"`
	ViewTime   string    `json:"view_time"`
	ViewPlace  string    `json:"view_place"`
	ViewTel    string    `json:"view_tel"`
	CreatTime  time.Time `json:"createTime"`
	ChangeTime time.Time `json:"changeTime"`
}

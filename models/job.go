package models

import (
	"MySystem/dao"
	"time"
)

// JobInfos 定义职位的基础信息结构体，标号为4
type JobInfos struct {
	Jid        string    `json:"jid"`
	Title      string    `json:"title"`
	Company    string    `json:"company"`
	Cid        string    `json:"cid"`
	Nature     string    `json:"nature"`
	Salary     string    `json:"salary"`
	Education  string    `json:"education"`
	Experience string    `json:"experience"`
	Province   string    `json:"province"`
	Address    string    `json:"address"`
	Require    string    `json:"require"`
	Describe   string    `json:"describe"`
	CreatTime  time.Time `json:"createTime"`
	ChangeTime time.Time `json:"changeTime"`
}

func PublishAJob(job *JobInfos) (err error) {
	err = dao.DB.Create(job).Error
	return
}

func GetJobList() (jobList []JobInfos, err error) {
	err = dao.DB.Find(&jobList).Error
	return
}

func QueryJobList(query string) (jobList []JobInfos, err error) {
	err = dao.DB.Where("title like ?", "%"+query+"%").Find(&jobList).Error
	return
}

func GetJobDetail(jid string) (jobList []JobInfos, err error) {
	err = dao.DB.Where("jid = ?", jid).Find(&jobList).Error
	return
}

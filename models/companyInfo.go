package models

import (
	"MySystem/dao"
	"time"
)

type CompanyInfo struct {
	Cid         string    `json:"cid"`
	Shortname   string    `json:"shortname"`
	Fullname    string    `json:"fullName"`
	Companytype string    `json:"companyType"`
	Address     string    `json:"address"`
	Industry    string    `json:"industry"`
	Scale       string    `json:"scale"`
	Logourl     string    `json:"logoUrl"`
	Brief       string    `json:"brief"`
	Creattime   time.Time `json:"createTime"`
	Changetime  time.Time `json:"changeTime"`
}

func GetComopany() (company []CompanyInfo, err error) {
	err = dao.DB.Find(&company).Error
	return
}

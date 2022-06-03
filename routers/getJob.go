package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
	"time"
)

// JobInfo 定义职位的基础信息结构体，标号为4
type JobInfo struct {
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
	Creattime  time.Time `json:"createTime"`
	Changetime time.Time `json:"changeTime"`
}

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

//定义结构体用于存放职位信息
type Rank struct {
	Jid         string `json:"jid"`
	Title       string `json:"title"`
	Shortname   string `json:"shortname"`
	Fullname    string `json:"fullname"`
	Companytype string `json:"companyType"`
	Nature      string `json:"nature"`
	Salary      string `json:"salary"`
	Education   string `json:"education"`
	Experience  string `json:"experience"`
	Address     string `json:"address"`
	Industry    string `json:"industry"`
	Scale       string `json:"scale"`
	Describe    string `json:"describe"`
	Require     string `json:"require"`
	Brief       string `json:"brief"`
}

// 获取所有职业信息
func GetJobInfoFromDatabase(DB *gorm.DB, query string) []Rank {
	//new结构体
	rank1 := make([]Rank, 0)
	var job []JobInfo
	DB.Find(&job)
	var company []CompanyInfo
	DB.Find(&company)

	var (
		Jid1         string
		Title1       string
		Salary1      string
		Cid          string
		Companytype1 string
		Nature1      string
		Education1   string
		Experience1  string
		Address1     string
		Shortname1   string
		Industry1    string
		Scale1       string
	)
	var count = 0 //定义count用来计数

	//遍历数据库
	for _, value2 := range job {
		Jid1 = value2.Jid
		Title1 = value2.Title
		Salary1 = value2.Salary
		Education1 = value2.Education
		Experience1 = value2.Experience
		Address1 = value2.Address
		Nature1 = value2.Nature
		Cid = value2.Cid
		for _, value3 := range company {
			if value3.Cid == Cid {
				Shortname1 = value3.Shortname
				Industry1 = value3.Industry
				Companytype1 = value3.Companytype
				Scale1 = value3.Scale
			}
		}
		//new结构体
		var rank = Rank{
			Jid:         Jid1,
			Title:       Title1,
			Salary:      Salary1,
			Companytype: Companytype1,
			Nature:      Nature1,
			Education:   Education1,
			Experience:  Experience1,
			Shortname:   Shortname1,
			Address:     Address1,
			Industry:    Industry1,
			Scale:       Scale1,
		}
		if count < 50 {
			if query == "" {
				rank1 = append(rank1, rank)
				count++
			} else {
				if strings.Index(rank.Title, query) != -1 {
					rank1 = append(rank1, rank)
					count++
				}
			}
		}
	}
	return rank1
}

// 获取所有职业信息
func GetJobDetailFromDatabase(DB *gorm.DB, query string) []Rank {
	//new结构体
	rank1 := make([]Rank, 0)
	var job []JobInfo
	DB.Find(&job)
	var company []CompanyInfo
	DB.Find(&company)

	var (
		Jid1        string
		Title1      string
		Salary1     string
		Cid         string
		Nature1     string
		Fullname1   string
		Brief1      string
		Education1  string
		Experience1 string
		Address1    string
		Shortname1  string
		Industry1   string
		Scale1      string
		Describe1   string
		Require1    string
	)

	//遍历数据库
	for _, value2 := range job {
		Jid1 = value2.Jid
		Title1 = value2.Title
		Salary1 = value2.Salary
		Education1 = value2.Education
		Experience1 = value2.Experience
		Address1 = value2.Address
		Nature1 = value2.Nature
		Cid = value2.Cid
		Require1 = value2.Require
		Describe1 = value2.Describe
		for _, value3 := range company {
			if value3.Cid == Cid {
				Shortname1 = value3.Shortname
				Industry1 = value3.Industry
				Fullname1 = value3.Fullname
				Brief1 = value3.Brief
				Scale1 = value3.Scale
			}
		}
		//new结构体
		var rank = Rank{
			Jid:        Jid1,
			Title:      Title1,
			Salary:     Salary1,
			Nature:     Nature1,
			Education:  Education1,
			Experience: Experience1,
			Shortname:  Shortname1,
			Address:    Address1,
			Industry:   Industry1,
			Scale:      Scale1,
			Describe:   Describe1,
			Require:    Require1,
			Fullname:   Fullname1,
			Brief:      Brief1,
		}
		if strings.Index(rank.Jid, query) != -1 {
			rank1 = append(rank1, rank)
		}
	}
	return rank1
}

// InfoTransfer 向前端传值
func InfoTransfer(c *gin.Context, rank []Rank) {
	//向前端传json数据：20组职位信息，用于主页面显示
	c.JSON(http.StatusOK, gin.H{
		"Total": len(rank),
		"Jobs":  rank,
	})
}

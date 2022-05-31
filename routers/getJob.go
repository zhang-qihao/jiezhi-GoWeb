package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
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
	Salary      string `json:"salary"`
	Education   string `json:"education"`
	Experience  string `json:"experience"`
	Address     string `json:"address"`
	Industry 	string `json:"industry"`
	Scale       string `json:"scale"`
}

//
func GetJobInfoFromDatabase(DB *gorm.DB) []Rank {
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
		Cid 		 string
		Education1   string
		Experience1  string
		Address1     string
		Shortname1   string
		Industry1	 string
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
		Cid = value2.Cid
		for _, value3 := range company {
			if value3.Cid == Cid {
				Shortname1 = value3.Shortname
				Industry1 = value3.Industry
				Scale1 = value3.Scale
			}
		}
		//new结构体
		var rank = Rank{
			Jid:         Jid1,
			Title:       Title1,
			Salary:      Salary1,
			Education:   Education1,
			Experience:  Experience1,
			Shortname:   Shortname1,
			Address:     Address1,
			Industry: 	 Industry1,
			Scale:       Scale1,
		}
		count++
		if count <= 20 {
			rank1 = append(rank1, rank)
		}
		//else if count <= 21 {
		//	rank2 = append(rank2, rank)
		//} else if count <= 24 {
		//	rank3 = append(rank3, rank)
		//} else {
		//	//超过25就跳出循环
		//	break
		//}
	}
	return rank1
}

// InfoTransfer 向前端传值
func InfoTransfer(c *gin.Context, DB *gorm.DB) {
	//获取要放在主页的职位信息
	rank1 := GetJobInfoFromDatabase(DB)
	//向前端传json数据：20组职位信息，用于主页面显示
	c.JSON(http.StatusOK, gin.H{
		"Total": len(rank1),
		"Jobs":  rank1,
	})
}

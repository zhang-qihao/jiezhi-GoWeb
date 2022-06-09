package controller

import (
	"MySystem/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

// 获取职业信息
func GetJob(c *gin.Context) {
	var (
		err     error
		job     []models.JobInfos
		company []models.CompanyInfo
		rank1   []Rank
	)
	query := c.DefaultQuery("query", "")
	rank1 = make([]Rank, 0)
	if query == "" {
		job, err = models.GetJobList()
	} else {
		job, err = models.QueryJobList(query)
	}
	if err != nil {
		return
	}
	company, err = models.GetComopany()
	if err != nil {
		return
	}

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
			rank1 = append(rank1, rank)
			count++
		}
	}

	//向前端传json数据：职位信息、总数，用于主页面显示
	c.JSON(http.StatusOK, gin.H{
		"Total": len(rank1),
		"Jobs":  rank1,
	})
}

// 获取所有职业信息
func GetJobDetail(c *gin.Context) {
	var (
		err     error
		job     []models.JobInfos
		company []models.CompanyInfo
		rank1   []Rank
	)
	jid := c.DefaultQuery("jid", "")
	rank1 = make([]Rank, 0)
	job, err = models.GetJobDetail(jid)
	if err != nil {
		return
	}
	company, err = models.GetComopany()
	if err != nil {
		return
	}

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
		rank1 = append(rank1, rank)
	}

	//向前端传json数据：职位信息、总数，用于主页面显示
	c.JSON(http.StatusOK, gin.H{
		"Total": len(rank1),
		"Jobs":  rank1,
	})
}

func PublishAJob(c *gin.Context) {
	var job models.JobInfos
	err := c.BindJSON(&job)
	if err != nil {
		return
	}
	if err := models.PublishAJob(&job); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1001,
			"msg":  "发布工作失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1000,
			"msg":  "发布工作成功",
		})
	}
}

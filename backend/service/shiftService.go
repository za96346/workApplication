package service

import (
	"backend/methods"
	"backend/table"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// work time
func FetchWorkTime (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
	_, company, err := CheckUserAndCompany(props)
	if err {return}

	year, err1 := strconv.Atoi((*props).Query("year"))
	month, err2 := strconv.Atoi((*props).Query("month"))
	userId, err3 := methods.AnyToInt64((*props).Query("userId"))

	res :=  new([]table.WorkTime)
	if err3 != nil {
		res = (*Mysql).SelectWorkTime(2, year, month, company.CompanyCode)
	} else if err2 != nil && err1 != nil {
		res = (*Mysql).SelectWorkTime(1, userId, company.CompanyCode)
	} else {
		res = (*Mysql).SelectWorkTime(3, year, month, userId, company.CompanyCode)
	}

	(*props).JSON(http.StatusOK, gin.H{
		"data": (*res),
		"message": StatusText().FindSuccess,
	})
}
func InsertWorkTime (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
	_, company, err := CheckUserAndCompany(props)
	if err {return}
	now := time.Now()

	workTime :=  new(table.WorkTime)
	if (*props).ShouldBindJSON(&workTime) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	(*workTime).CreateTime = now
	(*workTime).LastModify = now

	res := (*Mysql).SelectUser(5, company.CompanyCode, (*workTime).UserId)
	if methods.IsNotExited(res) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return
	}

	if err1, _ := (*Mysql).InsertWorkTime(workTime); !err1 {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().InsertFail,
		})
		return
	}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().InsertSuccess,
	})
}
func DeleteWorkTime (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
	workTimeId := (*props).Query("workTimeId")
	convWorkTimeId, conErr := methods.AnyToInt64(workTimeId)
	if conErr != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}
	_, company, err := CheckUserAndCompany(props)
	if err {return}

	res := (*Mysql).DeleteWorkTime(1, convWorkTimeId, company.CompanyCode)
	if !res {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().DeleteFail,
		})
		return
	}
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().DeleteSuccess,
	})
}
func UpdateWorkTime (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
	_, company, err := CheckUserAndCompany(props)
	if err {return}
	now := time.Now()

	workTime := new(table.WorkTime)
	if (*props).ShouldBindJSON(&workTime) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}
	(*workTime).LastModify = now
	Log.Println("hi")

	if !(*Mysql).UpdateWorkTime(0, workTime, company.CompanyCode) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().UpdateFail,
		})
		return
	}
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().UpdateSuccess,
	})
}

// paid Vocation
func FetchPaidVocation (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
}
func InsertPaidVocation (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
}
func DeletePaidVocation (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
}
func UpdatePaidVocation (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()
}
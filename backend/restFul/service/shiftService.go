package service

import (
	// "backend/handler/shiftEdit"
	"backend/methods"
	"backend/mysql/table"
	"fmt"
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

	res :=  new([]table.WorkTimeExtend)
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

// 班表
func FetchMonthShift (props *gin.Context, waitJob *sync.WaitGroup)  {
	defer panicHandle()
	defer (*waitJob).Done()

	// 年度 月份 部門
	year, _ := strconv.Atoi((*props).Query("year"))
	month, _ := strconv.Atoi((*props).Query("month"))
	banch, _ := methods.AnyToInt64((*props).Query("banch"))

	// 自己的資料
	_, company, err := CheckUserAndCompany(props)
	if err {return}

	// 拿取資料
	data := (*Mysql).SelectShift(
		0,
		banch,
		company.CompanyId,
		year,
		month,
	)
	// 找開始結束時間
	ft := fmt.Sprintf("%d-0%d-01", year, month)
	if month >= 10 {
		ft = fmt.Sprintf("%d-%d-01", year, month)
	}
	ftime, _ := time.Parse("2006-01-02", ft)

	y, m, _ := ftime.Date()
	thisMonth := time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
	end := thisMonth.AddDate(0, 1, -1)

	rowShiftTotal := (*Mysql).SelectShiftRowTotal(0, banch, company.CompanyId, year, month)
	columnShiftTotal := (*Mysql).SelectShiftColumnTotal(0, banch, company.CompanyId, year, month)

	(*props).JSON(http.StatusOK, gin.H{
		"ShiftData": *data,
		"EditUser": RemoveDuplicate(
			data,
			func(T table.ShiftExtend) int64 {
				return T.ShiftTable.UserId
			},
		),	// 找到可編輯使用者
		"OnlineUser": nil,
		"StartDay": thisMonth.Format("2006-01-02"),
		"EndDay": end.Format("2006-01-02"),
		"message": StatusText().FindSuccess,
		"RowsShiftTotal": *rowShiftTotal, //列的總計
		"ColumnsShiftTotal": *columnShiftTotal, // 欄的總計
	})
}

// 班表總計
func FetchTotalShift(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	// 年度 月份 部門
	year, _ := strconv.Atoi((*props).Query("year"))
	month, _ := strconv.Atoi((*props).Query("month"))
	banch, _ := methods.AnyToInt64((*props).Query("banch"))

	// 自己的資料
	_, company, err := CheckUserAndCompany(props)
	if err {return}

	// 找尋
	data := (*Mysql).SelectShiftTotal(
		0,
		banch,
		company.CompanyId,
		year,
		month,
	)
	(*props).JSON(http.StatusOK, gin.H{
		"data": *data,
		"message": StatusText().FindSuccess,
	})
	
}

// 班表歷程
func FetchShiftHistory(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	// 年度 月份 部門
	year, _ := strconv.Atoi((*props).Query("year"))
	month, _ := strconv.Atoi((*props).Query("month"))
	banch, _ := methods.AnyToInt64((*props).Query("banch"))

	data := (*Mysql).SelectShiftEditLog(0, banch, year, month)

	(*props).JSON(http.StatusOK, gin.H{
		"data": *data,
		"message": StatusText().FindSuccess,
	})
}
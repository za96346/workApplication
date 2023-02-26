package service

import (
	"backend/methods"
	"backend/mysql/table"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// take manage  is admin taking the managers' data

func FetchPerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	startYearProps := (*props).Query("startYear")
	startMonthProps := (*props).Query("startMonth")
	endYearProps := (*props).Query("endYear")
	endMonthProps := (*props).Query("endMonth")
	name := (*props).Query("name")
	banchIdProps := (*props).Query("banchId")
	workState := (*props).Query("workState")
	takeManage := (*props).Query("takeManage") // "Y"


	banchId, isBanchError := methods.AnyToInt64(banchIdProps)

	if len(startMonthProps) == 1 {
		startMonthProps = fmt.Sprintf("0%s", startMonthProps)
	}
	if len(endMonthProps) == 1 {
		endMonthProps = fmt.Sprintf("0%s", endMonthProps)
	}

	start := fmt.Sprintf("%s%s", startYearProps, startMonthProps)
	end := fmt.Sprintf("%s%s", endYearProps, endMonthProps)

	Log.Println("start, end",start, "--", end)

	user, company, err := CheckUserAndCompany(props)
	if err {return}

	res := []table.PerformanceExtend{}
	// 管理元 沒帶 banchId
	if user.Permession == 100 && (isBanchError != nil || takeManage == "Y") {
		res = *(*Mysql).SelectPerformance(
			0,
			company.CompanyId,
			start,
			end,
			name,
			name,
			name,
		)
	// 管理元 有帶 banchId
	} else if user.Permession == 100 && isBanchError == nil {
		banch := (*Mysql).SelectCompanyBanch(2, banchId)
		if !methods.IsNotExited(banch) {
			res = *(*Mysql).SelectPerformance(
				1,
				company.CompanyId,
				banchId,
				(*banch)[0].BanchName,
				start,
				end,
				name,
				name,
				name,
			)
		}
	// 主管 有帶 banchId 拿部門成員
	} else if user.Permession == 1 && isBanchError == nil {
		banch := (*Mysql).SelectCompanyBanch(2, user.Banch)
		if !methods.IsNotExited(banch) {
			res = *(*Mysql).SelectPerformance(
				1,
				company.CompanyId,
				user.Banch,
				(*banch)[0].BanchName,
				start,
				end,
				name,
				name,
				name,
			)
		}
	// 一般職員 或事 主管自己
	} else {
		res = *(*Mysql).SelectPerformance(
			2,
			user.UserId,
			start,
			end,
		)
	}
	// 判斷工作狀態
	val := []table.PerformanceExtend{}
	for _, v := range res {
		if workState == "on" && v.CompanyId != -1 {
			val = append(val, v)
		} else if workState == "off" && v.CompanyId == -1 {
			val = append(val, v)
		}
	}
	props.JSON(http.StatusOK, gin.H{
		"message": StatusText().FindSuccess,
		"data": val,
	})
}
func UpdatePerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	now := time.Now()
	performance := table.Performance {}
	if (*props).ShouldBindJSON(&performance) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}
	performance.LastModify = now
	user, company, err := CheckUserAndCompany(props)
	if err {return}

	updateStatus := false

	if user.Permession == 100 {
		banch := (*Mysql).SelectCompanyBanch(2, performance.BanchId)
		banchName := performance.BanchName
		if !methods.IsNotExited(banch) {
			banchName = (*banch)[0].BanchName
		}
		performance.BanchName = banchName
		updateStatus = (*Mysql).UpdatePerformance(
			0,
			&performance,
			performance.PerformanceId,
			company.CompanyId,
		)
		// 主管更改不是自己的
	} else if user.Permession == 1 &&
		&performance.UserId != &user.UserId {

		banch := (*Mysql).SelectCompanyBanch(2, user.Banch)
		banchName := ""
		if !methods.IsNotExited(banch) {
			banchName = (*banch)[0].BanchName
		}
		updateStatus = (*Mysql).UpdatePerformance(
			1,
			&performance,
			performance.PerformanceId,
			company.CompanyId,
			user.Banch,
			banchName,
		)
		// 一般權限 更改 或是主管更改自己的
	} else if user.Permession == 2 ||
		(user.Permession == 1 &&
		&performance.UserId == &user.UserId) {

		updateStatus = (*Mysql).UpdatePerformance(
			2,
			&performance,
			performance.PerformanceId,
			user.UserId,
		)
	}

	if !updateStatus {
		props.JSON(http.StatusOK, gin.H{
			"message": StatusText().UpdateFail,
		})
		return
	}
	props.JSON(http.StatusOK, gin.H{
		"message": StatusText().UpdateSuccess,
	})
}
func InsertPerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	now := time.Now()
	performance := table.Performance {}
	if (*props).ShouldBindJSON(&performance) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}
	performance.CreateTime = now
	performance.LastModify = now

	user, _, err := CheckUserAndCompany(props)
	if err {return}

	// 擋掉 主管新增自己
	if performance.UserId == user.UserId && user.Permession == 1 {
		props.JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().InsertFail,
		})
		return
	}

	// 檢查 月份
	if performance.Month < 1 || performance.Month > 12 {
		props.JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().NoMonthSelect,
		})
		return
	}

	// 檢查年份
	if performance.Year == 0 {
		props.JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().NoYearSelect,
		})
		return
	}

	// 如果為 主管 新增 強制把 資料部門改為 自己的部門
	if user.Permession == 1 {
		performance.BanchId = user.Banch
	}

	// 巡找部門
	findBanch := (*Mysql).SelectCompanyBanch(2, performance.BanchId)
	if !methods.IsNotExited(findBanch) {
		performance.BanchName = (*findBanch)[0].BanchName
	} else {
		props.JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().NotHaveBanch,
		})
		return
	}
	res, _ := (*Mysql).InsertPerformance(&performance)

	if !res {
		props.JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().InsertFail,
		})
		return
	}
	props.JSON(http.StatusOK, gin.H{
		"message": StatusText().InsertSuccess,
	})
}
func DeletePerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	performanceId, isIdErr := methods.AnyToInt64((*props).Query("performanceId"))

	if isIdErr != nil {
		props.JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	user, _, err := CheckUserAndCompany(props)
	if err {return}

	res := false
	if user.Permession == 100 {
		res = (*Mysql).DeletePerformance(0, performanceId)
	} else if user.Permession == 1 {
		res = (*Mysql).DeletePerformance(1, performanceId, user.Banch, user.UserId)
	}

	if !res {
		props.JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().DeleteFail,
		})
		return
	}
	props.JSON(http.StatusOK, gin.H{
		"message": StatusText().DeleteSuccess,
	})

}

// 複製功能
func CopyPerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	performance := new([]table.PerformanceExtend)

	request := new(struct {
		PerformanceId int64 `json:"PerformanceId"`
		IsResetGrade bool `json:"IsResetGrade"` // 重設 分數
		IsResetDirections bool `json:"IsResetDirections"` // 重設績效描述
	})
	
	if (*props).ShouldBindJSON(request) != nil {
		props.JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}
	user, company, err := CheckUserAndCompany(props)
	if err {return}

	// 依據權限選擇
	if user.Permession == 2 {
		performance = (*Mysql).SelectPerformance(
			5,
			request.PerformanceId,
			user.UserId,
		)
	} else if user.Permession == 1 {
		myBanch := (*Mysql).SelectCompanyBanch(2, user.Banch)
		if methods.IsNotExited(myBanch) {
			props.JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().NotHaveBanch,
			})
			return
		}
		performance = (*Mysql).SelectPerformance(
			4,
			request.PerformanceId,
			company.CompanyId,
			user.Banch,
			(*myBanch)[0].BanchName,
		)
	} else if user.Permession == 100 {
		performance = (*Mysql).SelectPerformance(
			3,
			request.PerformanceId,
			company.CompanyId,
		)
	}
	
	// 檢查是否有單筆的資料
	if methods.IsNotExited(performance) {
		props.JSON(http.StatusNotFound, gin.H{
			"message": StatusText().NoData,
		})
		return
	}
	
	// 設定新的資料
	if request.IsResetGrade {
		(*performance)[0].Attitude = 0
		(*performance)[0].Efficiency = 0
		(*performance)[0].Professional = 0
		(*performance)[0].BeLate = 0
		(*performance)[0].DayOffNotOnRule = 0
	}
	if request.IsResetDirections {
		(*performance)[0].Directions = ""
	}

	if (*performance)[0].Month == 12 {
		(*performance)[0].Year += 1
		(*performance)[0].Month = 1
	} else {
		(*performance)[0].Month += 1
	}

	now := time.Now()
	// 轉換 struct
	newP := table.Performance {
		UserId: (*performance)[0].UserId,
		Year: (*performance)[0].Year,
		Month: (*performance)[0].Month,
		BanchId: (*performance)[0].BanchId,
		Goal: (*performance)[0].Goal,
		Attitude: (*performance)[0].Attitude,
		Efficiency: (*performance)[0].Efficiency,
		Professional: (*performance)[0].Professional,
		Directions: (*performance)[0].Directions,
		BeLate: (*performance)[0].BeLate,
		DayOffNotOnRule: (*performance)[0].DayOffNotOnRule,
		BanchName: (*performance)[0].BanchName,
		CreateTime: now,
		LastModify: now,
	}

	if res, _ := (*Mysql).InsertPerformance(&newP); !res {
		props.JSON(http.StatusConflict, gin.H{
			"message": StatusText().CopyFail,
		})
		return
	}
	props.JSON(http.StatusOK, gin.H{
		"message": StatusText().CopySuccess,
	})
}

// 年度績效
func FetchYearPerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	// 拿取參數
	startYear := (*props).Query("startYear")
	endYear := (*props).Query("endYear")
	userName := (*props).Query("userName")

	user, company, err := CheckUserAndCompany(props)
	if err {return}

	res := new([]table.YearPerformance)
	// 根據 權限選擇
	if user.Permession == 100 {
		res = (*Mysql).SelectYearPerformance(
			0,
			company.CompanyId,
			startYear,
			endYear,
			userName,
			userName,
			userName,
		)
	} else if user.Permession == 1 {
		myBanch := (*Mysql).SelectCompanyBanch(2, user.Banch)
		if methods.IsNotExited(myBanch) {
			props.JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().NotHaveBanch,
			})
			return
		}
		res = (*Mysql).SelectYearPerformance(
			1,
			company.CompanyId,
			(*myBanch)[0].Id,
			(*myBanch)[0].BanchName,
			startYear,
			endYear,
			userName,
			userName,
			userName,
		)
	} else if user.Permession == 2 {
		res = (*Mysql).SelectYearPerformance(
			2,
			user.UserId,
			startYear,
			endYear,
			userName,
			userName,
			userName,
		)
	}
	(*props).JSON(http.StatusOK, gin.H {
		"message": "success",
		"data": *res,
	})
}
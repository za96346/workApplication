package service

import (
	"backend/methods"
	"backend/table"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func FetchPerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	startYearProps := (*props).Query("startYear")
	startMonthProps := (*props).Query("startMonth")
	endYearProps := (*props).Query("endYear")
	endMonthProps := (*props).Query("endMonth")
	name := (*props).Query("name")
	banchIdProps := (*props).Query("banchId")

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

	user, _, err := CheckUserAndCompany(props)
	if err {return}

	res := []table.PerformanceExtend{}
	// 管理元 沒帶 banchId
	if user.Permession == 100 && isBanchError != nil {
		res = *(*Mysql).SelectPerformance(
			0,
			user.CompanyCode,
			user.CompanyCode,
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
				user.CompanyCode,
				user.CompanyCode,
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
				user.CompanyCode,
				user.CompanyCode,
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
	props.JSON(http.StatusOK, gin.H{
		"message": "not bad",
		"data": res,
	})
}
func UpdatePerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	props.JSON(http.StatusOK, gin.H{
		"message": "not bad",
	})
}
func InsertPerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	props.JSON(http.StatusOK, gin.H{
		"message": "not bad",
	})
}
func DeletePerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	props.JSON(http.StatusOK, gin.H{
		"message": "not bad",
	})
}
package service

import (
	"backend/methods"
	"backend/table"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

func FetchPerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	yearProps := (*props).Query("year")
	monthProps := (*props).Query("month")
	banchIdProps := (*props).Query("banchId")

	year, _ := strconv.Atoi(yearProps)
	month, _ := strconv.Atoi(monthProps)
	banchId, isBanchError := methods.AnyToInt64(banchIdProps)

	user, _, err := CheckUserAndCompany(props)
	if err {return}

	res := []table.PerformanceExtend{}
	// 管理元 沒帶 banchId
	if user.Permession == 100 && isBanchError != nil {
		res = *(*Mysql).SelectPerformance(
			0,
			user.CompanyCode,
			user.CompanyCode,
			year,
			month,
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
				year,
				month,
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
				year,
				month,
			)
		}
	// 一般職員 或事 主管自己
	} else {
		res = *(*Mysql).SelectPerformance(2, user.UserId, year, month)
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
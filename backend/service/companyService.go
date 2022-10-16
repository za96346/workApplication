package service

import (
	"fmt"
	"net/http"
	"sync"
	"backend/methods"

	"github.com/gin-gonic/gin"
)

func FetchBanchAll(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	userId, existed := (*props).Get("UserId")


	// user id 是否存在
	if !existed {
		(*props).JSON(http.StatusInternalServerError, gin.H{
			"message": StatusText().UserIdNotFound,
		})
		return
	}

	// 轉換 user id
	convertUserId, err := methods.AnyToInt64(userId)
	fmt.Print(convertUserId, err)
	if err != nil {
		(*props).JSON(http.StatusInternalServerError, gin.H{
			"message": StatusText().UserIdNotFound + "2",
		})
		return
	}

	// 尋找資料
	res := (*dbHandle).SelectUser(1, convertUserId)
	if IsNotExited(res) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}

	// 尋找公司
	company := (*dbHandle).SelectCompany(2, (*res)[0].CompanyCode)
	if IsNotExited(company) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return
	}

	banch := (*dbHandle).SelectCompanyBanch(1, (*company)[0].CompanyId)
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().FindSuccess,
		"data": banch,
	})
}

func FetchBanchStyle(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	userId, existed := (*props).Get("UserId")
	wantBanchId := (*props).Query("banchId")

	// 轉換banch id
	convertBanchId, err := methods.AnyToInt64(wantBanchId)
	if err != nil {
		(*props).JSON(http.StatusInternalServerError, gin.H{
			"message": StatusText().BanchIdIsNotRight,
		})
		return
	}

	// user id 是否存在
	if !existed {
		(*props).JSON(http.StatusInternalServerError, gin.H{
			"message": StatusText().UserIdNotFound,
		})
		return
	}
	// 轉換 user id
	convertUserId, err := methods.AnyToInt64(userId)
	// fmt.Print(convertUserId, err)
	if err != nil {
		(*props).JSON(http.StatusInternalServerError, gin.H{
			"message": StatusText().UserIdNotFound + "2",
		})
		return
	}
	
	// 尋找自己資料
	res := (*dbHandle).SelectUser(1, convertUserId)
	if IsNotExited(res) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}

	// 尋找公司
	company := (*dbHandle).SelectCompany(2, (*res)[0].CompanyCode)
	if IsNotExited(res) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return
	}

	// 檢查部門是否是公司本身
	companyBanch := (*dbHandle).SelectCompanyBanch(1, (*company)[0].CompanyId)
	for _, v := range *companyBanch {

		if convertBanchId == v.Id {
			res := (*dbHandle).SelectBanchStyle(2, convertBanchId)
			(*props).JSON(http.StatusOK, gin.H{
				"message": StatusText().FindSuccess,
				"data": res,
			})
			return
		}
	}

	// 沒有此部門 或 不是此公司的部門
	(*props).JSON(http.StatusNotAcceptable, gin.H{
		"message": StatusText().BanchIdIsNotRight,
		"data": make([]int, 0),
	})
	return
	
}
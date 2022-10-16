package service

import (
	"backend/methods"
	"backend/table"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func FetchBanchAll(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	myCompany, exited := (*props).Get("MyCompany")
	if !exited {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return
	}

	switch v := myCompany.(type) {
	case *[]table.CompanyTable:
		banch := (*dbHandle).SelectCompanyBanch(1, (*v)[0].CompanyId)
		(*props).JSON(http.StatusOK, gin.H{
			"message": StatusText().FindSuccess,
			"data": banch,
		})
	default:
		if !exited {
			(*props).JSON(http.StatusNotFound, gin.H{
				"message": StatusText().IsNotHaveCompany,
			})
			return
		}
	}
	
}

func FetchBanchStyle(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	banchId := (*props).Query("banchId")
	convertBanchId, err := methods.AnyToInt64(banchId)
	if err != nil {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().BanchIdIsNotRight,
		})
		return
	}

	myCompany, exited := (*props).Get("MyCompany")
	if !exited {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return
	}

	switch types := myCompany.(type) {
	case *[]table.CompanyTable:
			// 檢查部門是否是公司本身
		companyBanch := (*dbHandle).SelectCompanyBanch(1, (*types)[0].CompanyId)
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
	default:
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return
	}
	
}

func FetchBanchRule(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	banchId := (*props).Query("banchId")
	convertBanchId, err := methods.AnyToInt64(banchId)
	if err != nil {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().BanchIdIsNotRight,
		})
		return
	}

	company, exited := (*props).Get("MyCompany")
	if !exited {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return
	}

	switch types := company.(type) {
	case *[]table.CompanyTable:
		companyBanch := (*dbHandle).SelectCompanyBanch(1, (*types)[0].CompanyId)
		for _, v := range *companyBanch {
			if convertBanchId == v.Id {
				res := (*dbHandle).SelectBanchRule(2, convertBanchId)
				(*props).JSON(http.StatusOK, gin.H{
					"message": StatusText().FindSuccess,
					"data": res,
				})
				return
			}
		}

		// 沒有此部門 或 不是此公司的部門
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().BanchIdIsNotRight,
		})
		return
	default:
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return
	}

}
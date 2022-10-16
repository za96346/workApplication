package service

import (
	"backend/methods"
	"backend/table"
	"net/http"
	"sync"
	"time"

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
	case table.CompanyTable:
		banch := (*dbHandle).SelectCompanyBanch(1, v.CompanyId)
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

// banch style
func FetchBanchStyle(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	// 檢查 banch id 
	banchId := (*props).Query("banchId")
	convertBanchId, err := methods.AnyToInt64(banchId)
	if err != nil {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().BanchIdIsNotRight,
		})
		return
	}

	// 檢查 my company 存在
	myCompany, exited := (*props).Get("MyCompany")
	if !exited {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return
	}

	// 轉換 my company
	company, a := methods.Assertion[table.CompanyTable](myCompany)
	if !a {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().AssertionFail,
		})
		return
	}

	// 檢查 部門是否在此公司
	if !BanchIsInCompany(convertBanchId, company.CompanyId) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().NotHaveBanch,
		})
		return
	}

	// 部門合法
	res := (*dbHandle).SelectBanchStyle(2, convertBanchId)
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().FindSuccess,
		"data": res,
	})
	
}

func UpdateBanchStyle(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	// 檢查body
	banchStyle := table.BanchStyle{
		LastModify: time.Now(),
	}

	if (*props).ShouldBindJSON(&banchStyle) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	// 檢查是否有此style id
	res := (*dbHandle).SelectBanchStyle(1, (*&banchStyle).StyleId)
	if methods.IsNotExited(res) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().StyleIdNotRight,
		})
		return
	}

	// 添加 banch id
	banchStyle.BanchId = (*res)[0].BanchId

	// 檢查是否是 company table type
	myCompany, exited := (*props).Get("MyCompany")
	if !exited {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return
	}
	company, a := methods.Assertion[table.CompanyTable](myCompany)
	if !a {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().AssertionFail,
		})
		return
	}

	// 檢查是否是 user table type
	myUserData, exited := (*props).Get("MyUserData")
	if !exited {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}
	user, a := methods.Assertion[table.UserTable](myUserData)
	if !a {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().AssertionFail,
		})
		return
	}

	// 檢查 部門是否在此公司
	if !BanchIsInCompany(banchStyle.BanchId, company.CompanyId) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().NotHaveBanch,
		})
		return
	}

	// 最高權限 更新
	if user.Permession == 100 {
		// 都可以更新
		if !(*dbHandle).UpdateBanchStyle(0, &banchStyle) {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().UpdateFail,
			})
			return
		}
	}
	// 主管權限 更新
	if user.Permession == 1 {
		// 只能更新自己的部門
		if  user.Banch != banchStyle.BanchId {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().OnlyCanUpDateYourBanch,
			})
			return
		}
		if !(*dbHandle).UpdateBanchStyle(0, &banchStyle) {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().UpdateFail,
			})
			return
		}
	}

	// 更新成功
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().UpdateSuccess,
	})
}

func InsertBanchStyle(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	now := time.Now()
	banchStyle := table.BanchStyle{
		CreateTime: now,
		LastModify: now,
	}

	if (*props).ShouldBindJSON(&banchStyle) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	// 檢查是否是 company table type
	myCompany, exited := (*props).Get("MyCompany")
	if !exited {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return
	}
	company, a := methods.Assertion[table.CompanyTable](myCompany)
	if !a {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().AssertionFail,
		})
		return
	}
	
	// 檢查是否是 user table type
	myUserData, exited := (*props).Get("MyUserData")
	if !exited {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}
	user, a := methods.Assertion[table.UserTable](myUserData)
	if !a {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().AssertionFail,
		})
		return
	}
	
	// 檢查 部門是否在此公司
	if !BanchIsInCompany(banchStyle.BanchId, company.CompanyId) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().NotHaveBanch,
		})
		return
	}

	// 最高權限 更新
	if user.Permession == 100 {
		if  v, _ := (*dbHandle).InsertBanchStyle(&banchStyle); !v {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().InsertFail,
			})
			return
		}
	}

	// 主管權限 更新
	if user.Permession == 1 {
		if user.Banch != banchStyle.BanchId {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().OnlyCanUpDateYourBanch,
			})
			return
		}
		if v, _ := (*dbHandle).InsertBanchStyle(&banchStyle); !v {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().InsertFail,
			})
			return
		}
	}


	// 更新成功
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().InsertSuccess,
	})
}


// banch rule
func FetchBanchRule(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	// 檢查 banch id 
	banchId := (*props).Query("banchId")
	convertBanchId, err := methods.AnyToInt64(banchId)
	if err != nil {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().BanchIdIsNotRight,
		})
		return
	}

	// 檢查 my company 存在
	myCompany, exited := (*props).Get("MyCompany")
	if !exited {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return
	}

	// 轉換 mycompany
	company, a := methods.Assertion[table.CompanyTable](myCompany)
	if !a {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().AssertionFail,
		})
		return
	}


	// 檢查 部門是否在此公司
	if !BanchIsInCompany(convertBanchId, company.CompanyId) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().NotHaveBanch,
		})
		return
	}

	// 部門合法
	res := (*dbHandle).SelectBanchRule(2, convertBanchId)
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().FindSuccess,
		"data": res,
	})

}

func UpdateBanchRule(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	// 檢查body
	banchRule := table.BanchRule{
		LastModify: time.Now(),
	}
	if (*props).ShouldBindJSON(&banchRule) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	res := (*dbHandle).SelectBanchRule(1, banchRule.RuleId)
	if methods.IsNotExited(res) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().RuleIdIsNotRight,
		})
		return
	}

	// 添加 banch id
	banchRule.BanchId = (*res)[0].BanchId

	// 檢查是否是 company table type
	myCompany, exited := (*props).Get("MyCompany")
	if !exited {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return
	}
	company, a := methods.Assertion[table.CompanyTable](myCompany)
	if !a {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().AssertionFail,
		})
		return
	}
	
	// 檢查是否是 user table type
	myUserData, exited := (*props).Get("MyUserData")
	if !exited {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}
	user, a := methods.Assertion[table.UserTable](myUserData)
	if !a {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().AssertionFail,
		})
		return
	}

	// 檢查 部門是否在此公司
	if !BanchIsInCompany(banchRule.BanchId, company.CompanyId) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().NotHaveBanch,
		})
		return
	}
	
	// 最高權限 更新
	if user.Permession == 100 {
		// 都可以更新
		if !(*dbHandle).UpdateBanchRule(0, &banchRule) {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().UpdateFail,
			})
			return
		}
	}

	// 主管權限 更新
	if user.Permession == 1 {
		if user.Banch != banchRule.BanchId {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().OnlyCanUpDateYourBanch,
			})
			return
		}
		if !(*dbHandle).UpdateBanchRule(0, &banchRule) {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().UpdateFail,
			})
			return
		}
	}

	// 更新成功
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().UpdateSuccess,
	})
}
func InsertBanchRule(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	
}
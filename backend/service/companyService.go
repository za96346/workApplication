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

func InsertBanch(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	now := time.Now()
	banch := table.CompanyBanchTable {}
	if (*props).ShouldBindJSON(&banch) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	} 

	_, company, err := CheckUserAndCompany(props)
	if err {return}

	banch.LastModify = now
	banch.CompanyId = company.CompanyId
	banch.BanchShiftStyle = ""
	banch.CreateTime = now

	// 新增
	res, _ := (*dbHandle).InsertCompanyBanch(&banch)
	if !res {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().InsertFail,
		})
		return
	}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().InsertSuccess,
	})
}

func UpdateBanch (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	now := time.Now()
	banch := table.CompanyBanchTable {}
	if (*props).ShouldBindJSON(&banch) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}
	_, company, err := CheckUserAndCompany(props)
	if err {return}

	banch.LastModify = now
	banch.CompanyId = company.CompanyId
	banch.BanchShiftStyle = ""

	// 檢查 部門是否在此公司
	if !BanchIsInCompany(banch.Id, company.CompanyId) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().NotHaveBanch,
		})
		return
	}

	res := (*dbHandle).UpdateCompanyBanch(0, &banch)
	if !res {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().UpdateFail,
		})
		return
	}

	(*props).JSON(http.StatusNotAcceptable, gin.H{
		"message": StatusText().UpdateSuccess,
	})
}

// banch style
func FetchBanchStyle(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	// 檢查 banch id 
	banchId := (*props).Query("banchId")
	convertBanchId, err1 := methods.AnyToInt64(banchId)
	if err1 != nil {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().BanchIdIsNotRight,
		})
		return
	}

	_, company, err := CheckUserAndCompany(props)
	if err {return}

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

	user, company, err := CheckUserAndCompany(props)
	if err {return}

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

	user, company, err := CheckUserAndCompany(props)
	if err {return}
	
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


	// 新增成功
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().InsertSuccess,
	})
}

func DeleteBanchStyle(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	request := table.BanchStyle{}
	if (*props).ShouldBindJSON(&request) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	// 拿取該部門資料
	banchStyle := (*dbHandle).SelectBanchStyle(1, request.StyleId)
	if methods.IsNotExited(banchStyle) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().DeleteFail + "," + StatusText().NotHaveBanch,
		})
		return
	}

	user, company, err := CheckUserAndCompany(props)
	if err {return}

	// 檢查 部門是否在此公司
	if !BanchIsInCompany((*banchStyle)[0].BanchId, company.CompanyId) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().NotHaveBanch,
		})
		return
	}


	//最高權限刪除
	if user.Permession == 100 {
		if v := (*dbHandle).DeleteBanchStyle(0, request.StyleId); !v {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().DeleteFail,
			})
			return
		} 
	}

	// 主管權限刪除
	if user.Permession == 1 {
		if user.Banch != (*banchStyle)[0].BanchId {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().OnlyCanDeleteYourBanch,
			})
			return
		}
		if v := (*dbHandle).DeleteBanchStyle(0, request.StyleId); !v {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().DeleteFail,
			})
			return
		}
	}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().DeleteSuccess,
	})
	return

}


// banch rule
func FetchBanchRule(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	// 檢查 banch id 
	banchId := (*props).Query("banchId")
	convertBanchId, err1 := methods.AnyToInt64(banchId)
	if err1 != nil {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().BanchIdIsNotRight,
		})
		return
	}

	_, company, err := CheckUserAndCompany(props)
	if err {return}


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

	user, company, err := CheckUserAndCompany(props)
	if err {return}

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
	now := time.Now()
	banchRule := table.BanchRule{
		CreateTime: now,
		LastModify: now,
	}
	if (*props).ShouldBindJSON(&banchRule) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	user, company, err := CheckUserAndCompany(props)
	if err {return}
	
	// 檢查 部門是否在此公司
	if !BanchIsInCompany(banchRule.BanchId, company.CompanyId) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().NotHaveBanch,
		})
		return
	}

	// 最高權限 更新
	if user.Permession == 100 {
		if v,_ := (*dbHandle).InsertBanchRule(&banchRule); !v {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().InsertFail,
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
		if v, _ := (*dbHandle).InsertBanchRule(&banchRule); !v {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().InsertFail,
			})
			return
		}
	}
	// 新增成功
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().InsertSuccess,
	})
}

func DeleteBanchRule(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
}

func FetchCompany (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	_, company, err := CheckUserAndCompany(props)
	if err {return}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().FindSuccess,
		"data": company,
	})
}

func InsertCompany (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
}

func UpdateCompany (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	request := table.CompanyTable{}
	if (*props).ShouldBindJSON(&request) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}
	request.LastModify = time.Now()

	_, company, err := CheckUserAndCompany(props)
	if err {return}
	// 判斷是不是同一家公司
	if request.CompanyId != company.CompanyId {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().CompanyNotEqual,
		})
		return
	}

	res := (*dbHandle).UpdateCompany(0, &request)
	if !res {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().UpdateFail,
		})
		return
	}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().UpdateSuccess,
	})
	return

}
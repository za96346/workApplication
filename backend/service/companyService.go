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

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().UpdateSuccess,
	})
}

func DeleteBanch(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

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

	if !(*dbHandle).DeleteCompanyBanch(0, convertBanchId) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().DeleteFail,
		})
		return
	}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().DeleteSuccess,
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

	RequestStyleId := (*props).Query("StyleId")
	StyleId, convErr := methods.AnyToInt64(RequestStyleId)
	if convErr != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	// 拿取該部門資料
	banchStyle := (*dbHandle).SelectBanchStyle(1, StyleId)
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
		if !(*dbHandle).DeleteBanchStyle(0, StyleId) {
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
		if !(*dbHandle).DeleteBanchStyle(0, StyleId) {
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
	RequestRuleId := (*props).Query("RuleId")
	RuleId, convErr := methods.AnyToInt64(RequestRuleId)
	if convErr != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	banchRule := (*dbHandle).SelectBanchRule(1, RuleId)
	if methods.IsNotExited(banchRule) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().DeleteFail + "," + StatusText().NotHaveBanch,
		})
		return
	}

	user, company, err := CheckUserAndCompany(props)
	if err {return}
	
	// 檢查 部門是否在此公司
	if !BanchIsInCompany((*banchRule)[0].BanchId, company.CompanyId) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().NotHaveBanch,
		})
		return
	}

	// 最高權限 刪除
	if user.Permession == 100 {
		if !(*dbHandle).DeleteBanchRule(0, RuleId) {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().DeleteFail,
			})
			return
		}
	}

	// 主管權限 刪除
	if user.Permession == 1 {
		if user.Banch != (*banchRule)[0].BanchId {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().OnlyCanDeleteYourBanch,
			})
			return
		}
		if !(*dbHandle).DeleteBanchRule(0, RuleId) {
			(*props).JSON(http.StatusNotAcceptable, gin.H{
				"message": StatusText().DeleteFail,
			})
			return
		}
	}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().DeleteSuccess,
	})
}

func FetchCompany (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	company := &[]table.CompanyTable{}

	companyCode := (*props).Query("companyCode")
	companyId, err := methods.AnyToInt64((*props).Query("companyId"))

	if companyCode != "" {
		company = (*dbHandle).SelectCompany(2, companyCode)
	} else if err == nil {
		company = (*dbHandle).SelectCompany(1, companyId)
	}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().FindSuccess,
		"data": *company,
	})
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

	user, company, err := CheckUserAndCompany(props)
	if err {return}

	// 判斷是不是公司負責人
	if company.BossId != user.UserId {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().YouAreNotBoss,
		})
		return
	}

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

func InsertCompany (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	userId, status := checkMineUserId(props)
	if !status {return}

	company := table.CompanyTable{}
	if (*props).ShouldBindJSON(&company) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	// 公司碼的長度
	if len(company.CompanyCode) < 10 {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().CompanyCodeIsNotTenLength,
		})
		return
	}

	// 更改company 的欄位

	now := time.Now()
	company.BossId = userId
	company.CreateTime = now
	company.LastModify = now
	
	if v, _ := (*dbHandle).InsertCompany(&company); !v {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().InsertFail,
		})
		return
	}

	// 更改負責人的資料

	user := (*dbHandle).SelectUser(1, userId)
	if methods.IsNotExited(user) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}
	(*user)[0].Banch = -1
	(*user)[0].Permession = 100
	(*user)[0].CompanyCode = company.CompanyCode
	(*user)[0].LastModify = now


	if !(*dbHandle).UpdateUser(0, &(*user)[0]) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": "更新負責人資料失敗",
		})
		return
	}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().InsertSuccess,
	})
}

func FetchWaitReply (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	user, _, err := CheckUserAndCompany(props)
	if err {return}


	if user.Permession == 100 {
		company := (*dbHandle).SelectCompany(2, user.CompanyCode)
		waitCompanyReply := &[]table.WaitCompanyReply{}
		if !methods.IsNotExited(company) {
			waitCompanyReply = (*dbHandle).Mysql.SelectWaitCompanyReply(5, (*company)[0].CompanyId)
		}
		props.JSON(http.StatusOK, gin.H{
			"data": *waitCompanyReply,
			"message": StatusText().FindSuccess,
		})
	} else {
		// 只能拿自己的
		waitCompanyReply := (*dbHandle).SelectWaitCompanyReply(2, user.UserId)
		if !methods.IsNotExited(waitCompanyReply) {
			(*waitCompanyReply)[0].UserName = user.UserName
		}
		props.JSON(http.StatusOK, gin.H{
			"data": *waitCompanyReply,
			"message": StatusText().FindSuccess,
		})
	}
}

func UpdateWaitCompanyReply (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	// 檢查格式
	request := table.WaitCompanyReply{}
	if (*props).ShouldBindJSON(&request) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	_, company, err := CheckUserAndCompany(props)
	if err {return}

	waitCompanyReply := (*dbHandle).SelectWaitCompanyReply(1, request.WaitId)
	if methods.IsNotExited(waitCompanyReply) {
		(*props).JSON(http.StatusBadRequest, gin.H{
			"message": StatusText().NoData,
		})
		return
	}
	if company.CompanyId != (*waitCompanyReply)[0].CompanyId {
		(*props).JSON(http.StatusBadRequest, gin.H{
			"message": StatusText().CompanyNotEqual,
		})
		return
	}

	request.LastModify = time.Now()
	if !(*dbHandle).UpdateWaitCompanyReply(0, &request) {
		(*props).JSON(http.StatusForbidden, gin.H{
			"message": StatusText().UpdateFail,
		})
		return
	}

	// 進入 公司
	if request.IsAccept == 2 {
		targetUser := (*dbHandle).SelectUser(1, (*waitCompanyReply)[0].UserId)
		if methods.IsNotExited(targetUser) {
			(*props).JSON(http.StatusBadRequest, gin.H{
				"message": StatusText().userDataNotFound,
			})
			return
		}
		(*targetUser)[0].CompanyCode = company.CompanyCode
		if !(*dbHandle).UpdateUser(0, &(*targetUser)[0]) {
			(*props).JSON(http.StatusForbidden, gin.H{
				"message": StatusText().UpdateFail,
			})
			return
		}
	}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().UpdateSuccess,
	})

}

func InsertWaitCompanyReply (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	now := time.Now()

	// 檢查格式
	request := table.UserTable{}
	if (*props).ShouldBindJSON(&request) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	user, _, err := CheckUserAndCompany(props)
	if err {return}

	findCompany := (*dbHandle).SelectCompany(2, request.CompanyCode)
	if methods.IsNotExited(findCompany) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().CompanyCodeIsNotRight,
		})
		return
	}

	waitCompanyReply := table.WaitCompanyReply{
		UserId: user.UserId,
		CompanyId: (*findCompany)[0].CompanyId,
		SpecifyTag: "您好我叫" + user.UserName,
		IsAccept: 1,
		CreateTime: now,
		LastModify: now,
	}
	res, _ := (*dbHandle).InsertWaitCompanyReply(&waitCompanyReply)
	if !res {
		(*props).JSON(http.StatusConflict, gin.H{
			"message": StatusText().InsertFail,
		})
		return
	}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().InsertSuccess,
	})
}

func FetchWeekendSetting (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	_, company, err := CheckUserAndCompany(props)
	if err {return}

	weekendSetting := (*dbHandle).SelectWeekendSetting(2, company.CompanyId)
	props.JSON(http.StatusOK, gin.H{
		"message": StatusText().FindSuccess,
		"data": *weekendSetting,
	})
}

func InsertWeekendSetting (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	now := time.Now()
	// 檢查格式
	request := table.WeekendSetting{}
	if (*props).ShouldBindJSON(&request) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	_, company, err := CheckUserAndCompany(props)
	if err {return}

	request.CreateTime = now
	request.LastModify = now
	request.CompanyId = company.CompanyId

	if v, _:= (*dbHandle).InsertWeekendSetting(&request); !v {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().InsertFail,
		})
		return
	}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().InsertSuccess,
	})

}

func DeleteWeekendSetting (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	// 檢查weekend id
	weekendId := (*props).Query("weekendId")
	convertWeekendId, err := methods.AnyToInt64(weekendId)
	if err != nil {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().WeekendIdNotRight,
		})
		return
	}

	// 判斷是否在公司
	findCompanyWeekend := (*dbHandle).SelectWeekendSetting(1, convertWeekendId)

	_, company, err1 := CheckUserAndCompany(props)
	if err1 {return}
	if !methods.IsNotExited(findCompanyWeekend) {
		if (*findCompanyWeekend)[0].CompanyId != company.CompanyId {
			(*props).JSON(http.StatusNotFound, gin.H{
				"message": StatusText().CompanyNotEqual,
			})
			return
		}
	}

	// 刪除
	if !(*dbHandle).DeleteWeekendSetting(0, convertWeekendId) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().DeleteFail,
		})
		return
	}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().DeleteSuccess,
	})
}
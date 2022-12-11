package service

import (
	"backend/methods"
	"backend/table"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)
// banch
func FetchBanchAll(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	_, company, err := CheckUserAndCompany(props)
	if err {return}

	res := (*Mysql).SelectCompanyBanch(1, company.CompanyId)
	(*props).JSON(http.StatusOK, (*res))
	
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
	res, _ := (*Mysql).InsertCompanyBanch(&banch)
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

	res := (*Mysql).UpdateCompanyBanch(1, &banch, company.CompanyCode)
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

	if !(*Mysql).DeleteCompanyBanch(1, convertBanchId, company.CompanyCode) {
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

	// 部門合法
	res := (*Mysql).SelectBanchStyle(3, convertBanchId, company.CompanyCode)
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

	user, company, err := CheckUserAndCompany(props)
	if err {return}

	// 最高權限 更新
	if user.Permession == 100 {
		// 都可以更新
		if !(*Mysql).UpdateBanchStyle(1, &banchStyle, company.CompanyCode) {
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
		if !(*Mysql).UpdateBanchStyle(1, &banchStyle, company.CompanyCode) {
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

	if methods.IsNotExited(
		(*Mysql).SelectCompanyBanch(3, banchStyle.BanchId, company.CompanyId),
	) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().NotHaveBanch,
		})
		return
	}
	

	// 最高權限 更新
	if user.Permession == 100 {
		if  v, _ := (*Mysql).InsertBanchStyle(&banchStyle); !v {
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
		if v, _ := (*Mysql).InsertBanchStyle(&banchStyle); !v {
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
	banchStyle := (*Mysql).SelectBanchStyle(1, StyleId)
	if methods.IsNotExited(banchStyle) {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().DeleteFail + "," + StatusText().NotHaveBanch,
		})
		return
	}

	user, company, err := CheckUserAndCompany(props)
	if err {return}

	//最高權限刪除
	if user.Permession == 100 {
		if !(*Mysql).DeleteBanchStyle(1, StyleId, company.CompanyCode) {
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
		if !(*Mysql).DeleteBanchStyle(1, StyleId, company.CompanyCode) {
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

// company info
func FetchCompany (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	company := &[]table.CompanyTable{}

	companyCode := (*props).Query("companyCode")
	companyId, err := methods.AnyToInt64((*props).Query("companyId"))

	if companyCode != "" {
		company = (*Mysql).SelectCompany(2, companyCode)
	} else if err == nil {
		company = (*Mysql).SelectCompany(1, companyId)
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

	request.CompanyId = company.CompanyId
	request.BossId = user.UserId
	request.CompanyCode = user.CompanyCode

	res := (*Mysql).UpdateCompany(0, &request)
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

	me, _, err := CheckUserAndCompany(props)
	if err {return}

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
	company.BossId = me.UserId
	company.CreateTime = now
	company.LastModify = now
	
	if v, _ := (*Mysql).InsertCompany(&company); !v {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().InsertFail,
		})
		return
	}

	// 更改負責人的資料
	me.Banch = -1
	me.Permession = 100
	me.CompanyCode = company.CompanyCode
	me.LastModify = now


	if !(*Mysql).UpdateUser(2, &me) {
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

	user, company, err := CheckUserAndCompany(props)
	if err {return}


	if user.Permession == 100 {
		waitCompanyReply := (*Mysql).SelectWaitCompanyReply(5, company.CompanyId)
		props.JSON(http.StatusOK, gin.H{
			"data": *waitCompanyReply,
			"message": StatusText().FindSuccess,
		})
	} else {
		// 只能拿自己的
		waitCompanyReply := (*Mysql).SelectWaitCompanyReply(2, user.UserId)
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

	waitCompanyReply := (*Mysql).SelectWaitCompanyReply(1, request.WaitId)
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
	if !(*Mysql).UpdateWaitCompanyReply(0, &request) {
		(*props).JSON(http.StatusForbidden, gin.H{
			"message": StatusText().UpdateFail,
		})
		return
	}

	// 進入 公司
	if request.IsAccept == 2 {
		targetUser := (*Mysql).SelectUser(1, (*waitCompanyReply)[0].UserId)
		if methods.IsNotExited(targetUser) {
			(*props).JSON(http.StatusBadRequest, gin.H{
				"message": StatusText().userDataNotFound,
			})
			return
		}
		(*targetUser)[0].CompanyCode = company.CompanyCode
		if !(*Mysql).UpdateUser(0, &(*targetUser)[0]) {
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

	findCompany := (*Mysql).SelectCompany(2, request.CompanyCode)
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
	res, _ := (*Mysql).InsertWaitCompanyReply(&waitCompanyReply)
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
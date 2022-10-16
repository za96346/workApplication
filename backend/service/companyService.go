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
	request := new(struct {
		StyleId int64
		Icon string  // 時段圖標
		TimeRangeName string // 時段名稱
		OnShiftTime string // 開始上班時間
		OffShiftTime string  //結束上班的時間
	})

	if (*props).ShouldBindJSON(request) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	// 檢查是否有此style id
	res := (*dbHandle).SelectBanchStyle(1, (*request).StyleId)
	if methods.IsNotExited(res) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().StyleIdNotRight,
		})
		return
	}

	banchStyle := table.BanchStyle{
		Icon: (*request).Icon,
		TimeRangeName: (*request).TimeRangeName,
		OnShiftTime: (*request).OnShiftTime,
		OffShiftTime: (*request).OffShiftTime,
		BanchId: (*res)[0].BanchId,
		LastModify: time.Now(),
		StyleId: (*request).StyleId,
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
	request := new(struct{
		MinPeople int // 限制最少員工
        MaxPeople int // 限制做多的員工
        WeekDay int // 星期幾 (1, 2, 3, 4, 5, 6, 7)
        WeekType int // 寒暑假 或 平常(1, 2, 3)
        OnShiftTime string // 開始上班時間
        OffShiftTime string  //結束上班的時間
        RuleId int64
	})
	if (*props).ShouldBindJSON(request) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	res := (*dbHandle).SelectBanchRule(1, request.RuleId)
	if methods.IsNotExited(res) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().RuleIdIsNotRight,
		})
		return
	}

	banchRule := table.BanchRule{
		BanchId: (*res)[0].BanchId,
		RuleId: (*request).RuleId,
		MinPeople: (*request).MinPeople,
		MaxPeople: (*request).MaxPeople,
		WeekDay: (*request).WeekDay,
		WeekType: (*request).WeekType,
		OnShiftTime: (*request).OnShiftTime,
		OffShiftTime: (*request).OffShiftTime,
		LastModify: time.Now(),
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
package service

import (
	"backend/methods"
	"backend/mysql"
	"backend/panicHandler"
	"backend/redis"
	"backend/table"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var Mysql = mysql.Singleton()
var Redis = redis.Singleton()
var panicHandle = panichandler.Recover

type statusText struct {
	LoginFail string
	NoUser string
	AccountOrPasswordError string
	RegisterFailNotAcceptDataFormat string
	AccountHasBeenRegisted string
	RegisterFail string
	RegisterSuccess string
	LoginSuccess string
	FormatError string
	UpdateSuccess string
	UpdateFail string
	UserIdNotFound string
	userDataNotFound string
	FindSuccess string
	EmailCaptchaSendSuccess string
	EmailCaptchaSendFail string
	EmailCaptchaIsNotRight string
	PasswordIsNotSame string
	EmailIsNotRight string
	GetCompanyCodeFailed string
	IsNotHaveCompany string
	BanchIdIsNotRight string
	StyleIdNotRight string
	NotHaveBanch string
	AssertionFail string
	OnlyCanUpDateYourBanch string
	RuleIdIsNotRight string
	InsertSuccess string
	InsertFail string
	CompanyCodeIsNotRight string
	CompanyNotEqual string
	PasswordNotSafe string
	CanNotUpdateBoss string
	DeleteFail string
	DeleteSuccess string
	OnlyCanDeleteYourBanch string
	QuitWorkUserInsertFail string
	YouAreNotBoss string
	CompanyCodeIsNotTenLength string
	NoData string
	WeekendIdNotRight string
	CompanyIdNotRight string
}
var statusTextInstance *statusText
var statusTextMux = new(sync.Mutex)

func StatusText() *statusText {
	if statusTextInstance == nil {
		statusTextMux.Lock()
		defer statusTextMux.Unlock()
		if statusTextInstance == nil {
			statusTextInstance = &statusText{
				LoginFail:  "登入失敗 請輸入有效的資料",
				NoUser: "沒有此使用者",
				AccountOrPasswordError:  "帳號或密碼錯誤",
				RegisterFailNotAcceptDataFormat: "註冊失敗， 資料不正確",
				AccountHasBeenRegisted: "此帳號已經被註冊了",
				RegisterFail: "註冊失敗",
				RegisterSuccess:  "註冊成功",
				LoginSuccess:  "登入成功",
				FormatError: "格式錯誤",
				UpdateSuccess: "更新成功",
				UpdateFail: "更新失敗",
				UserIdNotFound: "找不到使用者id",
				userDataNotFound: "找不到使用者資料",
				FindSuccess: "資料獲取成功",
				EmailCaptchaSendSuccess: "電子郵件驗證碼發送成功",
				EmailCaptchaSendFail: "電子郵件驗證碼發送失敗",
				EmailCaptchaIsNotRight: "驗證碼錯誤",
				PasswordIsNotSame: "密碼不相等",
				EmailIsNotRight: "電子信箱格式錯誤",
				GetCompanyCodeFailed: "獲取公司碼失敗",
				IsNotHaveCompany: "尚未有公司",
				BanchIdIsNotRight: "部門id不正確",
				StyleIdNotRight: "style id不正確",
				RuleIdIsNotRight: "rule id不正確",
				NotHaveBanch: "尚未有此部門",
				AssertionFail: "Assert Fail",
				OnlyCanUpDateYourBanch: "只可更新你的部門資料",
				InsertSuccess: "新增成功",
				InsertFail: "新增失敗",
				CompanyCodeIsNotRight: "公司碼不正確",
				CompanyNotEqual: "公司不同",
				PasswordNotSafe: "密碼小於8碼",
				CanNotUpdateBoss: "無法更新公司負責人的資料",
				DeleteFail: "刪除失敗",
				DeleteSuccess: "刪除成功",
				OnlyCanDeleteYourBanch: "只可刪除你的部門資料",
				QuitWorkUserInsertFail: "離職員工新增失敗",
				YouAreNotBoss: "非公司負責人，無法更新",
				CompanyCodeIsNotTenLength: "公司碼小於十碼",
				NoData: "沒有資料",
				WeekendIdNotRight: "weekendId 不正確",
				CompanyIdNotRight: "公司id 不正確",
			}
		}
	}
	return statusTextInstance
}

func BanchIsInCompany(banchId int64, companyId int64) bool {
	res := (*Mysql).SelectCompanyBanch(1, companyId)
	for _, v := range *res {
		if v.Id == banchId {
			return true
		}
	}
	return false
}

func CheckUserAndCompany(props *gin.Context) (table.UserTable, table.CompanyTable, bool) {
	// 檢查是否是 company table type
	myCompany, exited := (*props).Get("MyCompany")
	if !exited {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().IsNotHaveCompany,
		})
		return table.UserTable{}, table.CompanyTable{}, true
	}
	company, a := methods.Assertion[table.CompanyTable](myCompany)
	if !a {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().AssertionFail,
		})
		return table.UserTable{}, table.CompanyTable{}, true
	}
		
	// 檢查是否是 user table type
	myUserData, exited := (*props).Get("MyUserData")
	if !exited {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return table.UserTable{}, table.CompanyTable{}, true
	}
	user, a := methods.Assertion[table.UserTable](myUserData)
	if !a {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().AssertionFail,
		})
		return table.UserTable{}, table.CompanyTable{}, true
	}
	return user, company, false
}

func checkMineUserId (props *gin.Context) (int64, bool) {
	userId, existed := (*props).Get("UserId")

	// user id 尋找
	if !existed {
		(*props).JSON(http.StatusInternalServerError, gin.H{
			"message": StatusText().UserIdNotFound,
		})
		return -1, false
	}

	converUserId, err := methods.AnyToInt64(userId)
	if err != nil {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": "轉換格式失敗",
		})
		return -1, false
	}
	return converUserId, true
}
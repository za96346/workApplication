package service

import (
	"backend/mysql"
	"backend/panicHandler"
	"backend/redis"
	"backend/table"
	"sync"

	"github.com/gin-gonic/gin"
	"backend/logger"
)

var Mysql = mysql.Singleton()
var Redis = redis.Singleton()
var panicHandle = panichandler.Recover
var Log = logger.Logger()
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
	CopySuccess string
	CopyFail string
	AccountNotSafe string
	NoMonthSelect string
	NoYearSelect string
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
				CopySuccess: "複製成功",
				CopyFail: "複製失敗，可能有同筆資料",
				AccountNotSafe: "帳號小於五碼",
				NoMonthSelect: "請選擇正確月份",
				NoYearSelect: "請選擇正確年份",
			}
		}
	}
	return statusTextInstance
}

func CheckUserAndCompany(props *gin.Context) (table.UserTable, table.CompanyTable, bool) {

	Account, _ := (*props).Get("Account")
	UserId, _ := (*props).Get("UserId")
	CompanyCode, _ := (*props).Get("CompanyCode")
	UserName, _ := (*props).Get("UserName")
	// EmployeeNumber, _ := (*props).Get("EmployeeNumber")
	// OnWorkDay, _ := (*props).Get("OnWorkDay")
	Banch, _ := (*props).Get("Banch")
	Permession, _ := (*props).Get("Permession")
	CompanyId, _ := (*props).Get("CompanyId")
	BossId, _ := (*props).Get("BossId")


	convUserId, ok := UserId.(int64)
	if !ok {
		convUserId = int64(-10)
	}
	convAccount, ok := Account.(string)
	if !ok {
		convAccount = ""
	}
	convCompanyCode, ok := CompanyCode.(string)
	if !ok {
		convCompanyCode = ""
	}
	convUserName, ok := UserName.(string)
	if !ok {
		convUserName = ""
	}
	convBanch, ok := Banch.(int64)
	if !ok {
		convBanch = int64(-20)
	}
	convPremession, ok := Permession.(int)
	if !ok {
		convPremession = 2
	}
	convCompanyId, ok := CompanyId.(int64)
	if !ok {
		convCompanyId = int64(-20)
	}
	convBossId, ok := BossId.(int64)
	if !ok {
		convBossId = int64(-100)
	}


	user := table.UserTable{
		UserId: convUserId,
		Account: convAccount,
		CompanyCode: convCompanyCode,
		UserName: convUserName,
		Banch: convBanch,
		Permession: convPremession,
	}
	company := table.CompanyTable {
		CompanyId: convCompanyId,
		BossId: convBossId,
		CompanyCode: convCompanyCode,
	}
	
	return user, company, false
}

func UserExtendToUserTable (props *table.UserExtend) *table.UserTable {
	user := table.UserTable{
		UserId: (*props).UserId,
		CompanyCode: (*props).CompanyCode,
		Account: (*props).Account,
		Password: (*props).Password,
		UserName: (*props).UserName,
		EmployeeNumber: (*props).EmployeeNumber,
		OnWorkDay: (*props).OnWorkDay,
		Banch: (*props).Banch,
		Permession: (*props).Permession,
		MonthSalary: (*props).MonthSalary,
		PartTimeSalary: (*props).PartTimeSalary,
		CreateTime: (*props).CreateTime,
		LastModify: (*props).LastModify,
	}

	return &user
}

// func UserExtendToUserTable (props *table.UserTable) *table.UserExtend {
// 	user := table.UserExtend {

// 	}
// }
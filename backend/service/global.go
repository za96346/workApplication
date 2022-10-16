package service

import (
	"backend/handler"
	"backend/panicHandler"
	"sync"
)

var dbHandle = handler.Singleton()
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
			}
		}
	}
	return statusTextInstance
}

func BanchIsInCompany(banchId int64, companyId int64) bool {
	res := (*dbHandle).SelectCompanyBanch(1, companyId)
	for _, v := range *res {
		if v.Id == banchId {
			return true
		}
	}
	return false
}
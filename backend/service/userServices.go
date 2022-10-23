package service

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	// "strconv"
	"backend/methods"
	"backend/response"
	"backend/table"

	"github.com/gin-gonic/gin"
)
func FindSingleUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	userId := (*props).Query("userId")
	fmt.Println("userId => ", userId)
	// 尋找 userData
	intUserId, err := methods.AnyToInt64(userId)
	if err != nil {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": "轉換格式失敗",
		})
		return
	}
	res := (*dbHandle).SelectUser(1, intUserId)
	if methods.IsNotExited(res) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}
	data := []response.User{}

	data = append(data, response.User{
		UserId: (*res)[0].UserId,
		CompanyCode: (*res)[0].CompanyCode,
		OnWorkDay: (*res)[0].OnWorkDay,
		EmployeeNumber: (*res)[0].EmployeeNumber,
		UserName: (*res)[0].UserName,
		Banch: (*res)[0].Banch,
		Permession: (*res)[0].Permession,
		WorkState: "on",// 者個要去離職表找
	})
	// 尋找資料
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().FindSuccess,
		"data": data,
	})
	
}

func FindMine(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	userId, status := checkMineUserId(props)
	if !status {return}

	// 尋找資料
	res := (*dbHandle).SelectUser(1, userId)
	if methods.IsNotExited(res) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
			"data": *res,
		})
		return
	}

	(*res)[0].Password = ""

	// 找到資料
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().FindSuccess,
		"data": *res,
	})

}

func UpdateMine (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()

	mineId, status := checkMineUserId(props)
	if !status {return}

	user := table.UserTable{}
	if (*props).ShouldBindJSON(&user) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	mineUserData := (*dbHandle).SelectUser(1, mineId)
	if methods.IsNotExited(mineUserData) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}
	
	(*mineUserData)[0].UserName = user.UserName
	(*mineUserData)[0].LastModify = time.Now()

	if !(*dbHandle).UpdateUser(0, &(*mineUserData)[0]) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().UpdateFail,
		})
		return
	}

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().UpdateSuccess,
	})
}

func GetAllUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	companyCode, existed := (*props).Get("CompanyCode")
	if !existed {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().GetCompanyCodeFailed,
		})
		return
	}

	user, _, err := CheckUserAndCompany(props)
	if err {return}

	// 拿在職的員工資料
	userList := (*dbHandle).SelectUser(3, companyCode.(string))
	data := []response.User{}
	for _, v := range *userList {
		if user.Permession == 1 && v.Banch != user.Banch{
			continue
		}
		data = append(data, response.User{
			UserId: v.UserId,
			CompanyCode: v.CompanyCode,
			OnWorkDay: v.OnWorkDay,
			UserName: v.UserName,
			EmployeeNumber: v.EmployeeNumber,
			Banch: v.Banch,
			Permession: v.Permession,
			WorkState: "on",
		})
	}


	// 拿離職的員工資料
	quitWorkUser := (*dbHandle).SelectQuitWorkUser(3, companyCode.(string))
	for _, v := range *quitWorkUser {
		if user.Permession == 1 && v.Banch != user.Banch{
			continue
		}
		data = append(data, response.User{
			UserId: v.UserId,
			CompanyCode: v.CompanyCode,
			OnWorkDay: v.OnWorkDay,
			UserName: v.UserName,
			EmployeeNumber: v.EmployeeNumber,
			Banch: v.Banch,
			Permession: v.Permession,
			WorkState: "off", // 這個要去離職表找
		})
	}


	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().FindSuccess,
		"data": data,
	})
}

func UpdateUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	now := time.Now()

	// 檢查格式
	request := response.User{}
	if (*props).ShouldBindJSON(&request) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	// 抓曲目標user
	targetUser, existed := (*props).Get("targetUser")
	if !existed {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}

	// 斷言
	convertTargetUser, a := methods.Assertion[table.UserTable](targetUser)
	if !a {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().AssertionFail,
		})
		return
	}

	permession := request.Permession
	companyCode := request.CompanyCode
	banch := request.Banch

	// 當要被更改的人是公司負責人時
	//就要去判斷是不是自己更改自己
	//如果不是 則 回傳不能修改
	//如果是 永遠把負責人的權限設為管理者
	// 以及把公司碼射為自己的公司
	myUserData, myCompany, err := CheckUserAndCompany(props)
	if err {return}

	if convertTargetUser.UserId == myCompany.BossId {
		permession = 100
		companyCode = myCompany.CompanyCode
		banch = -1
		if myUserData.UserId != convertTargetUser.UserId {
			(*props).JSON(http.StatusForbidden, gin.H{
				"message": StatusText().CanNotUpdateBoss,
			})
			return
		}
	} else if request.WorkState == "off" {
		// 要去判斷 工作狀態並把它丟到quit work user table
		quitWorkUser := table.QuitWorkUser{
			CompanyCode: companyCode,
			EmployeeNumber: request.EmployeeNumber,
			Account: convertTargetUser.Account,
			UserName: convertTargetUser.UserName,
			OnWorkDay: request.OnWorkDay,
			Banch: banch,
			Permession: permession,
			CreateTime: convertTargetUser.CreateTime,
			LastModify: now,
			MonthSalary: 0,
			PartTimeSalary: 0,
			UserId: convertTargetUser.UserId,
		}
		if a, _ := (*dbHandle).InsertQuitWorkUser(&quitWorkUser); !a {
			(*props).JSON(http.StatusForbidden, gin.H{
				"message": StatusText().QuitWorkUserInsertFail,
			})
			return
		}
		companyCode = ""
	} else if request.WorkState == "on" {
		res := (*dbHandle).SelectQuitWorkUser(4, companyCode, convertTargetUser.UserId)
		if !methods.IsNotExited(res) {
			(*dbHandle).DeleteQuitWorkUser(0, (*res)[0].QuitId)
		}
	}

	user := table.UserTable{
		CompanyCode: companyCode,
		EmployeeNumber: request.EmployeeNumber,
		Password: convertTargetUser.Password,
		UserName: convertTargetUser.UserName,
		OnWorkDay: request.OnWorkDay,
		Banch: banch,
		Permession: permession,
		LastModify: now,
		MonthSalary: 0,
		PartTimeSalary: 0,
		UserId: convertTargetUser.UserId,
	}

	res := (*dbHandle).UpdateUser(0, &user)
	if !res {
		(*props).JSON(http.StatusForbidden, gin.H{
			"message": StatusText().UpdateFail,
		})
		return
	}
	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().UpdateSuccess,
	})
}

func ChangePassword (props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer waitJob.Done()

	// 確認格式
	changePwd := response.ChangePassword{}
	if (*props).ShouldBindJSON(&changePwd) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	userId, err := checkMineUserId(props)
	if !err {return}

	// get user
	res := (*dbHandle).SelectUser(1, userId)
	if methods.IsNotExited(res) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}

	// 驗證碼驗證
	rightCaptcha := (*dbHandle).Redis.SelectEmailCaptcha((*res)[0].Account)
	if rightCaptcha != changePwd.Captcha || rightCaptcha == -1 {
		(*props).JSON(http.StatusBadRequest, gin.H{
			"message": StatusText().EmailCaptchaIsNotRight,
		})
		return
	}

	// 驗證就密碼
	if changePwd.OldPassword != (*res)[0].Password {
		(*props).JSON(http.StatusBadRequest, gin.H{
			"message": "舊密碼不正確",
		})
		return
	}

	// 密碼驗證
	if changePwd.NewPassword != changePwd.NewPasswordConfirm {
		(*props).JSON(http.StatusUnprocessableEntity, gin.H{
			"message": StatusText().PasswordIsNotSame,
		})
		return
	}

	// 密碼長度驗證
	if len(changePwd.NewPassword) < 8 {
		(*props).JSON(http.StatusUnprocessableEntity, gin.H{
			"message": StatusText().PasswordNotSafe,
		})
		return
	}

	(*res)[0].Password = changePwd.NewPassword
	status := (*dbHandle).UpdateUser(0, &(*res)[0])
	if !status {
		(*props).JSON(http.StatusUnprocessableEntity, gin.H{
			"message": StatusText().UpdateFail,
		})
		return
	}
	(*dbHandle).Redis.DeleteCaptcha((*res)[0].Account)

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().UpdateSuccess,
	})
	return
	
}

func ForgetPassword(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer waitJob.Done()

	// 確認格式
	forgetPwd := response.ForgetPassword{}
	if (*props).ShouldBindJSON(&forgetPwd) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	// get user
	res := (*dbHandle).SelectUser(2, forgetPwd.Email)
	if methods.IsNotExited(res) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}

	// 驗證碼驗證
	rightCaptcha := (*dbHandle).Redis.SelectEmailCaptcha((*res)[0].Account)
	if rightCaptcha != forgetPwd.Captcha || rightCaptcha == -1 {
		(*props).JSON(http.StatusBadRequest, gin.H{
			"message": StatusText().EmailCaptchaIsNotRight,
		})
		return
	}

	// 密碼驗證
	if forgetPwd.NewPassword != forgetPwd.NewPasswordConfirm {
		(*props).JSON(http.StatusUnprocessableEntity, gin.H{
			"message": StatusText().PasswordIsNotSame,
		})
		return
	}

	// 密碼長度驗證
	if len(forgetPwd.NewPassword) < 8 {
		(*props).JSON(http.StatusUnprocessableEntity, gin.H{
			"message": StatusText().PasswordNotSafe,
		})
		return
	}

	(*res)[0].Password = forgetPwd.NewPassword
	status := (*dbHandle).UpdateUser(0, &(*res)[0])
	if !status {
		(*props).JSON(http.StatusUnprocessableEntity, gin.H{
			"message": StatusText().UpdateFail,
		})
		return
	}
	(*dbHandle).Redis.DeleteCaptcha((*res)[0].Account)

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().UpdateSuccess,
	})
	return
}
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
	res := (*Mysql).SelectUser(1, intUserId)
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

	user, _, err := CheckUserAndCompany(props)
	if err {return}

	// 尋找資料
	res := (*Mysql).SelectUser(1, user.UserId)
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

	me, _, err := CheckUserAndCompany(props)
	if err {return}

	user := table.UserTable{}
	if (*props).ShouldBindJSON(&user) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}

	mineUserData := (*Mysql).SelectUser(1, me.UserId)
	if methods.IsNotExited(mineUserData) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}
	
	(*mineUserData)[0].UserName = user.UserName
	(*mineUserData)[0].LastModify = time.Now()

	if !(*Mysql).UpdateUser(0, UserExtendToUserTable(&(*mineUserData)[0])) {
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
	workState := (*props).Query("workState")
	name := (*props).Query("name")
	banch, isBanchError := methods.AnyToInt64((*props).Query("banch"))

	user, company, err := CheckUserAndCompany(props)
	if err {return}

	data := []response.User{}
	// 管理員 沒帶部門查詢
	if user.Permession == 100 && isBanchError != nil {
		data = *((*Mysql).SelectAllUser(
			0,
			company.CompanyCode,
			company.CompanyCode,
			name,
			name,
			name,
		))
	// 主管查詢 或是 管理員 有帶部門查詢
	} else if user.Permession == 1 ||
		(user.Permession == 100 && isBanchError == nil) {
		b := user.Banch
		if user.Permession == 100 {
			b = banch
		}
		data = *((*Mysql).SelectAllUser(
			1,
			company.CompanyCode,
			company.CompanyCode,
			b,
			b,
			name,
			name,
			name,
		))
	}

	// 判斷工作狀態
	val := []response.User{}
	for _, v := range data {
		if (workState == "all") {
			val = append(val, v)
		} else if workState == v.WorkState {
			val = append(val, v)
		}
	}


	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().FindSuccess,
		"data": val,
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



	// 當要被更改的人是公司負責人時
	// 就要去判斷是不是自己更改自己
	// 如果不是 則 回傳不能修改
	// 如果是 永遠把負責人的權限設為管理者
	// 以及把公司碼射為自己的公司
	myUserData, myCompany, err := CheckUserAndCompany(props)
	if err {return}

	permession := request.Permession
	companyCode := myCompany.CompanyCode
	banch := request.Banch

	if request.UserId == myCompany.BossId {
		permession = 100
		companyCode = myCompany.CompanyCode
		banch = -1
		if myUserData.UserId != request.UserId {
			(*props).JSON(http.StatusForbidden, gin.H{
				"message": StatusText().CanNotUpdateBoss,
			})
			return
		}
	} else if request.WorkState == "off" {
		// 要去判斷 工作狀態並把它丟到quit work user table
		(*Mysql).InsertQuitWorkUserBySelectUser(request.UserId, myCompany.CompanyCode)
		companyCode = ""
	} else if request.WorkState == "on" {
		(*Mysql).DeleteQuitWorkUser(1, request.UserId, myCompany.CompanyCode)
		companyCode = myCompany.CompanyCode
	}
	Log.Println("companyCode => ", companyCode)

	user := table.UserTable{
		CompanyCode: companyCode,
		EmployeeNumber: request.EmployeeNumber,
		OnWorkDay: request.OnWorkDay,
		Banch: banch,
		Permession: permession,
		LastModify: now,
		MonthSalary: 0,
		PartTimeSalary: 0,
		UserId: request.UserId,
	}

	res := (*Mysql).UpdateUser(1, &user, myCompany.CompanyCode)
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

	user, _, err := CheckUserAndCompany(props)
	if err {return}

	// get user
	res := (*Mysql).SelectUser(1, user.UserId)
	if methods.IsNotExited(res) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}

	// 驗證碼驗證
	rightCaptcha := (*Redis).SelectEmailCaptcha((*res)[0].Account)
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
	status := (*Mysql).UpdateUser(0, UserExtendToUserTable(&(*res)[0]))
	if !status {
		(*props).JSON(http.StatusUnprocessableEntity, gin.H{
			"message": StatusText().UpdateFail,
		})
		return
	}
	(*Redis).DeleteCaptcha((*res)[0].Account)

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
	res := (*Mysql).SelectUser(2, forgetPwd.Email)
	if methods.IsNotExited(res) {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": StatusText().userDataNotFound,
		})
		return
	}

	// 驗證碼驗證
	rightCaptcha := (*Redis).SelectEmailCaptcha((*res)[0].Account)
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
	status := (*Mysql).UpdateUser(0, UserExtendToUserTable(&(*res)[0]))
	if !status {
		(*props).JSON(http.StatusUnprocessableEntity, gin.H{
			"message": StatusText().UpdateFail,
		})
		return
	}
	(*Redis).DeleteCaptcha((*res)[0].Account)

	(*props).JSON(http.StatusOK, gin.H{
		"message": StatusText().UpdateSuccess,
	})
	return
}
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
	userId, existed := (*props).Get("UserId")

	// user id 尋找
	if !existed {
		(*props).JSON(http.StatusInternalServerError, gin.H{
			"message": StatusText().UserIdNotFound,
		})
		return
	}

	converUserId, err := methods.AnyToInt64(userId)
	if err != nil {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": "轉換格式失敗",
		})
		return
	}

	// 尋找資料
	res := (*dbHandle).SelectUser(1, converUserId)
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
	userList := (*dbHandle).SelectUser(3, companyCode.(string))
	data := []response.User{}
	for _, v := range *userList {
		data = append(data, response.User{
			UserId: v.UserId,
			CompanyCode: v.CompanyCode,
			OnWorkDay: v.OnWorkDay,
			UserName: v.UserName,
			EmployeeNumber: v.EmployeeNumber,
			Banch: v.Banch,
			Permession: v.Permession,
			WorkState: "on", // 這個要去離職表找
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
	}

	
	// 要去判斷 工作狀態並把它丟到quit work user table


	

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
	(*props).JSON(http.StatusNotAcceptable, gin.H{
		"message": StatusText().UpdateSuccess,
	})
}

func DeleteUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer waitJob.Done()
	// deleteUser := []pojo.User{}
}

func ForgetPassword (props *gin.Context, waitJob *sync.WaitGroup) {
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

	// 驗證就密碼
	if forgetPwd.OldPassword != (*res)[0].Password {
		(*props).JSON(http.StatusBadRequest, gin.H{
			"message": "舊密碼不正確",
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
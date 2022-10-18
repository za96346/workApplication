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
		WorkState: (*res)[0].WorkState,
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
			WorkState: v.WorkState,
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

	user := table.UserTable{
		CompanyCode: request.CompanyCode,
		EmployeeNumber: request.EmployeeNumber,
		Password: convertTargetUser.Password,
		UserName: convertTargetUser.UserName,
		OnWorkDay: request.OnWorkDay,
		Banch: request.Banch,
		Permession: request.Permession,
		WorkState: request.WorkState,
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
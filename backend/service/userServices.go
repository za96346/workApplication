package service

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	// "strconv"
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
	intUserId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		(*props).JSON(http.StatusNotFound, gin.H{
			"message": "轉換格式失敗",
		})
		return
	}
	res := (*dbHandle).SelectUser(1, intUserId)
	if len(*res) == 0 {
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

	// 尋找資料
	res := (*dbHandle).SelectUser(1, userId.(int64))
	if len(*res) == 0 {
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
			Banch: v.Banch,
			Permession: v.Permession,
			WorkState: v.WorkState,
		})
	}
	(*props).JSON(http.StatusNotAcceptable, gin.H{
		"message": StatusText().FindSuccess,
		"data": data,
	})
}

func UpdateUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	user := table.UserTable{}
	// 格式錯誤
	if (*props).ShouldBindJSON(&user) != nil {
		(*props).JSON(http.StatusNotAcceptable, gin.H{
			"message": StatusText().FormatError,
		})
		return
	}
	user.LastModify = time.Now()
	res := (*dbHandle).UpdateUser(0, &user)
	fmt.Println(user)

	// 更新資料成功
	if res {
		(*props).JSON(http.StatusOK, gin.H{
			"message": StatusText().UpdateSuccess,
		})
		return
	// 更新資料失敗
	} else {
		(*props).JSON(http.StatusBadRequest, gin.H{
			"message": StatusText().UpdateFail,
		})
		return
	}
}

func DeleteUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer waitJob.Done()
	// deleteUser := []pojo.User{}
}
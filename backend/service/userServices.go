package service

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	// "strconv"
	"backend/table"

	"github.com/gin-gonic/gin"
)
func FindSingleUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	fmt.Println("前端傳來資料 =>", props.Params[0])
	//get
	(*props).JSON(http.StatusOK, "hi")
}

func FindMine(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()


}

func UpdateUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	user := table.UserTable{}
	(*props).ShouldBindJSON(&user)
	user.LastModify = time.Now()
	res := (*dbHandle).UpdateUser(0, &user)
	fmt.Println(user)
	if res {
		(*props).JSON(http.StatusOK, "更新成功")
	} else {
		(*props).JSON(http.StatusNotFound, "更新失敗")
	}
	// account, _ := (*props).Get("Account")
	// userId, _ := (*props).Get("UserId")
	// fmt.Println("im get data =>", account, userId)
}

func DeleteUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer waitJob.Done()
	// deleteUser := []pojo.User{}
}
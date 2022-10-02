package service

import (
	"fmt"
	"net/http"
	"sync"

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

func CreateUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	// post
	defer (*waitJob).Done()
	user := table.UserTable{}
	(*props).ShouldBindJSON(&user)

	fmt.Println("do create user", &user)
	(*dbHandle).InsertUser(&user)
	(*props).JSON(http.StatusOK, user)
}

func UpdateUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	user := table.UserTable{}
	(*props).ShouldBindJSON(&user)
	res := (*dbHandle).UpdateUser(0, &user)
	fmt.Println(user)
	if res {
		(*props).JSON(http.StatusOK, "更新成功")
	} else {
		(*props).JSON(http.StatusNotFound, "更新失敗")
	}
}

func DeleteUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer waitJob.Done()
	// deleteUser := []pojo.User{}
}
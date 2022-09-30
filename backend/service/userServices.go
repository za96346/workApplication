package service

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"sync"

	// "strconv"
	"backend/pojo"
	"backend/table"
	// "backend/database"
	"github.com/gin-gonic/gin"
	"backend/handler"
)
var userList = []pojo.User{}
func FindSingleUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer (*waitJob).Done()
	fmt.Println("前端傳來資料 =>", props.Params[0])
	//get
	(*props).JSON(http.StatusOK, "hi")
}

func CreateUser(props *gin.Context, waitJob *sync.WaitGroup) {
	// post
	defer (*waitJob).Done()
	user := table.UserTable{}
	(*props).ShouldBindJSON(&user)

	fmt.Println("do create user", &user)
	(*handler.Singleton()).InsertUser(&user)
	(*props).JSON(http.StatusOK, user)
}

func UpdateUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer (*waitJob).Done()
	user := table.UserTable{}
	(*props).ShouldBindJSON(&user)
	res := (*handler.Singleton()).UpdateUser(0, &user)
	if res {
		(*props).JSON(http.StatusOK, "it ok")
	} else {
		(*props).JSON(http.StatusNotFound, "新增失敗")
	}
}

func DeleteUser(props *gin.Context, waitJob *sync.WaitGroup) {
	defer waitJob.Done()
	// deleteUser := []pojo.User{}
}
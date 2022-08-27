package service

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	// "strconv"

	"backend/pojo"

	"github.com/gin-gonic/gin"
)
var userList = []pojo.User{}
func FindSingleUser(props *gin.Context) {
	fmt.Println("前端傳來資料 =>", props.Params[0])
	//get
	props.JSON(http.StatusOK, userList)
}

func CreateUser(props *gin.Context) {
	fmt.Println("前端傳來資料 =>", props.Params[0])
	//post
	props.JSON(http.StatusOK, userList)
}

func UpdateUser(props *gin.Context) {
	//put
	// update

	user := []pojo.User{}
	err := props.ShouldBindJSON(&user)
	if err != nil {
		props.JSON(http.StatusUnprocessableEntity, "資料格式錯誤")
		return;
	}
	for amount := 0; amount < len(user); amount++ {
		userList = append(userList, user[amount])
	}
	log.Println("update user receive => ", userList, userList[len(userList) - 1])
	props.JSON(http.StatusOK, "資料修改成功")
}

func DeleteUser(props *gin.Context) {
	// deleteUser := []pojo.User{}
}
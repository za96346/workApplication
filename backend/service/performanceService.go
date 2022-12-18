package service

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func FetchPerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	props.JSON(http.StatusOK, gin.H{
		"message": "not bad",
	})
}
func UpdatePerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	props.JSON(http.StatusOK, gin.H{
		"message": "not bad",
	})
}
func InsertPerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	props.JSON(http.StatusOK, gin.H{
		"message": "not bad",
	})
}
func DeletePerformance(props *gin.Context, waitJob *sync.WaitGroup) {
	defer panicHandle()
	defer (*waitJob).Done()
	props.JSON(http.StatusOK, gin.H{
		"message": "not bad",
	})
}
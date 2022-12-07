package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)
func(dbObj *DB) InsertLog(data *table.Log) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertLog(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectLateExcused(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertLateExcused(&value)
	// 		}
	// 	}()
	// }
	return isOk, id
}

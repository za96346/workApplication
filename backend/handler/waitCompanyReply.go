package handler

import (
	"backend/table"

	// "fmt"
	// "time"
	"backend/panicHandler"
)

func(dbObj *DB) restoreWaitCompanyReplyAll() {
	defer panichandler.Recover()
	(*(*dbObj).waitCompanyReplyLock) = true
	(*dbObj).Redis.DeleteKeyWaitCompanyReply()
	arr := (*dbObj.Mysql).SelectWaitCompanyReply(0)
	forEach(arr, (*dbObj.Redis).InsertWaitCompanyReply)
	(*(*dbObj).waitCompanyReplyLock) = false
}


func(dbObj *DB) InsertWaitCompanyReply(data *table.WaitCompanyReply) (bool, int64) {
	defer panichandler.Recover()
	isOk, id := (*dbObj).Mysql.InsertWaitCompanyReply(data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectWaitCompanyReply(1, id)
	// 		for _, value := range *res {
	// 			(*dbObj).Redis.InsertWaitCompanyReply(&value)
	// 		}
	// 	}()
	// }
	return isOk, id
}

func(dbObj *DB) UpdateWaitCompanyReply(updateKey int, data *table.WaitCompanyReply) bool {
	defer panichandler.Recover()
	isOk := (*dbObj).Mysql.UpdateWaitCompanyReply(updateKey, data)
	// if isOk {
	// 	go func ()  {
	// 		res := (*dbObj).Mysql.SelectWaitCompanyReply(1, int64((*data).WaitId))
	// 		for _, v := range *res {
	// 			(*dbObj).Redis.InsertWaitCompanyReply(&v)
	// 		}
	// 	}()
	// }
	return isOk
}

func(dbObj *DB) DeleteWaitCompanyReply(deleteKey int, waitId int64) bool {
	defer panichandler.Recover()
	res := (*dbObj).Mysql.DeleteWaitCompanyReply(deleteKey, waitId)
	// if res {
	// 	go func ()  {
	// 		(*dbObj).Redis.DeleteWaitCompanyReply(deleteKey, waitId)
	// 	}()	
	// }
	
	return res
}

// 0 => all, value => nil
//  1 => waitId, value => int64
//  2 => userId, value => int64
//  3 => companyId, value => int64
//  4 => comapnyId && userId, value => int64, int64
func(dbObj *DB) SelectWaitCompanyReply(selectKey int, value... interface{}) *[]table.WaitCompanyReply {
	defer panichandler.Recover()
	return selectAllHandler(
		func() *[]table.WaitCompanyReply {
			return (*dbObj.Redis).SelectWaitCompanyReply(selectKey, value...)
		},
		func() *[]table.WaitCompanyReply {
			return (*dbObj.Mysql).SelectWaitCompanyReply(selectKey, value...)
		},
		(*dbObj).waitCompanyReplyLock,
	)
}
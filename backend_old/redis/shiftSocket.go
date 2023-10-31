package redis

import (
	panichandler "backend/panicHandler"
	"backend/response"
	"encoding/json"
	"strconv"
	"time"
)

func(dbObj *DB) EnterShiftRoom (banchId int64, value response.Member) {
	defer panichandler.Recover()
	jsonData, _ := json.Marshal(value)
	conBanchId := strconv.FormatInt(banchId, 10)
	conUserId := strconv.FormatInt(value.UserId, 10)
	(*dbObj).RedisOfShiftSocket.HSet(conBanchId, conUserId, jsonData)
}

func (dbObj *DB) LeaveShiftRoom (banchId int64, userId int64) {
	defer panichandler.Recover()
	conBanchId := strconv.FormatInt(banchId, 10)
	conUserId := strconv.FormatInt(userId, 10)
	(*dbObj).RedisOfShiftSocket.HDel(conBanchId, conUserId)
}

func (dbObj *DB) OnlineShiftUser (banchId int64) {
	defer panichandler.Recover()
	conBanchId := strconv.FormatInt(banchId, 10)
	(*dbObj).RedisOfShiftSocket.HVals(conBanchId)
}

func (dbObj *DB) GetShiftRoomUser (banchId int64) *[]response.Member {
	defer panichandler.Recover()
	return forEach(
		func() ([]string, error) {
			return (*dbObj).RedisOfShiftSocket.HVals(strconv.FormatInt(banchId, 10)).Result()
		},
		func(v response.Member) bool {return true},
	)
}

func (dbObj *DB) InsertShiftData(banchId int64, value response.Shift) {
	defer panichandler.Recover()

	jsonData, _ := json.Marshal(value)
	conBanchId := strconv.FormatInt(banchId, 10)
	positionIdx := value.Date + "__" + strconv.FormatInt(value.UserId,10)
	(*dbObj).RedisOfShiftData.HSet(conBanchId, positionIdx, jsonData)
}

// 拿取 編輯 班表的資料
func (dbObj *DB) GetShiftData (banchId int64, year int, month int) *[]response.Shift {
	defer panichandler.Recover()
	return forEach(
		func() ([]string, error) {
			return (*dbObj).RedisOfShiftData.HVals(strconv.FormatInt(banchId, 10)).Result()
		},
		func(v response.Shift) bool {
			date, _ := time.Parse("2006-01-02", v.Date)
			if date.Year() == year && int(date.Month()) == month {
				return true
			}
			return false
		},
	)
}

// 刪除 單一 部門 所有 班表資料
func (dbObj *DB) DeleteShiftData (banchId int64) {
	defer panichandler.Recover()
	(*dbObj).RedisOfShiftData.Del(strconv.FormatInt(banchId, 10)).Result()
}

// 刪除 單一 部門 的 單一比 班表資料
func (dbObj *DB) DeleteSingleShiftData (banchId int64, value response.Shift) {
	defer panichandler.Recover()
	positionIdx := value.Date + "__" + strconv.FormatInt(value.UserId,10)
	(*dbObj).RedisOfShiftData.HDel(strconv.FormatInt(banchId, 10), positionIdx).Result()
}

// 紀錄 房間的狀態
// 1. 本月資料是否已經送出
func (dbObj *DB) InsertShiftRoomStatus(banchId int64, value any) {
	defer panichandler.Recover()
	jsonData, _ := json.Marshal(value)
	conBanchId := strconv.FormatInt(banchId, 10) + "-roomStatus"
	(*dbObj).RedisOfShiftSocket.HSet(conBanchId, "status", jsonData)
}

func (dbObj *DB) GetShiftRoomStatus(banchId int64) *map[string]any {
	defer panichandler.Recover()
	conBanchId := strconv.FormatInt(banchId, 10) + "-roomStatus"
	v, _ := (*dbObj).RedisOfShiftSocket.HGet(conBanchId, "status").Result()

	rooStatus := new(map[string]any)
	json.Unmarshal([]byte(v), rooStatus)
	return rooStatus
}

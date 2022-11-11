package redis

import (
	panichandler "backend/panicHandler"
	"backend/response"
	"encoding/json"
	"strconv"
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
	positionIdx := value.Date
	(*dbObj).RedisOfShiftData.HSet(conBanchId, positionIdx, jsonData)
}

func (dbObj *DB) GetShiftData (banchId int64) *[]response.Shift {
	defer panichandler.Recover()
	return forEach(
		func() ([]string, error) {
			return (*dbObj).RedisOfShiftData.HVals(strconv.FormatInt(banchId, 10)).Result()
		},
		func(v response.Shift) bool {return true},
	)
}
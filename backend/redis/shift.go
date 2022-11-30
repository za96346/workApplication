package redis

import (
	"encoding/json"
	"strconv"
	panichandler "backend/panicHandler"
	"backend/table"
	"sync"

)

// 0 => all, value => nil
//  1 => 班表id, value => int64
func(dbObj *DB) SelectShift(selectKey int, value... interface{}) *[]table.ShiftTable {
	defer panichandler.Recover()
	tableKey := (*dbObj).table[4]
	switch selectKey {
	case 0:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).RedisDb.HVals(tableKey).Result()
			},
			func(v table.ShiftTable) bool {return true},
		)
	case 1:
		return forEach(
			func() ([]string, error) {
				return (*dbObj).hmGet(tableKey, value...)
			},
			func(v table.ShiftTable) bool {return true},
		)
	default:
		return &[]table.ShiftTable{}
	}
}

func(dbObj *DB) InsertShift(data *table.ShiftTable) bool {
	defer panichandler.Recover()
	(*dbObj).shiftMux.Lock()
	defer (*dbObj).shiftMux.Unlock()
	key := strconv.FormatInt((*data).ShiftId, 10)
	jsonData, _ := json.Marshal(*data)
	_, err := (*dbObj).RedisDb.HSet((*dbObj).table[4], key, jsonData).Result()
	
	(*dbObj).checkErr(err)
	return true
}

// 班表的唯一id
func(dbObj *DB) DeleteShift(deleteKey int, shiftId int64) bool {
	defer panichandler.Recover()
	(*dbObj).shiftMux.Lock()
	defer (*dbObj).shiftMux.Unlock()

	(*dbObj).RedisDb.HDel((*dbObj).table[4], strconv.FormatInt(shiftId, 10))

	wg := new(sync.WaitGroup)
	wg.Add(6)
	// (*dbObj).SelectShiftChange(1)
	go func ()  {
		defer wg.Done()
		res := (*dbObj).SelectShiftOverTime(2, shiftId)
		for _, v := range *res {
			(*dbObj).DeleteShiftOverTime(0, v.CaseId)
		}
	}()
	go func ()  {
		defer wg.Done()
		res := (*dbObj).SelectDayOff(2, shiftId)
		for _, v := range *res {
			(*dbObj).DeleteDayOff(0, v.CaseId)
		}
	}()
	go func ()  {
		defer wg.Done()
		res := (*dbObj).SelectLateExcused(2, shiftId)
		for _, v := range *res {
			(*dbObj).DeleteLateExcused(0, v.CaseId)
		}
	}()
	go func ()  {
		defer wg.Done()
		res := (*dbObj).SelectForgetPunch(2, shiftId)
		for _, v := range *res {
			(*dbObj).DeleteForgetPunch(0, v.CaseId)
		}
	}()
	go func ()  {
		defer wg.Done()
		res1 := (*dbObj).SelectShiftChange(2, shiftId)
		for _, v := range *res1 {
			(*dbObj).DeleteShiftChange(0, v.CaseId)
		}
	}()
	go func ()  {
		defer wg.Done()
		res1 := (*dbObj).SelectShiftChange(3, shiftId)
		for _, v := range *res1 {
			(*dbObj).DeleteShiftChange(0, v.CaseId)
		}
	}()
	wg.Wait()
	return true
}

func(dbObj *DB) DeleteKeyShift(){
	defer panichandler.Recover()
	(*dbObj).RedisDb.Del((*dbObj).table[4])
}
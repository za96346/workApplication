package handler

import (
	"backend/database"
	"backend/redis"
	"fmt"

)
func Init() {
	(*database.MysqlSingleton()).Conn()
	(*redis.RedisSingleton()).Conn()
	(*redis.RedisSingleton()).RedisDb.FlushAll()
	TakeAllFromMysql()
}

func TakeAllFromMysql() {

	userArr := (*database.MysqlSingleton()).SelectUserAll()
	_ = (*redis.RedisSingleton()).InsertUserAll(userArr)

	userPreferenceArr := (*database.MysqlSingleton()).SelectUserPreferenceAll()
	_ = (*redis.RedisSingleton()).InsertUserPreferenceAll(userPreferenceArr)

	companyArr := (*database.MysqlSingleton()).SelectCompanyAll()
	_ = (*redis.RedisSingleton()).InsertCompanyAll(companyArr)

	companyBanchArr := (*database.MysqlSingleton()).SelectCompanyBanchAll()
	_ = (*redis.RedisSingleton()).InsertCompanyBanchAll(companyBanchArr)

	shiftArr := (*database.MysqlSingleton()).SelectShiftAll()
	_ = (*redis.RedisSingleton()).InsertShiftAll(shiftArr)

	shiftChangeArr := (*database.MysqlSingleton()).SelectShiftChangeAll()
	_ = (*redis.RedisSingleton()).InsertShiftChangeAll(shiftChangeArr)

	shiftOverTimeArr := (*database.MysqlSingleton()).SelectShiftOverTimeAll()
	_ = (*redis.RedisSingleton()).InsertShiftOverTimeAll(shiftOverTimeArr)

	dayOffArr := (*database.MysqlSingleton()).SelectDayOffAll()
	_ = (*redis.RedisSingleton()).InsertDayOffAll(dayOffArr)

	lateExcusedArr := (*database.MysqlSingleton()).SelectLateExcusedAll()
	_ = (*redis.RedisSingleton()).InsertLateExcusedAll(lateExcusedArr)

	forgetPunchArr := (*database.MysqlSingleton()).SelectForgetPunchAll()
	_ = (*redis.RedisSingleton()).InsertForgetPunch(forgetPunchArr)

	SelectUserAll()
	SelectCompanyAll()
	SelectUserPreferenceAll()
	SelectCompanyBanchAll()
	SelectShiftAll()
	SelectShiftChangeAll()
	SelectShiftOverTimeAll()
	SelectDayOffAll()
	SelectForgetPunchAll()
	SelectLateExcusedAll()
}
func SelectUserAll() {
	res := (*redis.RedisSingleton()).SelectUserAll()
	if len(*res) == 0 {
		fmt.Println("user 是空的")
	}
	_ = (*database.MysqlSingleton()).SelectUserAll()
	fmt.Println(*res)
}
func SelectUserPreferenceAll() {
	res := (*redis.RedisSingleton()).SelectUserPreferenceAll()
	if len(*res) == 0 {
		fmt.Println("userPreference 是空的")
	}
}
func SelectCompanyAll() {
	res := (*redis.RedisSingleton()).SelectCompanyAll()
	if len(*res) == 0 {
		fmt.Println("company 是空的")
	}
}
func SelectCompanyBanchAll() {
	res := (*redis.RedisSingleton()).SelectCompanyBanchAll()
	if len(*res) == 0 {
		fmt.Println("companyBanch 是空的")
	}
}
func SelectShiftAll() {
	res := (*redis.RedisSingleton()).SelectShiftAll()
	if len(*res) == 0 {
		fmt.Println("shift 是空的")
	}
}
func SelectShiftChangeAll() {
	res := (*redis.RedisSingleton()).SelectShiftChangeAll()
	if len(*res) == 0 {
		fmt.Println("shiftChange 是空的")
	}
}
func SelectShiftOverTimeAll() {
	res := (*redis.RedisSingleton()).SelectShiftOverTimeAll()
	if len(*res) == 0 {
		fmt.Println("shiftOverTime 是空的")
	}
}
func SelectDayOffAll() {
	res := (*redis.RedisSingleton()).SelectDayOffAll()
	if len(*res) == 0 {
		fmt.Println("dayOff 是空的")
	}
}
func SelectForgetPunchAll() {
	res := (*redis.RedisSingleton()).SelectForgetPunchAll()
	if len(*res) == 0 {
		fmt.Println("ForgetPunch 是空的")
	}
}
func SelectLateExcusedAll() {
	res := (*redis.RedisSingleton()).SelectLateExcusedAll()
	if len(*res) == 0 {
		fmt.Println("lateExcused 是空的")
	}
}
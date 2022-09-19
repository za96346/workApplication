package handler

import (
	"backend/database"
	"backend/redis"
)

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
}
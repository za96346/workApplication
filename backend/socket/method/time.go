package method

import (
	"backend/redis"
	"fmt"
	"time"
)

// 開始 結束 開始年 開始月
func GetNextMonthSE() (string, string, int, int) {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start := thisMonth.AddDate(0, 1, 0)
	end := thisMonth.AddDate(0, 2, -1)
	return start.Format("2006-01-02"), end.Format("2006-01-02"), start.Year(), int(start.Month())
}
//   step 1 => 每個月的1號到14號的等待階段
//  step 2 => 每個月的倒數第15天，開啟編輯
//   step 3 => 每個月的倒數第5天 到 第3天，一般人員結束編輯，進入組長以上的確認階段
//   step 4 => 結算，禁止任何人編輯
func CheckWhichStep (banchId int64) int {
	year, month, day := time.Now().Date()
	_, _, startYear, startMonth := GetNextMonthSE()

	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	_, _, lastDay := thisMonth.AddDate(0, 1, -1).Date()
	diff := lastDay - day
	// fmt.Println("相差幾天 =>", diff)

	// 這邊要去檢查 redis room status
	// 如果 有 提交了 本月的資料 就 回傳 4
	v := (*redis.Singleton()).GetShiftRoomStatus(banchId)
	fullDate := fmt.Sprintln(startYear, "/", startMonth)

	// step 1
	if diff > 5 && diff <= 15 {
		return 2
	// step 2
	} else if (*v)["LastFinishedYearMonth"] == fullDate {
		return 4
	// step 3
	} else if diff <= 5 && diff >= 3 {
		return 3
	// step 4
	}  else {
		return 1
	}
}
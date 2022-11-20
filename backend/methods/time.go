package methods

import (
	"fmt"
	"time"
)

func GetNextMonthSE() (string, string) {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start := thisMonth.AddDate(0, 1, 0).Format("2006-01-02")
	end := thisMonth.AddDate(0, 2, -1).Format("2006-01-02")
	return start, end
}

//  step 1 => 每個月的倒數第15天，開啟編輯
//   step 2 => 每個月的倒數第5天 到 第3天，一般人員結束編輯，進入組長以上的確認階段
//   step 3 => 結算，禁止任何人編輯
//   step 4 => 每個月的1號到14號的等待階段
func CheckWhichStep () int {
	year, month, day := time.Now().Date()

	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	_, _, lastDay := thisMonth.AddDate(0, 1, -1).Date()
	diff := lastDay - day
	fmt.Println("相差幾天 =>", diff)
	// step 1
	if diff > 5 && diff <= 15 {
		return 1
	// step 2
	} else if diff <= 5 && diff >= 3 {
		return 2
	// step 3
	} else if diff < 3 {
		return 3
	//step 4
	} else {
		return 4
	}
}
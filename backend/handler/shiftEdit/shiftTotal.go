package shiftEdit

import (
	panichandler "backend/panicHandler"
	"backend/response"
	"math"
	"time"
)

// 班表總計
func ShiftTotal(data *[]response.Shift) (*map[int64]float64, *map[string]float64) {
	defer panichandler.Recover()

	userShiftTotalContainer := map[int64]float64{} // rows 的總時數
	dayShiftTotalContainer := map[string]float64{} // columns 的總時數

	for _, item := range *data {
		// if item.OffShiftTime > item.OnShiftTime {

		// }
		onDate, _ :=  time.Parse("2006-01-02 15:04:05", item.Date + " " + item.OnShiftTime)
		offDate, _ :=  time.Parse("2006-01-02 15:04:05", item.Date + " " + item.OffShiftTime)
		restTime, _ := time.Parse("15:04:05", item.RestTime)
		basicTime, _ := time.Parse("15:04:05", "00:00:00") // 給休息時間相減

		// 判斷有沒有過一天 過了一天要在加24小時
		if offDate.Before(onDate) {
			d, _ := time.ParseDuration("24h")
			offDate = offDate.Add(d)
		}
		subHour := (offDate.Sub(onDate).Seconds() - restTime.Sub(basicTime).Seconds()) / 60 / 60

		userShiftTotalContainer[item.UserId] = math.Round((subHour + userShiftTotalContainer[item.UserId]) * 10) / 10
		dayShiftTotalContainer[item.Date] = math.Round((subHour + dayShiftTotalContainer[item.Date]) * 10) / 10
	}
	return &userShiftTotalContainer, &dayShiftTotalContainer
}
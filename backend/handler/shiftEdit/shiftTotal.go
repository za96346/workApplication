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

		userShiftTotalContainer[item.UserId] = math.Round((offDate.Sub(onDate).Hours() + userShiftTotalContainer[item.UserId]) * 10) / 10
		dayShiftTotalContainer[item.Date] = math.Round((offDate.Sub(onDate).Hours() + dayShiftTotalContainer[item.Date]) * 10) / 10
	}
	return &userShiftTotalContainer, &dayShiftTotalContainer
}
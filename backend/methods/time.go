package methods

import "time"

func GetNextMonthSE() (string, string) {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	start := thisMonth.AddDate(0, 1, 0).Format("2006-01-02")
	end := thisMonth.AddDate(0, 2, -1).Format("2006-01-02")
	return start, end
}
package repository


type LogRepository interface {
	GetNewLogId(int) int
}
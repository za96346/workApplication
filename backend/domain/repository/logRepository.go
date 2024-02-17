package repository


type LogRepository interface {
	GetNewLogId(companyId int) int
}
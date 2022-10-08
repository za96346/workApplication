package handler

import "backend/table"

type Method interface {
	FindBanch()
}

func(dbObj *DB) FindBanch(banchId int64) string {
	banch := ""
	if b := (*dbObj).SelectCompanyBanch(2, banchId); len(*b) == 0 {
		banch = ""
	} else {
		banch = (*b)[0].BanchName
	}
	return banch
}

func(dbObj *DB) FindCompany(companyCode string) *[]table.CompanyTable {
	return (*dbObj).SelectCompany(2, companyCode);
}
package database

import (
	"fmt"
	"sync"
)
var userTableMux sync.RWMutex
var UserTableInstance *UserTable

type UserTable struct {
	UserId float64
	CompanyCode string
	Account string
	Password string
	OnWorkDay string
	Banch string
	Permession string
	Work_state string
	CreateTime string
	LastModify string
	MonthSalary float64
	PartTimeSalary float64
}
type CompanyTable struct {
	CompanyId float64
	CompanyCode string
	CompanyName string
	CompanyLocation string
	CompanyPhoneNumber string
	TermStart string
	TermEnd string
	CreateTime string
	LastModify string
}

func UserTableSingleton() *UserTable {
	if UserTableInstance == nil {
		userTableMux.Lock()
		defer userTableMux.Unlock()
		if UserTableInstance == nil {
			UserTableInstance = &UserTable{}
			return UserTableInstance
		}
	}
	return UserTableInstance
}
var queryMux *sync.Mutex
var SqlQuery *sqlQuery
type sqlQuery struct {
	User userQuery
	UserPreference userPerferenceQuery
	Company companyQuery
	CompanyBanch companyBanchQuery
}

type userQuery struct {
	InsertAll string
	InsertSingle string
	selectAll string
	selectSingle string
}
type userPerferenceQuery struct {
	InsertAll string
	InsertSingle string
}
type companyQuery struct {
	InsertAll string
	InsertSingle string
	SelectSingle string
}
type companyBanchQuery struct {
	InsertAll string
}

func SqlQuerySingleton() *sqlQuery {
	queryMux = new(sync.Mutex)
	if SqlQuery == nil {
		queryMux.Lock()
		defer queryMux.Unlock()
		if SqlQuery == nil {
			SqlQuery = &sqlQuery{}
			addUserQuery()
			addCompanyQuery()
			addUserPreferenceQuery()
			addCompanyBanchQuery()
			return SqlQuery
		}
	}
	return SqlQuery
}

func addUserQuery() {
	fmt.Println(SqlQuery)
	SqlQuery.User.InsertAll = `insert into user(
		companyCode,
		account,
		password,
		onWorkDay,
		banch,
		permession,
		workState,
		createTime,
		lastModify,
		monthSalary,
		partTimeSalary
		) values(
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	);`;
	SqlQuery.User.selectAll = `select * from user;`;
	SqlQuery.User.selectSingle = `select * from user where account=? or userId=?;`;
}
func addUserPreferenceQuery() {
	SqlQuery.UserPreference.InsertAll = `insert into userPreference(
		userId,
		style,
		fontSize,
		selfPhoto,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?
	);`;
}
func addCompanyQuery() {
	SqlQuery.Company.InsertAll = `insert into company(
		companyCode,
		companyName,
		companyLocation,
		companyPhoneNumber,
		termStart,
		termEnd,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?, ?, ?, ?
	);`;
	SqlQuery.Company.SelectSingle = `select * from company where companyId=? or companyCode=?;`;
}
func addCompanyBanchQuery() {
	SqlQuery.CompanyBanch.InsertAll = `insert into companyBanch(
		companyId,
		banchName,
		banchShiftStyle,
		createTime,
		lastModify
		) values(
			?, ?, ?, ?, ?
	);`;
}
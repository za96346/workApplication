package mysql

import (
	"backend/table"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)
func DataBaseInit() {
	defer simulateData();
	(*Singleton()).MysqlDB.Exec(`drop table user;`)
	(*Singleton()).MysqlDB.Exec(`drop table userPreference;`)
	(*Singleton()).MysqlDB.Exec(`drop table shift;`)
	(*Singleton()).MysqlDB.Exec(`drop table shiftChange;`)
	(*Singleton()).MysqlDB.Exec(`drop table shiftOverTime;`)
	(*Singleton()).MysqlDB.Exec(`drop table forgetPunch;`)
	(*Singleton()).MysqlDB.Exec(`drop table dayOff;`)
	(*Singleton()).MysqlDB.Exec(`drop table lateExcused;`)
	(*Singleton()).MysqlDB.Exec(`drop table company;`)
	(*Singleton()).MysqlDB.Exec(`drop table companyBanch;`)
/// user table
	_, err := (*Singleton()).MysqlDB.Exec(`
		create table user(
			userId bigint not null AUTO_INCREMENT unique,
			companyCode varchar(50),
			account varchar(50) primary key,
			password varchar(50),
			onWorkDay timestamp,
			banch bigint,
			permession int,
			workState varchar(50),
			createTime timestamp,
			lastModify timestamp,
			monthSalary int,
			partTimeSalary int
		);
	`)
	if err == nil {
		fmt.Println("user table is created success")
	} else {
		fmt.Println("user table is created failed")
		log.Fatal(err)
	}
/// userPreference table
	_, err = (*Singleton()).MysqlDB.Exec(`
		create table userPreference(
			userId bigint primary key,
			style varchar(50),
			fontSize varchar(3),
			selfPhoto blob,
			createTime timestamp,
			lastModify timestamp
		);
	`)
	if err == nil {
		fmt.Println("userPreference table is created success")
	} else {
		fmt.Println("userPreference table is created failed")
		log.Fatal(err)
	}
///shift table
	_, err = (*Singleton()).MysqlDB.Exec(`
		create table shift(
			shiftId bigint not null AUTO_INCREMENT unique,
			userId bigint,
			onShiftTime timestamp,
			offShiftTime timestamp,
			punchIn timestamp,
			punchOut timestamp,
			specifyTag varchar(50),
			createTime timestamp,
			lastModify timestamp
		);
	`)
	if err == nil {
		fmt.Println("shift table is created success")
	} else {
		fmt.Println("shift table is created failed")
		log.Fatal(err)
	}
///shiftChange table
	_, err = (*Singleton()).MysqlDB.Exec(`
	create table shiftChange(
		caseId bigint not null unique auto_increment,
		initiatorShiftId bigint,
		requestedShiftId bigint,
		reason varchar(200),
		caseProcess varchar(10),
		specifyTag varchar(50),
		createTime timestamp,
		lastModify timestamp
	);
	`)
	if err == nil {
		fmt.Println("shift table is created success")
	} else {
		fmt.Println("shift table is created failed")
		log.Fatal(err)
	}

///shiftOverTime table
	_, err = (*Singleton()).MysqlDB.Exec(`
	create table shiftOverTime(
		caseId bigint not null unique auto_increment,
		shiftId bigint,
		initiatorOnOverTime timestamp,
		initiatorOffOverTime timestamp,
		reason varchar(200),
		caseProcess varchar(10),
		specifyTag varchar(50),
		createTime timestamp,
		lastModify timestamp
	);
	`)
	if err == nil {
		fmt.Println("shiftOverTime table is created success")
	} else {
		fmt.Println("shiftOverTime table is created failed")
		log.Fatal(err)
	}
///forgetPunch table
	_, err =  (*Singleton()).MysqlDB.Exec(`
		create table forgetPunch(
			caseId bigint not null unique auto_increment,
			shiftId bigint,
			targetPunch varchar(3),
			reason varchar(200),
			caseProcess varchar(10),
			specifyTag varchar(50),
			createTime timestamp,
			lastModify timestamp
		);
	`)
	if err == nil {
		fmt.Println("forgetPunch table is created success")
	} else {
		fmt.Println("forgetPunch table is created failed")
		log.Fatal(err)
	}
/// dayOff table
	_, err = (*Singleton()).MysqlDB.Exec(`
		create table dayOff(
			caseId bigint not null unique auto_increment,
			shiftId bigint,
			dayOffType varchar(10),
			reason varchar(200),
			caseProcess varchar(10),
			specifyTag varchar(50),
			createTime timestamp,
			lastModify timestamp
		);
	`)
	if err == nil {
		fmt.Println("dayOff table is created success")
	} else {
		fmt.Println("dayOff table is created failed")
		log.Fatal(err)
	}

/// lateExcused table
	_, err = (*Singleton()).MysqlDB.Exec(`
		create table lateExcused(
			caseId bigint not null unique auto_increment,
			shiftId bigint,
			lateExcusedType varchar(10),
			reason varchar(200),
			caseProcess varchar(10),
			specifyTag varchar(50),
			createTime timestamp,
			lastModify timestamp
		);
	`)
	if err == nil {
		fmt.Println("lateExcused table is created success")
	} else {
		fmt.Println("lateExcused table is created failed")
		log.Fatal(err)
	}

/// company table
	_, err = (*Singleton()).MysqlDB.Exec(`
		create table company(
			companyId bigint not null unique auto_increment,
			companyCode varchar(50) unique,
			companyName varchar(200),
			companyLocation varchar(200),
			companyPhoneNumber varchar(20),
			termStart timestamp,
			termEnd timestamp,
			createTime timestamp,
			lastModify timestamp
		);
	`)
	if err == nil {
		fmt.Println("company table is created success")
	} else {
		fmt.Println("company table is created failed")
		log.Fatal(err)
	}

/// companyBanch table
	_, err = (*Singleton()).MysqlDB.Exec(`
		create table companyBanch(
			id bigint not null unique auto_increment,
			companyId bigint,
			banchName varchar(50),
			banchShiftStyle varchar(200),
			createTime timestamp,
			lastModify timestamp
		);
	`)
	if err == nil {
		fmt.Println("companyBanch table is created success")
	} else {
		fmt.Println("companyBanch table is created failed")
		log.Fatal(err)
	}


/// userPreference alter
	_, err = (*Singleton()).MysqlDB.Exec("alter table userPreference add foreign key(userId) references user(userId) on update cascade on delete cascade;")

	if err == nil {
		fmt.Println("userPreference alter foreign key success")
	} else {
		fmt.Println("userPreference alter foreign key failed")
		log.Fatal(err)
	}
/// shift alter
	_, err = (*Singleton()).MysqlDB.Exec(`
		alter table shift add primary key (onShiftTime, offShiftTime, shiftId, userId);`)

	if err == nil {
		fmt.Println("shift alter primary key success")
	} else {
		fmt.Println("shift alter primary key failed")
		log.Fatal(err)
	}
	
	_, err = (*Singleton()).MysqlDB.Exec("alter table shift add foreign key(userId) references user(userId) on update cascade on delete cascade;")

	if err == nil {
		fmt.Println("shift alter foreign key success")
	} else {
		fmt.Println("shift alter foreign key failed")
		log.Fatal(err)
	}

	_, err = (*Singleton()).MysqlDB.Exec("alter table shift auto_increment=1;")

	if err == nil {
		fmt.Println("shift alter set auto increment success")
	} else {
		fmt.Println("shift alter auto increment failed")
		log.Fatal(err)
	}

/// shiftChange alter
	_, err = (*Singleton()).MysqlDB.Exec(`
	alter table shiftChange add primary key (initiatorShiftId, requestedShiftId);`)

	if err == nil {
		fmt.Println("shiftChange alter primary key success")
	} else {
		fmt.Println("shiftChange alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*Singleton()).MysqlDB.Exec("alter table shiftChange add foreign key(initiatorShiftId) references shift(shiftId) on update cascade on delete cascade;")

	if err == nil {
		fmt.Println("shiftChange alter foreign key success")
	} else {
		fmt.Println("shiftChange alter foreign key failed")
		log.Fatal(err)
	}

	_, err = (*Singleton()).MysqlDB.Exec("alter table shiftChange add foreign key (requestedShiftId) references shift(shiftId) on update cascade on delete cascade;")

	if err == nil {
		fmt.Println("shiftChange alter foreign key success")
	} else {
		fmt.Println("shiftChange alter foreign key failed")
		log.Fatal(err)
	}
/// shiftOverTime alter
	_, err = (*Singleton()).MysqlDB.Exec(`
	alter table shiftOverTime add primary key (shiftId, initiatorOnOverTime, initiatorOffOverTime);`)

	if err == nil {
		fmt.Println("shiftOverTime alter primary key success")
	} else {
		fmt.Println("shiftOverTime alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*Singleton()).MysqlDB.Exec("alter table shiftOverTime add foreign key (shiftId) references shift(shiftId) on update cascade on delete cascade;")

	if err == nil {
		fmt.Println("shiftOverTime alter foreign key success")
	} else {
		fmt.Println("shiftOverTime alter foreign key failed")
		log.Fatal(err)
	}
/// forgetPunch alter
	_, err = (*Singleton()).MysqlDB.Exec(`
	alter table forgetPunch add primary key (shiftId, targetPunch);`)

	if err == nil {
		fmt.Println("forgetPunch alter primary key success")
	} else {
		fmt.Println("forgetPunch alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*Singleton()).MysqlDB.Exec(`alter table forgetPunch add foreign key (shiftId) references shift(shiftId) on update cascade on delete cascade;`)

	if err == nil {
		fmt.Println("forgetPunch alter foreign key success")
	} else {
		fmt.Println("forgetPunch alter foreign key failed")
		log.Fatal(err)
	}
/// dayOff alter
	_, err = (*Singleton()).MysqlDB.Exec(`alter table dayOff add primary key (shiftId, dayOffType);`)

	if err == nil {
		fmt.Println("dayOff alter primary key success")
	} else {
		fmt.Println("dayOff alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*Singleton()).MysqlDB.Exec(`alter table dayOff add foreign key (shiftId) references shift(shiftId) on update cascade on delete cascade;`)

	if err == nil {
		fmt.Println("dayOff alter foreign key success")
	} else {
		fmt.Println("dayOff alter foreign key failed")
		log.Fatal(err)
	}
/// lateExecused alter
	_, err = (*Singleton()).MysqlDB.Exec(`alter table lateExcused add primary key (shiftId, lateExcusedType);`)

	if err == nil {
		fmt.Println("lateExcused alter primary key success")
	} else {
		fmt.Println("lateExcused alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*Singleton()).MysqlDB.Exec(`alter table lateExcused add foreign key (shiftId) references shift(shiftId) on update cascade on delete cascade;`)

	if err == nil {
		fmt.Println("lateExcused alter foreign key success")
	} else {
		fmt.Println("lateExcused alter foreign key failed")
		log.Fatal(err)
	}
/// company alter
	_, err = (*Singleton()).MysqlDB.Exec(`alter table company add primary key (companyId, companyCode);`)

	if err == nil {
		fmt.Println("company alter primary key success")
	} else {
		fmt.Println("company alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*Singleton()).MysqlDB.Exec("alter table company auto_increment=1;")

	if err == nil {
		fmt.Println("company alter set auto increment success")
	} else {
		fmt.Println("company alter set auto increment failed")
		log.Fatal(err)
	}

/// companyBanch alter
	_, err = (*Singleton()).MysqlDB.Exec(`alter table companyBanch add primary key (companyId, banchName);`)

	if err == nil {
		fmt.Println("companyBanch alter primary key success")
	} else {
		fmt.Println("companyBanch alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*Singleton()).MysqlDB.Exec(`alter table companyBanch add foreign key (companyId) references company(companyId) on update cascade on delete cascade;`)

	if err == nil {
		fmt.Println("companyBanch alter foreign key success")
	} else {
		fmt.Println("companyBanch alter foreign key failed")
		log.Fatal(err)
	}
// user alter
// companyCode 被刪除後 set null
	_, err = (*Singleton()).MysqlDB.Exec("alter table user auto_increment=1;")

	if err == nil {
		fmt.Println("user alter set auto increment success")
	} else {
		fmt.Println("user alter set auto increment failed")
		log.Fatal(err)
	}

	_, err = (*Singleton()).MysqlDB.Exec(`alter table user add foreign key (companyCode) references company(companyCode) on update cascade on delete set null;`)

	if err == nil {
		fmt.Println("user alter foreign key success")
	} else {
		fmt.Println("user alter foreign key failed")
		log.Fatal(err)
	}

}
func simulateData() {
	wg := new(sync.WaitGroup)
    (*wg).Add(1)
	go func ()  {
		//company
		company := table.CompanyTable{
			CompanyCode: "",
			CompanyName: "",
			CompanyLocation: "",
			CompanyPhoneNumber: "",
			TermStart: time.Now(),
			TermEnd: time.Now(),
			CreateTime: time.Now(),
			LastModify: time.Now(),
		}
		(*Singleton()).InsertCompany(&company)
		company = table.CompanyTable{
			CompanyCode: "fei32fej",
			CompanyName: "xx股份有限公司",
			CompanyLocation: "台中市大甲區ｘｘｘ",
			CompanyPhoneNumber: "0906930873",
			TermStart: time.Now(),
			TermEnd: time.Now(),
			CreateTime: time.Now(),
			LastModify: time.Now(),
		}
		resStatus, _ := (*Singleton()).InsertCompany(&company)
		handleError(resStatus)
		// company banch
		resData := (*Singleton()).SelectCompany(2, "fei32fej")
		fmt.Println("接收SelectCompanySingle 記憶體位置 => ", resData, "\n")
		companyBanch := table.CompanyBanchTable{
			CompanyId :(*resData)[0].CompanyId,
			BanchName: "公關組",
			BanchShiftStyle: "{}",
			CreateTime: time.Now(),
			LastModify: time.Now(),
		}
		resStatus, _ = (*Singleton()).InsertCompanyBanch(&companyBanch)
		handleError(resStatus)
		(*wg).Done()
	}()
	(*wg).Wait()
	(*wg).Add(11)
	for i := 0; i <= 10; i++ {
		i := i
		go func (i int)  {
			// user
			user := table.UserTable{
				UserId : int64(0),
				CompanyCode: "fei32fej",
				Account: "account" + strconv.Itoa(i),
				Password: "aa20010722",
				OnWorkDay: time.Now(),
				Banch: 1,
				Permession: 0,
				WorkState: "on",
				MonthSalary: 30000 + i,
				PartTimeSalary: 130 + i,
				CreateTime: time.Now(),
				LastModify: time.Now(),
			}
			resStatus, _ := (*Singleton()).InsertUser(&user)
			// userPreference
			resData := (*Singleton()).SelectUser(2, "account" + strconv.Itoa(i)) //拿使用者資料
			userPreference := table.UserPreferenceTable{
				UserId: (*resData)[0].UserId,
				Style: "{style}",
				FontSize: "12",
				SelfPhoto: "pic",
				CreateTime: time.Now(),
				LastModify: time.Now(),
			}
			resStatus, _ = (*Singleton()).InsertUserPreference(&userPreference)
			handleError(resStatus)
			// shift
			for shiftStep := 0; shiftStep <= 30; shiftStep++ {
				hours, _ :=  time.ParseDuration("1h")
				oneDay, _ := time.ParseDuration(fmt.Sprint(strconv.Itoa(shiftStep * 24), "h"))
				shift := table.ShiftTable{
					UserId: (*resData)[0].UserId,
					OnShiftTime: time.Now().Add(-8 * hours).Add(oneDay),
					OffShiftTime: time.Now().Add(oneDay),
					SpecifyTag: "nothing",
					PunchIn: time.Now(),
					PunchOut: time.Now(),
					CreateTime: time.Now(),
					LastModify: time.Now(),
				}
				_, _ = (*Singleton()).InsertShift(&shift)
			}
			wg.Done()
		}(i)
	}
	// shift change
	shiftChange := table.ShiftChangeTable{
		InitiatorShiftId: int64(1),
		RequestedShiftId: int64(2),
		Reason: "我那天有事",
		CaseProcess: "manager",
		SpecifyTag: "hi",
		CreateTime: time.Now(),
		LastModify: time.Now(),
	}
	shiftOverTime := table.ShiftOverTimeTable{
		ShiftId: int64(1),
		InitiatorOnOverTime: time.Now(),
		InitiatorOffOverTime: time.Now(),
		Reason: "妹有原因",
		CaseProcess: "manager",
		SpecifyTag: "hi",
		CreateTime: time.Now(),
		LastModify: time.Now(),
	}
	forgetPunch := table.ForgetPunchTable{
		ShiftId: int64(1),
		TargetPunch: "上班",
		Reason: "妹有原因",
		CaseProcess: "manager",
		SpecifyTag: "hi",
		CreateTime: time.Now(),
		LastModify: time.Now(),
	}
	dayOff := table.DayOffTable{
		ShiftId: int64(1),
		DayOffType: "事假",
		Reason: "妹有原因",
		CaseProcess: "manager",
		SpecifyTag: "hi",
		CreateTime: time.Now(),
		LastModify: time.Now(),
	}
	lateExcused := table.LateExcusedTable{
		ShiftId: int64(1),
		LateExcusedType: "遲到",
		Reason: "妹有原因",
		CaseProcess: "manager",
		SpecifyTag: "hi",
		CreateTime: time.Now(),
		LastModify: time.Now(),
	}
	(*wg).Wait()
	(*Singleton()).InsertShiftChange(&shiftChange)
	(*Singleton()).InsertShiftOverTime(&shiftOverTime)
	(*Singleton()).InsertForgetPunch(&forgetPunch)
	(*Singleton()).InsertDayOff(&dayOff)
	(*Singleton()).InsertLateExcused(&lateExcused)
}
func handleError(resStatus bool) {
	if resStatus {
		// do success handle
	} else {
		//do fail handle
	}
}
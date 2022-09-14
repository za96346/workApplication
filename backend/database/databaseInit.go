package database

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"sync"
)
func DataBaseInit() {
	defer simulateData();
	MysqlDB.Exec(`drop table user;`)
	MysqlDB.Exec(`drop table userPreference;`)
	MysqlDB.Exec(`drop table shift;`)
	MysqlDB.Exec(`drop table shiftChange;`)
	MysqlDB.Exec(`drop table shiftOverTime;`)
	MysqlDB.Exec(`drop table forgetPunch;`)
	MysqlDB.Exec(`drop table dayOff;`)
	MysqlDB.Exec(`drop table lateExcused;`)
	MysqlDB.Exec(`drop table company;`)
	MysqlDB.Exec(`drop table companyBanch;`)
/// user table
	_, err = MysqlDB.Exec(`
		create table user(
			userId int not null AUTO_INCREMENT unique,
			companyCode varchar(50),
			account varchar(50) primary key,
			password varchar(50),
			onWorkDay varchar(50),
			banch varchar(50),
			permession varchar(50),
			workState varchar(50),
			createTime varchar(50),
			lastModify varchar(50),
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
	_, err = MysqlDB.Exec(`
		create table userPreference(
			userId int primary key,
			style varchar(50),
			fontSize varchar(3),
			selfPhoto blob,
			createTime varchar(50),
			lastModify varchar(50)
		);
	`)
	if err == nil {
		fmt.Println("userPreference table is created success")
	} else {
		fmt.Println("userPreference table is created failed")
		log.Fatal(err)
	}
///shift table
	_, err = MysqlDB.Exec(`
		create table shift(
			shiftId int unique,
			userId int,
			onShiftTime varchar(50),
			offShiftTime varchar(50),
			punchIn varchar(50),
			punchOut varchar(50),
			specifyTag varchar(50),
			createTime varchar(50),
			lastModify varchar(50)
		);
	`)
	if err == nil {
		fmt.Println("shift table is created success")
	} else {
		fmt.Println("shift table is created failed")
		log.Fatal(err)
	}
///shiftChange table
	_, err = MysqlDB.Exec(`
	create table shiftChange(
		initiatorShiftId int,
		requestedShiftId int,
		reson varchar(200),
		caseProcess varchar(10),
		specifyTag varchar(50),
		createTime varchar(50),
		lastModify varchar(50)
	);
	`)
	if err == nil {
		fmt.Println("shift table is created success")
	} else {
		fmt.Println("shift table is created failed")
		log.Fatal(err)
	}

///shiftOverTime table
	_, err = MysqlDB.Exec(`
	create table shiftOverTime(
		shiftId int,
		initiatorOnOverTime varchar(50),
		initiatorOffOverTime varchar(50),
		reson varchar(200),
		caseProcess varchar(10),
		specifyTag varchar(50),
		createTime varchar(50),
		lastModify varchar(50)
	);
	`)
	if err == nil {
		fmt.Println("shiftOverTime table is created success")
	} else {
		fmt.Println("shiftOverTime table is created failed")
		log.Fatal(err)
	}
///forgetPunch table
	_, err = MysqlDB.Exec(`
		create table forgetPunch(
			shiftId int,
			targetPunch varchar(3),
			reson varchar(200),
			caseProcess varchar(10),
			specifyTag varchar(50),
			createTime varchar(50),
			lastModify varchar(50)
		);
	`)
	if err == nil {
		fmt.Println("forgetPunch table is created success")
	} else {
		fmt.Println("forgetPunch table is created failed")
		log.Fatal(err)
	}
/// dayOff table
	_, err = MysqlDB.Exec(`
		create table dayOff(
			shiftId int,
			dayOffType varchar(10),
			reson varchar(200),
			caseProcess varchar(10),
			specifyTag varchar(50),
			createTime varchar(50),
			lastModify varchar(50)
		);
	`)
	if err == nil {
		fmt.Println("dayOff table is created success")
	} else {
		fmt.Println("dayOff table is created failed")
		log.Fatal(err)
	}

/// lateExcused table
	_, err = MysqlDB.Exec(`
		create table lateExcused(
			shiftId int,
			lateExcusedType varchar(10),
			reson varchar(200),
			caseProcess varchar(10),
			specifyTag varchar(50),
			createTime varchar(50),
			lastModify varchar(50)
		);
	`)
	if err == nil {
		fmt.Println("lateExcused table is created success")
	} else {
		fmt.Println("lateExcused table is created failed")
		log.Fatal(err)
	}

/// company table
	_, err = MysqlDB.Exec(`
		create table company(
			companyId int unique auto_increment,
			companyCode varchar(50) unique,
			companyName varchar(200),
			companyLocation varchar(200),
			companyPhoneNumber varchar(20),
			termStart varchar(50),
			termEnd varchar(50),
			createTime varchar(50),
			lastModify varchar(50)
		);
	`)
	if err == nil {
		fmt.Println("company table is created success")
	} else {
		fmt.Println("company table is created failed")
		log.Fatal(err)
	}

/// companyBanch table
	_, err = MysqlDB.Exec(`
		create table companyBanch(
			companyId int,
			banchName varchar(50),
			banchShiftStyle varchar(200),
			createTime varchar(50),
			lastModify varchar(50)
		);
	`)
	if err == nil {
		fmt.Println("companyBanch table is created success")
	} else {
		fmt.Println("companyBanch table is created failed")
		log.Fatal(err)
	}


/// userPreference alter
	_, err = MysqlDB.Exec("alter table userPreference add foreign key(userId) references user(userId) on update cascade;")

	if err == nil {
		fmt.Println("userPreference alter foreign key success")
	} else {
		fmt.Println("userPreference alter foreign key failed")
		log.Fatal(err)
	}
/// shift alter
	_, err = MysqlDB.Exec(`
		alter table shift add primary key (onShiftTime, offShiftTime, shiftId, userId);`)

	if err == nil {
		fmt.Println("shift alter primary key success")
	} else {
		fmt.Println("shift alter primary key failed")
		log.Fatal(err)
	}
	
	_, err = MysqlDB.Exec("alter table shift add foreign key(userId) references user(userId) on update cascade;")

	if err == nil {
		fmt.Println("shift alter foreign key success")
	} else {
		fmt.Println("shift alter foreign key failed")
		log.Fatal(err)
	}

	_, err = MysqlDB.Exec("alter table shift auto_increment=734028;")

	if err == nil {
		fmt.Println("shift alter set auto increment success")
	} else {
		fmt.Println("shift alter auto increment failed")
		log.Fatal(err)
	}

/// shiftChange alter
	_, err = MysqlDB.Exec(`
	alter table shiftChange add primary key (initiatorShiftId, requestedShiftId);`)

	if err == nil {
		fmt.Println("shiftChange alter primary key success")
	} else {
		fmt.Println("shiftChange alter primary key failed")
		log.Fatal(err)
	}

	_, err = MysqlDB.Exec("alter table shiftChange add foreign key(initiatorShiftId) references shift(shiftId) on update cascade;")

	if err == nil {
		fmt.Println("shiftChange alter foreign key success")
	} else {
		fmt.Println("shiftChange alter foreign key failed")
		log.Fatal(err)
	}

	_, err = MysqlDB.Exec("alter table shiftChange add foreign key (requestedShiftId) references shift(shiftId) on update cascade;")

	if err == nil {
		fmt.Println("shiftChange alter foreign key success")
	} else {
		fmt.Println("shiftChange alter foreign key failed")
		log.Fatal(err)
	}
/// shiftOverTime alter
	_, err = MysqlDB.Exec(`
	alter table shiftOverTime add primary key (shiftId, initiatorOnOverTime, initiatorOffOverTime);`)

	if err == nil {
		fmt.Println("shiftOverTime alter primary key success")
	} else {
		fmt.Println("shiftOverTime alter primary key failed")
		log.Fatal(err)
	}

	_, err = MysqlDB.Exec("alter table shiftOverTime add foreign key (shiftId) references shift(shiftId) on update cascade;")

	if err == nil {
		fmt.Println("shiftOverTime alter foreign key success")
	} else {
		fmt.Println("shiftOverTime alter foreign key failed")
		log.Fatal(err)
	}
/// forgetPunch alter
	_, err = MysqlDB.Exec(`
	alter table forgetPunch add primary key (shiftId, targetPunch);`)

	if err == nil {
		fmt.Println("forgetPunch alter primary key success")
	} else {
		fmt.Println("forgetPunch alter primary key failed")
		log.Fatal(err)
	}

	_, err = MysqlDB.Exec(`alter table forgetPunch add foreign key (shiftId) references shift(shiftId) on update cascade;`)

	if err == nil {
		fmt.Println("forgetPunch alter foreign key success")
	} else {
		fmt.Println("forgetPunch alter foreign key failed")
		log.Fatal(err)
	}
/// dayOff alter
	_, err = MysqlDB.Exec(`alter table dayOff add primary key (shiftId, dayOffType);`)

	if err == nil {
		fmt.Println("dayOff alter primary key success")
	} else {
		fmt.Println("dayOff alter primary key failed")
		log.Fatal(err)
	}

	_, err = MysqlDB.Exec(`alter table dayOff add foreign key (shiftId) references shift(shiftId) on update cascade;`)

	if err == nil {
		fmt.Println("dayOff alter foreign key success")
	} else {
		fmt.Println("dayOff alter foreign key failed")
		log.Fatal(err)
	}
/// lateExecused alter
	_, err = MysqlDB.Exec(`alter table lateExcused add primary key (shiftId, lateExcusedType);`)

	if err == nil {
		fmt.Println("lateExcused alter primary key success")
	} else {
		fmt.Println("lateExcused alter primary key failed")
		log.Fatal(err)
	}

	_, err = MysqlDB.Exec(`alter table lateExcused add foreign key (shiftId) references shift(shiftId) on update cascade;`)

	if err == nil {
		fmt.Println("lateExcused alter foreign key success")
	} else {
		fmt.Println("lateExcused alter foreign key failed")
		log.Fatal(err)
	}
/// company alter
	_, err = MysqlDB.Exec(`alter table company add primary key (companyId, companyCode);`)

	if err == nil {
		fmt.Println("company alter primary key success")
	} else {
		fmt.Println("company alter primary key failed")
		log.Fatal(err)
	}

	_, err = MysqlDB.Exec("alter table company auto_increment=3034;")

	if err == nil {
		fmt.Println("company alter set auto increment success")
	} else {
		fmt.Println("company alter set auto increment failed")
		log.Fatal(err)
	}

/// companyBanch alter
	_, err = MysqlDB.Exec(`alter table companyBanch add primary key (companyId, banchName);`)

	if err == nil {
		fmt.Println("companyBanch alter primary key success")
	} else {
		fmt.Println("companyBanch alter primary key failed")
		log.Fatal(err)
	}

	_, err = MysqlDB.Exec(`alter table companyBanch add foreign key (companyId) references company(companyId) on update cascade;`)

	if err == nil {
		fmt.Println("companyBanch alter foreign key success")
	} else {
		fmt.Println("companyBanch alter foreign key failed")
		log.Fatal(err)
	}
// user alter
	_, err = MysqlDB.Exec("alter table user auto_increment=34;")

	if err == nil {
		fmt.Println("user alter set auto increment success")
	} else {
		fmt.Println("user alter set auto increment failed")
		log.Fatal(err)
	}

	_, err = MysqlDB.Exec(`alter table user add foreign key (companyCode) references company(companyCode) on update cascade;`)

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
		resStatus = (*DBSingleton()).InsertCompanyAll(
			"fei32fej",
			"xx股份有限公司",
			"台中市大甲區ｘｘｘ",
			"0906930873",
			time.Now(),
			time.Now(),
			time.Now(),
			time.Now(),
		)
		if resStatus {
			// do success handle
		} else {
			//do fail handle
		}
		resData := (*DBSingleton()).SelectCompanySingle("fei32fej")
		fmt.Println("接收SelectCompanySingle 記憶體位置 => ", resData, "\n")
		resStatus = (*DBSingleton()).InsertCompanyBanchAll(
			(*resData).CompanyId,
			"公關組",
			"{}",
			time.Now(),
			time.Now(),
		)
		if resStatus {
			// do success handle
		} else {
			//do fail handle
		}
		(*wg).Done()
	}()
	(*wg).Wait()
	(*wg).Add(2)
	for i := 0; i <=1; i++ {
		i := i
		go func (i int)  {
			resStatus = (*DBSingleton()).InsertUserAll(
				"fei32fej",
				"account" + strconv.Itoa(i),
				"1234",
				"2022-02-11 10;10:00",
				"公關組",
				"admin",
				"on",
				time.Now(),
				time.Now(),
				30000 + i,
				130 + i,
			)
			resData := (*DBSingleton()).SelectUserSingle("account" + strconv.Itoa(i))
			fmt.Println("接收SelectUserSingle 記憶體位置 => ", resData, "\n")
			resStatus = (*DBSingleton()).InsertUserPreferenceAll(
				(*resData).UserId,
				"{style}",
				"12",
				"pic",
				time.Now(),
				time.Now(),
			)
			if resStatus {
				// do success handle
			} else {
				//do fail handle
			}
			wg.Done()
		}(i)
	}
	(*wg).Wait()
	resData := (*DBSingleton()).SelectUserAll()
	fmt.Println("接收SelectUserAll 記憶體位置 => ", resData, "\n")
	fmt.Println(" => ", (*resData)[0], (*resData)[1])
}
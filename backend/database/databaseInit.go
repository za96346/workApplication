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
	(*MysqlSingleton()).MysqlDB.Exec(`drop table user;`)
	(*MysqlSingleton()).MysqlDB.Exec(`drop table userPreference;`)
	(*MysqlSingleton()).MysqlDB.Exec(`drop table shift;`)
	(*MysqlSingleton()).MysqlDB.Exec(`drop table shiftChange;`)
	(*MysqlSingleton()).MysqlDB.Exec(`drop table shiftOverTime;`)
	(*MysqlSingleton()).MysqlDB.Exec(`drop table forgetPunch;`)
	(*MysqlSingleton()).MysqlDB.Exec(`drop table dayOff;`)
	(*MysqlSingleton()).MysqlDB.Exec(`drop table lateExcused;`)
	(*MysqlSingleton()).MysqlDB.Exec(`drop table company;`)
	(*MysqlSingleton()).MysqlDB.Exec(`drop table companyBanch;`)
/// user table
	_, err := (*MysqlSingleton()).MysqlDB.Exec(`
		create table user(
			userId int not null AUTO_INCREMENT unique,
			companyCode varchar(50),
			account varchar(50) primary key,
			password varchar(50),
			onWorkDay timestamp,
			banch varchar(50),
			permession varchar(50),
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
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`
		create table userPreference(
			userId int primary key,
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
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`
		create table shift(
			shiftId int not null AUTO_INCREMENT unique,
			userId int,
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
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`
	create table shiftChange(
		initiatorShiftId int,
		requestedShiftId int,
		reson varchar(200),
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
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`
	create table shiftOverTime(
		shiftId int,
		initiatorOnOverTime timestamp,
		initiatorOffOverTime timestamp,
		reson varchar(200),
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
	_, err =  (*MysqlSingleton()).MysqlDB.Exec(`
		create table forgetPunch(
			shiftId int,
			targetPunch varchar(3),
			reson varchar(200),
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
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`
		create table dayOff(
			shiftId int,
			dayOffType varchar(10),
			reson varchar(200),
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
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`
		create table lateExcused(
			shiftId int,
			lateExcusedType varchar(10),
			reson varchar(200),
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
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`
		create table company(
			companyId int not null unique auto_increment,
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
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`
		create table companyBanch(
			companyId int,
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
	_, err = (*MysqlSingleton()).MysqlDB.Exec("alter table userPreference add foreign key(userId) references user(userId) on update cascade;")

	if err == nil {
		fmt.Println("userPreference alter foreign key success")
	} else {
		fmt.Println("userPreference alter foreign key failed")
		log.Fatal(err)
	}
/// shift alter
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`
		alter table shift add primary key (onShiftTime, offShiftTime, shiftId, userId);`)

	if err == nil {
		fmt.Println("shift alter primary key success")
	} else {
		fmt.Println("shift alter primary key failed")
		log.Fatal(err)
	}
	
	_, err = (*MysqlSingleton()).MysqlDB.Exec("alter table shift add foreign key(userId) references user(userId) on update cascade;")

	if err == nil {
		fmt.Println("shift alter foreign key success")
	} else {
		fmt.Println("shift alter foreign key failed")
		log.Fatal(err)
	}

	_, err = (*MysqlSingleton()).MysqlDB.Exec("alter table shift auto_increment=1;")

	if err == nil {
		fmt.Println("shift alter set auto increment success")
	} else {
		fmt.Println("shift alter auto increment failed")
		log.Fatal(err)
	}

/// shiftChange alter
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`
	alter table shiftChange add primary key (initiatorShiftId, requestedShiftId);`)

	if err == nil {
		fmt.Println("shiftChange alter primary key success")
	} else {
		fmt.Println("shiftChange alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*MysqlSingleton()).MysqlDB.Exec("alter table shiftChange add foreign key(initiatorShiftId) references shift(shiftId) on update cascade;")

	if err == nil {
		fmt.Println("shiftChange alter foreign key success")
	} else {
		fmt.Println("shiftChange alter foreign key failed")
		log.Fatal(err)
	}

	_, err = (*MysqlSingleton()).MysqlDB.Exec("alter table shiftChange add foreign key (requestedShiftId) references shift(shiftId) on update cascade;")

	if err == nil {
		fmt.Println("shiftChange alter foreign key success")
	} else {
		fmt.Println("shiftChange alter foreign key failed")
		log.Fatal(err)
	}
/// shiftOverTime alter
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`
	alter table shiftOverTime add primary key (shiftId, initiatorOnOverTime, initiatorOffOverTime);`)

	if err == nil {
		fmt.Println("shiftOverTime alter primary key success")
	} else {
		fmt.Println("shiftOverTime alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*MysqlSingleton()).MysqlDB.Exec("alter table shiftOverTime add foreign key (shiftId) references shift(shiftId) on update cascade;")

	if err == nil {
		fmt.Println("shiftOverTime alter foreign key success")
	} else {
		fmt.Println("shiftOverTime alter foreign key failed")
		log.Fatal(err)
	}
/// forgetPunch alter
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`
	alter table forgetPunch add primary key (shiftId, targetPunch);`)

	if err == nil {
		fmt.Println("forgetPunch alter primary key success")
	} else {
		fmt.Println("forgetPunch alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*MysqlSingleton()).MysqlDB.Exec(`alter table forgetPunch add foreign key (shiftId) references shift(shiftId) on update cascade;`)

	if err == nil {
		fmt.Println("forgetPunch alter foreign key success")
	} else {
		fmt.Println("forgetPunch alter foreign key failed")
		log.Fatal(err)
	}
/// dayOff alter
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`alter table dayOff add primary key (shiftId, dayOffType);`)

	if err == nil {
		fmt.Println("dayOff alter primary key success")
	} else {
		fmt.Println("dayOff alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*MysqlSingleton()).MysqlDB.Exec(`alter table dayOff add foreign key (shiftId) references shift(shiftId) on update cascade;`)

	if err == nil {
		fmt.Println("dayOff alter foreign key success")
	} else {
		fmt.Println("dayOff alter foreign key failed")
		log.Fatal(err)
	}
/// lateExecused alter
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`alter table lateExcused add primary key (shiftId, lateExcusedType);`)

	if err == nil {
		fmt.Println("lateExcused alter primary key success")
	} else {
		fmt.Println("lateExcused alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*MysqlSingleton()).MysqlDB.Exec(`alter table lateExcused add foreign key (shiftId) references shift(shiftId) on update cascade;`)

	if err == nil {
		fmt.Println("lateExcused alter foreign key success")
	} else {
		fmt.Println("lateExcused alter foreign key failed")
		log.Fatal(err)
	}
/// company alter
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`alter table company add primary key (companyId, companyCode);`)

	if err == nil {
		fmt.Println("company alter primary key success")
	} else {
		fmt.Println("company alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*MysqlSingleton()).MysqlDB.Exec("alter table company auto_increment=1;")

	if err == nil {
		fmt.Println("company alter set auto increment success")
	} else {
		fmt.Println("company alter set auto increment failed")
		log.Fatal(err)
	}

/// companyBanch alter
	_, err = (*MysqlSingleton()).MysqlDB.Exec(`alter table companyBanch add primary key (companyId, banchName);`)

	if err == nil {
		fmt.Println("companyBanch alter primary key success")
	} else {
		fmt.Println("companyBanch alter primary key failed")
		log.Fatal(err)
	}

	_, err = (*MysqlSingleton()).MysqlDB.Exec(`alter table companyBanch add foreign key (companyId) references company(companyId) on update cascade;`)

	if err == nil {
		fmt.Println("companyBanch alter foreign key success")
	} else {
		fmt.Println("companyBanch alter foreign key failed")
		log.Fatal(err)
	}
// user alter
	_, err = (*MysqlSingleton()).MysqlDB.Exec("alter table user auto_increment=1;")

	if err == nil {
		fmt.Println("user alter set auto increment success")
	} else {
		fmt.Println("user alter set auto increment failed")
		log.Fatal(err)
	}

	_, err = (*MysqlSingleton()).MysqlDB.Exec(`alter table user add foreign key (companyCode) references company(companyCode) on update cascade;`)

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
		//公司
		resStatus := (*MysqlSingleton()).InsertCompanyAll(
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
		// 公司部門
		resData := (*MysqlSingleton()).SelectCompanySingle("fei32fej")
		fmt.Println("接收SelectCompanySingle 記憶體位置 => ", resData, "\n")
		resStatus = (*MysqlSingleton()).InsertCompanyBanchAll(
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
	(*wg).Add(11)
	for i := 0; i <= 10; i++ {
		i := i
		go func (i int)  {
			resStatus := (*MysqlSingleton()).InsertUserAll(
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
			resData := (*MysqlSingleton()).SelectUserSingle("account" + strconv.Itoa(i)) //拿使用者資料
			resStatus = (*MysqlSingleton()).InsertUserPreferenceAll(
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
			for shiftStep := 0; shiftStep <= 30; shiftStep++ {
				hours, _ :=  time.ParseDuration("1h")
				oneDay, _ := time.ParseDuration(fmt.Sprint(strconv.Itoa(shiftStep * 24), "h"))
				_ = (*MysqlSingleton()).InsertShiftAll(
					(*resData).UserId,
					time.Now().Add(-8 * hours).Add(oneDay),
					time.Now().Add(oneDay),
					time.Now(),
					time.Now(),
					time.Now(),
					time.Now(),
					"nothing",
				)
			}
			wg.Done()
		}(i)
	}
	(*wg).Wait()
	(*MysqlSingleton()).SelectCompanyAll()
	(*MysqlSingleton()).SelectCompanyBanchAll()
	(*MysqlSingleton()).SelectUserAll()
	(*MysqlSingleton()).SelectUserPreferenceAll()
	(*MysqlSingleton()).SelectShiftAll()
	(*MysqlSingleton()).SelectShiftChangeAll()
	(*MysqlSingleton()).SelectShiftOverTimeAll()
	(*MysqlSingleton()).SelectForgetPunchAll()
	(*MysqlSingleton()).SelectDayOffAll()
	(*MysqlSingleton()).SelectLateExcusedAll()
}
// func handleError(resStatus *MysqlSingleton()) {
// 	if resStatus {
// 		// do success handle
// 	} else {
// 		//do fail handle
// 	}
// }
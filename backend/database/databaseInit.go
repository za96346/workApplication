package database

import (
	"database/sql"
	"fmt"
	"log"
)

func DataBaseInit(DB *sql.DB) {
	// DB.Query(`drop database workApplication;`)
	// DB.Query(`create database workApplication;`)
	// DB.Query(`use workApplication;`)
	DB.Exec(`drop table user;`)
	DB.Exec(`drop table userPreference;`)
	DB.Exec(`drop table shift;`)
	DB.Exec(`drop table shiftChange;`)
	DB.Exec(`drop table shiftOverTime;`)
	DB.Exec(`drop table forgetPunch;`)
	DB.Exec(`drop table dayOff;`)
	DB.Exec(`drop table lateExcused;`)
	DB.Exec(`drop table company;`)
	DB.Exec(`drop table companyBanch;`)
/// user table
	_, err = DB.Exec(`
		create table user(
			userId int not null AUTO_INCREMENT unique,
			companyCode varchar(50),
			account varchar(50) primary key,
			password varchar(50),
			onWorkDay varchar(50),
			banch varchar(50),
			permession varchar(50),
			wrokState varchar(50),
			createTime varchar(50),
			lastModify varchar(50),
			monthSalary varchar(50),
			ParTimeSalary varchar(50)
		);
	`)
	if err == nil {
		fmt.Println("user table is created success")
	} else {
		fmt.Println("user table is created failed")
		log.Fatal(err)
	}
/// userPreference table
	_, err = DB.Exec(`
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
	_, err = DB.Exec(`
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
	_, err = DB.Exec(`
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
	_, err = DB.Exec(`
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
	_, err = DB.Exec(`
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
	_, err = DB.Exec(`
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
	_, err = DB.Exec(`
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
	_, err = DB.Exec(`
		create table company(
			companyId int unique,
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
	_, err = DB.Exec(`
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
	_, err = DB.Exec("alter table userPreference add foreign key(userId) references user(userId) on update cascade;")

	if err == nil {
		fmt.Println("userPreference alter foreign key success")
	} else {
		fmt.Println("userPreference alter foreign key failed")
		log.Fatal(err)
	}
/// shift alter
	_, err = DB.Exec(`
		alter table shift add primary key (onShiftTime, offShiftTime, shiftId, userId);`)

	if err == nil {
		fmt.Println("shift alter primary key success")
	} else {
		fmt.Println("shift alter primary key failed")
		log.Fatal(err)
	}
	
	_, err = DB.Exec("alter table shift add foreign key(userId) references user(userId) on update cascade;")

	if err == nil {
		fmt.Println("shift alter foreign key success")
	} else {
		fmt.Println("shift alter foreign key failed")
		log.Fatal(err)
	}

	_, err = DB.Exec("alter table shift auto_increment=734028;")

	if err == nil {
		fmt.Println("shift alter set auto increment success")
	} else {
		fmt.Println("shift alter auto increment failed")
		log.Fatal(err)
	}

/// shiftChange alter
	_, err = DB.Exec(`
	alter table shiftChange add primary key (initiatorShiftId, requestedShiftId);`)

	if err == nil {
		fmt.Println("shiftChange alter primary key success")
	} else {
		fmt.Println("shiftChange alter primary key failed")
		log.Fatal(err)
	}

	_, err = DB.Exec("alter table shiftChange add foreign key(initiatorShiftId) references shift(shiftId) on update cascade;")

	if err == nil {
		fmt.Println("shiftChange alter foreign key success")
	} else {
		fmt.Println("shiftChange alter foreign key failed")
		log.Fatal(err)
	}

	_, err = DB.Exec("alter table shiftChange add foreign key (requestedShiftId) references shift(shiftId) on update cascade;")

	if err == nil {
		fmt.Println("shiftChange alter foreign key success")
	} else {
		fmt.Println("shiftChange alter foreign key failed")
		log.Fatal(err)
	}
	defer DB.Close()
/// shiftOverTime alter
	_, err = DB.Exec(`
	alter table shiftOverTime add primary key (shiftId, initiatorOnOverTime, initiatorOffOverTime);`)

	if err == nil {
		fmt.Println("shiftOverTime alter primary key success")
	} else {
		fmt.Println("shiftOverTime alter primary key failed")
		log.Fatal(err)
	}

	_, err = DB.Exec("alter table shiftOverTime add foreign key (shiftId) references shift(shiftId) on update cascade;")

	if err == nil {
		fmt.Println("shiftOverTime alter foreign key success")
	} else {
		fmt.Println("shiftOverTime alter foreign key failed")
		log.Fatal(err)
	}
/// forgetPunch alter
	_, err = DB.Exec(`
	alter table forgetPunch add primary key (shiftId, targetPunch);`)

	if err == nil {
		fmt.Println("forgetPunch alter primary key success")
	} else {
		fmt.Println("forgetPunch alter primary key failed")
		log.Fatal(err)
	}

	_, err = DB.Exec(`alter table forgetPunch add foreign key (shiftId) references shift(shiftId) on update cascade;`)

	if err == nil {
		fmt.Println("forgetPunch alter foreign key success")
	} else {
		fmt.Println("forgetPunch alter foreign key failed")
		log.Fatal(err)
	}
/// dayOff alter
	_, err = DB.Exec(`alter table dayOff add primary key (shiftId, dayOffType);`)

	if err == nil {
		fmt.Println("dayOff alter primary key success")
	} else {
		fmt.Println("dayOff alter primary key failed")
		log.Fatal(err)
	}

	_, err = DB.Exec(`alter table dayOff add foreign key (shiftId) references shift(shiftId) on update cascade;`)

	if err == nil {
		fmt.Println("dayOff alter foreign key success")
	} else {
		fmt.Println("dayOff alter foreign key failed")
		log.Fatal(err)
	}
/// lateExecused alter
	_, err = DB.Exec(`alter table lateExcused add primary key (shiftId, lateExcusedType);`)

	if err == nil {
		fmt.Println("lateExcused alter primary key success")
	} else {
		fmt.Println("lateExcused alter primary key failed")
		log.Fatal(err)
	}

	_, err = DB.Exec(`alter table lateExcused add foreign key (shiftId) references shift(shiftId) on update cascade;`)

	if err == nil {
		fmt.Println("lateExcused alter foreign key success")
	} else {
		fmt.Println("lateExcused alter foreign key failed")
		log.Fatal(err)
	}
/// company alter
	_, err = DB.Exec(`alter table company add primary key (companyId, companyCode);`)

	if err == nil {
		fmt.Println("company alter primary key success")
	} else {
		fmt.Println("company alter primary key failed")
		log.Fatal(err)
	}

/// companyBanch alter
	_, err = DB.Exec(`alter table companyBanch add primary key (companyId, banchName);`)

	if err == nil {
		fmt.Println("companyBanch alter primary key success")
	} else {
		fmt.Println("companyBanch alter primary key failed")
		log.Fatal(err)
	}

	_, err = DB.Exec(`alter table companyBanch add foreign key (companyId) references company(companyId) on update cascade;`)

	if err == nil {
		fmt.Println("companyBanch alter foreign key success")
	} else {
		fmt.Println("companyBanch alter foreign key failed")
		log.Fatal(err)
	}
// user alter
	_, err = DB.Exec("alter table user auto_increment=9000034;")

	if err == nil {
		fmt.Println("user alter set auto increment success")
	} else {
		fmt.Println("user alter set auto increment failed")
		log.Fatal(err)
	}

	_, err = DB.Exec(`alter table user add foreign key (companyCode) references company(companyCode) on update cascade;`)

	if err == nil {
		fmt.Println("user alter foreign key success")
	} else {
		fmt.Println("user alter foreign key failed")
		log.Fatal(err)
	}

}
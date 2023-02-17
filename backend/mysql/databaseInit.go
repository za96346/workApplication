package mysql

import (
	"backend/mysql/table"
	"fmt"
	"strconv"
	"time"
)
func DataBaseInit() {
	simulateData();
	
}
func simulateData() {
	for x := 0; x < 2; x++ {
		addCompany(x)
		addUser(x)
	}
	addShift(0)
}
func addCompany(x int) {
	boss := 0
	if x == 0 {
		boss = 2
	} else {
		boss = 11
	}
		//company
		company := table.CompanyTable{
			CompanyCode: "company" + strconv.Itoa(x),
			CompanyName: "xx股份有限公司",
			CompanyLocation: "台中市大甲區ｘｘｘ",
			CompanyPhoneNumber: "0906930873",
			BossId: int64(boss),
			SettlementDate: 26,
			TermStart: time.Now(),
			TermEnd: time.Now(),
			CreateTime: time.Now(),
			LastModify: time.Now(),
		}
		_, _ = (*Singleton()).InsertCompany(&company)
		resData := (*Singleton()).SelectCompany(2, "company" + strconv.Itoa(x))
		for i := 0; i < 1; i++ {
		// company banch
			companyBanch := table.CompanyBanchTable{
				CompanyId :(*resData)[0].CompanyId,
				BanchName: "xx組" + strconv.Itoa(i),
				BanchShiftStyle: "{}",
				CreateTime: time.Now(),
				LastModify: time.Now(),
			}
			(*Singleton()).InsertCompanyBanch(&companyBanch)
			// for i := 0; i < 2; i++ {
			// 	banchStyle := table.BanchStyle{
			// 		BanchId: id,
			// 		OnShiftTime: "09:00",
			// 		OffShiftTime: "18:00",
			// 		Icon: ">'..'<",
			// 		TimeRangeName: "平日早班",
			// 		CreateTime: time.Now(),
			// 		LastModify: time.Now(),
			// 	}
			// 	(*Singleton()).InsertBanchStyle(&banchStyle)
	
			// 	banchRule := table.BanchRule{
			// 		BanchId: id,
			// 		MaxPeople: 2,
			// 		MinPeople: 1,
			// 		WeekDay: 1 + i,
			// 		WeekType: 2,
			// 		OnShiftTime: "09:00",
			// 		OffShiftTime: "18:00",
			// 		CreateTime: time.Now(),
			// 		LastModify: time.Now(),
			// 	}
			// 	(*Singleton()).InsertBanchRule(&banchRule)
			// }
		}
}

func addUser(x int) {
	for i := 0; i < 10; i++ {
			permession := 2
			if i == 1 {
				permession = 100
			}
			// user
			user := table.UserTable{
				CompanyCode: "company" + strconv.Itoa(x),
				Account: "account" + strconv.Itoa(i + x * 10),
				Password: "aa20010722",
				UserName: "siou" + strconv.Itoa(i),
				EmployeeNumber: "a0000" + strconv.Itoa(i),
				OnWorkDay: time.Now(),
				Banch: int64(x + 1),
				Permession: permession,
				MonthSalary: 30000 + i,
				PartTimeSalary: 130 + i,
				CreateTime: time.Now(),
				LastModify: time.Now(),
			}
			_, id := (*Singleton()).InsertUser(&user)
			// userPreference
			// resData := (*Singleton()).SelectUser(2, "account" + strconv.Itoa(i)) //拿使用者資料
			userPreference := table.UserPreferenceTable{
				UserId: id,
				Style: "{style}",
				FontSize: "12",
				SelfPhoto: "pic",
				CreateTime: time.Now(),
				LastModify: time.Now(),
			}
			_, _ = (*Singleton()).InsertUserPreference(&userPreference)
			// shift
			for shiftStep := 0; shiftStep <= 30; shiftStep++ {
				hours, _ :=  time.ParseDuration("1h")
				oneDay, _ := time.ParseDuration(fmt.Sprint(strconv.Itoa(shiftStep * 24), "h"))
				shift := table.ShiftTable{
					UserId: id,
					BanchStyleId: int64(1),
					OnShiftTime: time.Now().Add(-8 * hours).Add(oneDay),
					OffShiftTime: time.Now().Add(oneDay),
					RestTime: "01:00:00",
					SpecifyTag: "nothing",
					PunchIn: time.Now(),
					PunchOut: time.Now(),
					CreateTime: time.Now(),
					LastModify: time.Now(),
				}
				_, _ = (*Singleton()).InsertShift(&shift)
			}
	}
}
func addShift(x int) {
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
		(*Singleton()).InsertShiftChange(&shiftChange)
		(*Singleton()).InsertShiftOverTime(&shiftOverTime)
		(*Singleton()).InsertForgetPunch(&forgetPunch)
		(*Singleton()).InsertDayOff(&dayOff)
		(*Singleton()).InsertLateExcused(&lateExcused)
}
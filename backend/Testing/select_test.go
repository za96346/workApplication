package testing

import (
	"backend/handler"
	"backend/methods"
	"backend/mysql"
	"backend/redis"
	"backend/table"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectUser(t *testing.T) {
	compareFunc := func(v table.UserTable, vNext table.UserTable) bool {
		if v.UserId > vNext.UserId {
			return true
		}
		return false
	}
	handler.Init("../.env")

	// test 0
	r := (*redis.Singleton()).SelectUser(0)
	m := (*mysql.Singleton()).SelectUser(0)
	res := testEq(r, m, compareFunc)

	assert.Equal(t, res, true)

	// test 1
	r = (*redis.Singleton()).SelectUser(1, int64(1))
	m = (*mysql.Singleton()).SelectUser(1, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectUser(1, int64(1))
	m = (*mysql.Singleton()).SelectUser(1, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	//test 2
	r = (*redis.Singleton()).SelectUser(2, "account0")
	m = (*mysql.Singleton()).SelectUser(2, "account0")
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectUser(2, "account0")
	m = (*mysql.Singleton()).SelectUser(2, "account1")
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 3
	r = (*redis.Singleton()).SelectUser(3, "company1")
	m = (*mysql.Singleton()).SelectUser(3, "company1")
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectUser(3, "company0")
	m = (*mysql.Singleton()).SelectUser(3, "company100")
	// fmt.Println(r, m)
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)
}

func TestSelectUserPreference(t *testing.T) {
	compareFunc := func(v table.UserPreferenceTable, vNext table.UserPreferenceTable) bool {
		if v.UserId > vNext.UserId {
			return true
		}
		return false
	}

	// test 0
	r := (*redis.Singleton()).SelectUserPreference(0)
	m := (*mysql.Singleton()).SelectUserPreference(0)
	res := testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	// test 1
	r = (*redis.Singleton()).SelectUserPreference(1, int64(1))
	m = (*mysql.Singleton()).SelectUserPreference(1, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectUserPreference(1, int64(1))
	m = (*mysql.Singleton()).SelectUserPreference(1, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

}

func TestSelectCompany (t *testing.T) {
	compareFunc := func(v table.CompanyTable, vNext table.CompanyTable) bool {
		if v.CompanyId > vNext.CompanyId {
			return true
		}
		return false
	}

	// test 0
	r := (*redis.Singleton()).SelectCompany(0)
	m := (*mysql.Singleton()).SelectCompany(0)
	res := testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	// test 1
	r = (*redis.Singleton()).SelectCompany(1, int64(1))
	m = (*mysql.Singleton()).SelectCompany(1, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectCompany(1, int64(1))
	m = (*mysql.Singleton()).SelectCompany(1, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 2
	r = (*redis.Singleton()).SelectCompany(2, "company0")
	m = (*mysql.Singleton()).SelectCompany(2, "company0")
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectCompany(2, "company0")
	m = (*mysql.Singleton()).SelectCompany(2, "company100")
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

}

func TestSelectCompanyBanch(t *testing.T) {
	compareFunc := func(v table.CompanyBanchTable, vNext table.CompanyBanchTable) bool {
		if v.Id > vNext.Id {
			return true
		}
		return false
	}

	// test 0
	r := (*redis.Singleton()).SelectCompanyBanch(0)
	m := (*mysql.Singleton()).SelectCompanyBanch(0)
	res := testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	// test 1
	r = (*redis.Singleton()).SelectCompanyBanch(1, int64(1))
	m = (*mysql.Singleton()).SelectCompanyBanch(1, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectCompanyBanch(1, int64(1))
	m = (*mysql.Singleton()).SelectCompanyBanch(1, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 2
	r = (*redis.Singleton()).SelectCompanyBanch(2, int64(1))
	m = (*mysql.Singleton()).SelectCompanyBanch(2, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectCompanyBanch(2, int64(1))
	m = (*mysql.Singleton()).SelectCompanyBanch(2, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)
}

func TestSelectShift(t *testing.T) {
	compareFunc := func(v table.ShiftTable, vNext table.ShiftTable) bool {
		if v.ShiftId > vNext.ShiftId {
			return true
		}
		return false
	}

	// test 0
	r := (*redis.Singleton()).SelectShift(0)
	m := (*mysql.Singleton()).SelectShift(0)
	res := testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	// test 1
	r = (*redis.Singleton()).SelectShift(1, int64(1))
	m = (*mysql.Singleton()).SelectShift(1, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectShift(1, int64(1))
	m = (*mysql.Singleton()).SelectShift(1, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)
}

func TestSelectShiftChange(t *testing.T) {
	compareFunc := func(v table.ShiftChangeTable, vNext table.ShiftChangeTable) bool {
		if v.CaseId > vNext.CaseId {
			return true
		}
		return false
	}
	// handler.Init("../.env")
	// test 0
	r := (*redis.Singleton()).SelectShiftChange(0)
	m := (*mysql.Singleton()).SelectShiftChange(0)
	res := testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	// test 1
	r = (*redis.Singleton()).SelectShiftChange(1, int64(1))
	m = (*mysql.Singleton()).SelectShiftChange(1, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectShiftChange(1, int64(1))
	m = (*mysql.Singleton()).SelectShiftChange(1, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 2
	r = (*redis.Singleton()).SelectShiftChange(2, int64(1))
	m = (*mysql.Singleton()).SelectShiftChange(2, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectShiftChange(2, int64(1))
	m = (*mysql.Singleton()).SelectShiftChange(2, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 3
	r = (*redis.Singleton()).SelectShiftChange(3, int64(1))
	m = (*mysql.Singleton()).SelectShiftChange(3, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectShiftChange(3, int64(1))
	m = (*mysql.Singleton()).SelectShiftChange(3, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)
	
}

func TestSelectShiftOverTime(t *testing.T) {
	compareFunc := func(v table.ShiftOverTimeTable, vNext table.ShiftOverTimeTable) bool {
		if v.CaseId > vNext.CaseId {
			return true
		}
		return false
	}
	// handler.Init("../.env")
	// test 0
	r := (*redis.Singleton()).SelectShiftOverTime(0)
	m := (*mysql.Singleton()).SelectShiftOverTime(0)
	res := testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	// test 1
	r = (*redis.Singleton()).SelectShiftOverTime(1, int64(1))
	m = (*mysql.Singleton()).SelectShiftOverTime(1, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectShiftOverTime(1, int64(1))
	m = (*mysql.Singleton()).SelectShiftOverTime(1, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 2
	r = (*redis.Singleton()).SelectShiftOverTime(2, int64(1))
	m = (*mysql.Singleton()).SelectShiftOverTime(2, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)
	
	r = (*redis.Singleton()).SelectShiftOverTime(2, int64(1))
	m = (*mysql.Singleton()).SelectShiftOverTime(2, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)
	
}

func TestSelectForgetPunch(t *testing.T) {
	compareFunc := func(v table.ForgetPunchTable, vNext table.ForgetPunchTable) bool {
		if v.CaseId > vNext.CaseId {
			return true
		}
		return false
	}
	// handler.Init("../.env")
	// test 0
	r := (*redis.Singleton()).SelectForgetPunch(0)
	m := (*mysql.Singleton()).SelectForgetPunch(0)
	res := testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	// test 1
	r = (*redis.Singleton()).SelectForgetPunch(1, int64(1))
	m = (*mysql.Singleton()).SelectForgetPunch(1, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectForgetPunch(1, int64(1))
	m = (*mysql.Singleton()).SelectForgetPunch(1, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 2
	r = (*redis.Singleton()).SelectForgetPunch(2, int64(1))
	m = (*mysql.Singleton()).SelectForgetPunch(2, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)
	
	r = (*redis.Singleton()).SelectForgetPunch(2, int64(1))
	m = (*mysql.Singleton()).SelectForgetPunch(2, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)
	
}

func TestSelectLateExcused(t *testing.T) {
	compareFunc := func(v table.LateExcusedTable, vNext table.LateExcusedTable) bool {
		if v.CaseId > vNext.CaseId {
			return true
		}
		return false
	}
	// handler.Init("../.env")

	// test 0
	r := (*redis.Singleton()).SelectLateExcused(0)
	m := (*mysql.Singleton()).SelectLateExcused(0)
	res := testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	// test 1
	r = (*redis.Singleton()).SelectLateExcused(1, int64(1))
	m = (*mysql.Singleton()).SelectLateExcused(1, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectLateExcused(1, int64(1))
	m = (*mysql.Singleton()).SelectLateExcused(1, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 2
	r = (*redis.Singleton()).SelectLateExcused(2, int64(1))
	m = (*mysql.Singleton()).SelectLateExcused(2, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)
	
	r = (*redis.Singleton()).SelectLateExcused(2, int64(1))
	m = (*mysql.Singleton()).SelectLateExcused(2, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)
}

func TestSelectDayOff(t *testing.T) {
	compareFunc := func(v table.DayOffTable, vNext table.DayOffTable) bool {
		if v.CaseId > vNext.CaseId {
			return true
		}
		return false
	}
	// handler.Init("../.env")
	// test 0
	r := (*redis.Singleton()).SelectDayOff(0)
	m := (*mysql.Singleton()).SelectDayOff(0)
	res := testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	// test 1
	r = (*redis.Singleton()).SelectDayOff(1, int64(1))
	m = (*mysql.Singleton()).SelectDayOff(1, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectDayOff(1, int64(1))
	m = (*mysql.Singleton()).SelectDayOff(1, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 2
	r = (*redis.Singleton()).SelectDayOff(2, int64(1))
	m = (*mysql.Singleton()).SelectDayOff(2, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)
	
	r = (*redis.Singleton()).SelectDayOff(2, int64(1))
	m = (*mysql.Singleton()).SelectDayOff(2, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)
	
}

func TestSelectBanchStyle(t *testing.T) {
	compareFunc := func(v table.BanchStyle, vNext table.BanchStyle) bool {
		if v.StyleId > vNext.StyleId {
			return true
		}
		return false
	}

	// test 0
	r := (*redis.Singleton()).SelectBanchStyle(0)
	m := (*mysql.Singleton()).SelectBanchStyle(0)
	res := testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	// test 1
	r = (*redis.Singleton()).SelectBanchStyle(1, int64(1))
	m = (*mysql.Singleton()).SelectBanchStyle(1, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectBanchStyle(1, int64(1))
	m = (*mysql.Singleton()).SelectBanchStyle(1, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 2
	r = (*redis.Singleton()).SelectBanchStyle(2, int64(1))
	m = (*mysql.Singleton()).SelectBanchStyle(2, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectBanchStyle(2, int64(1))
	m = (*mysql.Singleton()).SelectBanchStyle(2, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)
	
}

func TestSelectBanchRule(t *testing.T) {
	compareFunc := func(v table.BanchRule, vNext table.BanchRule) bool {
		if v.RuleId > vNext.RuleId {
			return true
		}
		return false
	}

	// test 0
	r := (*redis.Singleton()).SelectBanchRule(0)
	m := (*mysql.Singleton()).SelectBanchRule(0)
	res := testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	// test 1
	r = (*redis.Singleton()).SelectBanchRule(1, int64(1))
	m = (*mysql.Singleton()).SelectBanchRule(1, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectBanchRule(1, int64(1))
	m = (*mysql.Singleton()).SelectBanchRule(1, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 2
	r = (*redis.Singleton()).SelectBanchRule(2, int64(1))
	m = (*mysql.Singleton()).SelectBanchRule(2, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectBanchRule(2, int64(1))
	m = (*mysql.Singleton()).SelectBanchRule(2, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)
	
}

func TestSelectQuitWorkUser(t *testing.T) {
	compareFunc := func(v table.QuitWorkUser, vNext table.QuitWorkUser) bool {
		if v.QuitId > vNext.QuitId {
			return true
		}
		return false
	}
	// handler.Init("../.env")

	// test 0
	r := (*redis.Singleton()).SelectQuitWorkUser(0)
	m := (*mysql.Singleton()).SelectQuitWorkUser(0)
	res := testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	// test 1
	r = (*redis.Singleton()).SelectQuitWorkUser(1, int64(1))
	m = (*mysql.Singleton()).SelectQuitWorkUser(1, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectQuitWorkUser(1, int64(1))
	m = (*mysql.Singleton()).SelectQuitWorkUser(1, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 2
	r = (*redis.Singleton()).SelectQuitWorkUser(2, int64(1))
	m = (*mysql.Singleton()).SelectQuitWorkUser(2, int64(1))
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectQuitWorkUser(2, int64(0))
	m = (*mysql.Singleton()).SelectQuitWorkUser(2, int64(2))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 3
	r = (*redis.Singleton()).SelectQuitWorkUser(3, "company0")
	m = (*mysql.Singleton()).SelectQuitWorkUser(3, "company0")
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectQuitWorkUser(3, "company1")
	m = (*mysql.Singleton()).SelectQuitWorkUser(3, "company0")
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)

	// test 4
	r = (*redis.Singleton()).SelectQuitWorkUser(4, "company0", 1)
	m = (*mysql.Singleton()).SelectQuitWorkUser(4, "company0", 1)
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)
	
	r = (*redis.Singleton()).SelectQuitWorkUser(4, "company0", int64(0))
	m = (*mysql.Singleton()).SelectQuitWorkUser(4, "company0", int64(100))
	res = testEq(r, m, compareFunc)
	assert.NotEqual(t, res, true)
}

func testEq[T comparable](a1 *[]T, b1 *[]T, compareFunc func(T, T) bool) bool {
	a := methods.BubbleSorted(a1, compareFunc)
	b := methods.BubbleSorted(b1, compareFunc)
    if len(*a) != len(*b) {
        return false
    }
    for i := range *a {
        if (*a)[i] != (*b)[i] {
            return false
        }
    }
    return true
}
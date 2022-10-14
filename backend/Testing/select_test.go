package testing

import (
	"backend/handler"
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
	handler.Init()

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
	r = (*redis.Singleton()).SelectUser(3, "fei32fej")
	m = (*mysql.Singleton()).SelectUser(3, "fei32fej")
	res = testEq(r, m, compareFunc)
	assert.Equal(t, res, true)

	r = (*redis.Singleton()).SelectUser(3, "fei32fej")
	m = (*mysql.Singleton()).SelectUser(3, "fmm")
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
	handler.Init()

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

func sorted[T comparable](arr *[]T, compareFunc func(T, T) bool) *[]T {
	a := *arr
	for oldStep := len(*arr) - 1;oldStep > 0; oldStep -- {
		for currentStep := 0; currentStep < oldStep; currentStep++ {
			if compareFunc(a[currentStep], a[currentStep + 1]) {
				b := a[currentStep]
				a[currentStep] = a[currentStep + 1]
				a[ currentStep + 1] = b
			}
		}
	}

	return &a
}

func testEq[T comparable](a1 *[]T, b1 *[]T, compareFunc func(T, T) bool) bool {
	a := sorted(a1, compareFunc)
	b := sorted(b1, compareFunc)
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
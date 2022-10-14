package methods

import (
	"fmt"
	"strconv"
)

func AnyToInt64(num any) (int64, any) {
	switch v := num.(type) {
	case int:
		return strconv.ParseInt(fmt.Sprint(v), 10, 64)
	case int64:
		return v, nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	default:
		return int64(-100), "hoa"
	}
}
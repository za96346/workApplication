package method

import "reflect"

/**
    @TODO 可以把 array param 換成是 pointer
*/
func InArray(array interface{}, val interface{}) (exists bool, index int) {
    exists = false
    index = -1

    switch reflect.TypeOf(array).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(array)

        for i := 0; i < s.Len(); i++ {
            if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
                index = i
                exists = true
                return
            }
        }
    }

    return
}


func ConvertSliceToInt(in []any) (out []int) {
    out = make([]int, 0, len(in))
    for _, v := range in {
        out = append(out, int(v.(float64)))
    }
    return
}
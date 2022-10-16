package methods

func IsNotExited[T any](value *[]T) bool {
	if len(*value) == 0 {
		return true
	}
	return false
}
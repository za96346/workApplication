package methods

func BubbleSorted[T comparable](arr *[]T, compareFunc func(T, T) bool) *[]T {
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
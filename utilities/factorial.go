package logisticsmaps

/*
Factorializer is a recursive function that will return the factorial of an integer */
func Factorializer(input int) int {

	if input > 1 && input < 10000 {
		return input * Factorializer(input-1)
	} else {
		return 1
	}
}

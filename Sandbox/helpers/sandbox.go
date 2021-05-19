package helpers

func PointToInc(number *int) {
	*number++
}

func StringTester() string {
	var test string = "Hello, Dave"

	return test[:2]
}

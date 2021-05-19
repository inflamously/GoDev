package helpers

import (
	"fmt"
	"testing"
)

func TestPointToInc(t *testing.T) {
	number := 1
	PointToInc(&number)

	fmt.Println(number)

	if number == 1 {
		t.Errorf("PointToInc failed due to invalid number")
	}
}

func TestStringTester(t *testing.T) {
	var str string = StringTester()
	if str == "" {
		t.Errorf("StringTester failed due to value empty.")
	}
}

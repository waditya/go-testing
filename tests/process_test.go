package main

import (
	"testing"

	"github.com/waditya/go-testing/process"
)

func TestCheckEven(testingPointer *testing.T) {

	input1 := 10
	expectedOutput1 := "YES"
	result1 := process.CheckEven(input1)

	if expectedOutput1 != result1 {
		testingPointer.Errorf("For : %v, Expected Value: %v, Got : %v", input1, expectedOutput1, result1)
	}

	input2 := 9
	expectedOutput2 := "NO"
	result2 := process.CheckEven(input2)

	if expectedOutput2 != result2 {
		testingPointer.Errorf("For : %v, Expected Value: %v, Got : %v", input2, expectedOutput2, result2)
	}
}

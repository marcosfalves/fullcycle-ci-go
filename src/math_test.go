package main

import "testing"

const RESULT_TEST_MESSAGE = "Result is invalid: Result %d. Expected: %d"

func TestSum(t *testing.T) {

	result := sum(15, 15)

	if result != 30 {
		t.Errorf(RESULT_TEST_MESSAGE, result, 30)
	}
}

func TestSub(t *testing.T) {

	result := sub(100, 15)

	if result != 85 {
		t.Errorf(RESULT_TEST_MESSAGE, result, 85)
	}
}

func TestTimes(t *testing.T) {

	result := times(10, 5)

	if result != 50 {
		t.Errorf(RESULT_TEST_MESSAGE, result, 50)
	}
}

func TestSumX(t *testing.T) {

	result := sumX(10, 5)

	if result != 25 {
		t.Errorf(RESULT_TEST_MESSAGE, result, 25)
	}
}

func TestSumY(t *testing.T) {

	result := sumY(10, 5)

	if result != 20 {
		t.Errorf(RESULT_TEST_MESSAGE, result, 20)
	}
}

func TestDiv(t *testing.T) {

	result := div(100, 2)

	if result != 50 {
		t.Errorf(RESULT_TEST_MESSAGE, result, 50)
	}
}

func TestAvg(t *testing.T) {

	result := avg(50, 75, 100)

	if result != 75 {
		t.Errorf(RESULT_TEST_MESSAGE, result, 75)
	}
}

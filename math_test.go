package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(15, 15)

	if result != 30 {
		t.Errorf("Result is invalid: Result %d. Expected: %d", result, 30)
	}
}

func TestSub(t *testing.T) {

	result := sub(100, 15)

	if result != 85 {
		t.Errorf("Result is invalid: Result %d. Expected: %d", result, 85)
	}
}

func TestTimes(t *testing.T) {

	result := times(10, 5)

	if result != 50 {
		t.Errorf("Result is invalid: Result %d. Expected: %d", result, 50)
	}
}

func TestSumX(t *testing.T) {

	result := sumX(10, 5)

	if result != 25 {
		t.Errorf("Result is invalid: Result %d. Expected: %d", result, 25)
	}
}

func TestSumY(t *testing.T) {

	result := sumY(10, 5)

	if result != 20 {
		t.Errorf("Result is invalid: Result %d. Expected: %d", result, 20)
	}
}

func TestDiv(t *testing.T) {

	result := div(100, 2)

	if result != 50 {
		t.Errorf("Result is invalid: Result %d. Expected: %d", result, 50)
	}
}

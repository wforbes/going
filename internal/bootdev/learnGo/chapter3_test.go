// taken from boot.dev Learn Go: Chapter 3 Lesson 3 ( https://www.boot.dev/lessons/a729ff01-7620-45a6-b330-7efb72bda67b )
package learnGo

import (
	"fmt"
	"testing"
)

func TestChapter3Lesson3(t *testing.T) {
	type testCase struct {
		tier     string
		expected int
	}

	testCases := []testCase{
		{"basic", 10000},
		{"premium", 15000},
		{"enterprise", 50000},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getMonthlyPrice(test.tier)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, test.tier, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, test.tier, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func TestChapter3Lesson7(t *testing.T) {
	type testCase struct {
		costPerSend  int
		numLastMonth int
		numThisMonth int
		expected     int
	}

	testCases := []testCase{
		{2, 89, 102, 26},
		{2, 98, 104, 12},
		{3, 50, 40, -30},
		{3, 60, 60, 0},
	}
	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := monthlyBillIncrease(test.costPerSend, test.numLastMonth, test.numThisMonth)
		_ = getBillForMonth(0, 0)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %v
Actual:     %v
Fail
`, test.costPerSend, test.numLastMonth, test.numThisMonth, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.costPerSend, test.numLastMonth, test.numThisMonth, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)

}

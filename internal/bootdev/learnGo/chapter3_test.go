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

func TestChapter3Lesson8(t *testing.T) {
	type testCase struct {
		tier     string
		expected string
	}
	testCases := []testCase{
		{"basic", "You get 1,000 texts per month for $30 per month."},
		{"premium", "You get 50,000 texts per month for $60 per month."},
		{"enterprise", "You get unlimited texts per month for $100 per month."},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getProductMessage(test.tier)
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

func TestChapter3Lesson9(t *testing.T) {
	type testCase struct {
		age                   int
		exYearsUntilAdult     int
		exYearsUntilDrinking  int
		exYearsUntilCarRental int
	}

	testCases := []testCase{
		{4, 14, 17, 21},
		{18, 0, 3, 7},
		{22, 0, 0, 3},
		{25, 0, 0, 0},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(test.age)
		if yearsUntilAdult != test.exYearsUntilAdult ||
			yearsUntilDrinking != test.exYearsUntilDrinking ||
			yearsUntilCarRental != test.exYearsUntilCarRental {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v, %v)
Actual:     (%v, %v, %v)
Fail
`, test.age, test.exYearsUntilAdult, test.exYearsUntilDrinking, test.exYearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v, %v)
Actual:     (%v, %v, %v)
Pass
`, test.age, test.exYearsUntilAdult, test.exYearsUntilDrinking, test.exYearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental)
		}
	}
	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func TestChapter3Lesson12(t *testing.T) {
	type testCase struct {
		age                   int
		exYearsUntilAdult     int
		exYearsUntilDrinking  int
		exYearsUntilCarRental int
	}

	testCases := []testCase{
		{4, 14, 17, 21},
		{18, 0, 3, 7},
		{22, 0, 0, 3},
		{25, 0, 0, 0},
		{35, 0, 0, 0},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental := yearsUntilEvents(test.age)
		if yearsUntilAdult != test.exYearsUntilAdult ||
			yearsUntilDrinking != test.exYearsUntilDrinking ||
			yearsUntilCarRental != test.exYearsUntilCarRental {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v, %v)
Actual:     (%v, %v, %v)
Fail
`, test.age, test.exYearsUntilAdult, test.exYearsUntilDrinking, test.exYearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  (%v, %v, %v)
Actual:     (%v, %v, %v)
Pass
`, test.age, test.exYearsUntilAdult, test.exYearsUntilDrinking, test.exYearsUntilCarRental, yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental)
		}
	}
	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func TestChapter3Lesson15(t *testing.T) {
	type testCase struct {
		message       string
		formatter     func(string) string
		formatterName string
		expected      string
	}

	testCases := []testCase{
		{"hello", addExclamation, "addExclamation", "TEXTIO: hello!!!"},
		{"hello there", addPeriod, "addPeriod", "TEXTIO: hello there..."},
		{"moor der ehT", reverseString, "reverseString", "TEXTIO: The red room"},
	}

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := reformat(test.message, test.formatter)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Fail
`, test.message, test.formatterName, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.message, test.formatterName, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
func addPeriod(s string) string {
	return s + "."
}
func addExclamation(s string) string {
	return s + "!"
}
func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

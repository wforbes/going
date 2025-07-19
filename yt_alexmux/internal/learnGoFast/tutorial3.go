// following along with youtube video
//
//	"Learn GO Fast: Full Tutorial"
//	https://www.youtube.com/watch?v=8uiZC0l4Ajw
//	This file includes content from 13:13 - 19:30
package learnGoFast

import (
	"errors"
	"fmt"
)

func RunT3() {
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v \n", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v \n", result, remainder)
	}

	// could also use a switch statement
	// Note! break statements are implied
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v \n", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v \n", result, remainder)
	}

	//and here's a conditional switch statement
	switch remainder {
	case 0:
		fmt.Printf("The division was exact\n")
	case 1, 2:
		fmt.Printf("The division was close\n")
	default:
		fmt.Printf("The division wasnt close\n")
	}
}

func printMe(printVal string) {
	fmt.Println(printVal)
}

//   - must set the return type in the function definition
//   - the return type can be multiple values
//   - adding error handling involves just having an early return that
//     returns empty values along with the error object that can be checked
func intDivision(numerator, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

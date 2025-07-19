// following along with youtube video
//
//	"Learn GO Fast: Full Tutorial"
//	https://www.youtube.com/watch?v=8uiZC0l4Ajw
//	This file includes content from 26:35 - 31:06
package learnGoFast

import (
	"fmt"
	"strings"
)

func RunT5() {
	demoStrings()
	demoRunes()
}

func demoStrings() {
	var myString = "résumé"
	fmt.Println(myString)
	//we can an index of the string as array
	var index0 = myString[0]
	fmt.Println(index0) // but its a number...
	// we can see its type with %T format
	fmt.Printf("%v, %T \n", index0, index0) // its a uint8
	// and we can iterate through the string to see all the array indicies
	for i, v := range myString {
		fmt.Println(i, v)
	} // notice it skips index 2 (the index after the é symbol)

	// Go represents strings with utf-8
	//	each char is represented with ascii (up to 128 codes or 7-bits)
	//	so uint8 goes 0-255 by default, making it a good choice
	//That's the number we're seeing with a normal english alphabet character
	fmt.Println(myString[0]) // <-- like this, it's 'r' so that's 114 ascii code
	// Each utf-8 encoded string is represented by a variable set of bytes
	//	'a' would only require 1 byte, but 🧠 emojis or き require 3 or 4 bytes
	var ki_hiragana = "き"
	var brain_emoji = "🧠"
	fmt.Printf("Length of き: %v \n", len(ki_hiragana))
	fmt.Printf("Length of 🧠: %v \n", len(brain_emoji))

	// so since é is code 233 (bigger than 128, less than 2047) it requires 2 bytes
	// 	it takes 2 memory ...spaces? hence we see iterating through the string "résumé"
	//	it goes from index 0, 1, ... 3, 4 - because index 2 is taken by the second half of é
	var fancy_e = "é"
	fmt.Printf("Length of é: %v \n", len(fancy_e))
	// and splitting these 2 bytes up we see the first half and second half of what's ultimately
	//	the full encoding for the é character
	fmt.Printf("Index 0 of \"é\": %v \n", fancy_e[0])
	fmt.Printf("Index 1 of \"é\": %v \n", fancy_e[1])
}

func demoRunes() {
	// an easier way to deal with iterating and indexing strings
	//	is to cast them as an array of runes
	var myString = []rune("résumé")
	var indexed = myString[1]
	// a rune returns 233 (the actual code for é)
	// 	instead of 195,169 in two parts
	fmt.Printf("%v, %T \n", indexed, indexed)
	// runes are unicode point numbers that represent the characters
	//	and they're just an alias for int32
	// when we iterate over this we get a set of continous
	//	indices and the values actually represent the code of each char
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// we can declare a rune type with single quotes
	var myRune = 'a'
	fmt.Printf("myRune = %v\n", myRune)

	// we can concatenate strings
	var strSlice = []string{"p", "o", "t", "a", "t", "o"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("%v \n", catStr)

	// strings are immutable in Go
	//	so we can't modify them once created
	// catStr[0] = 'a' <--- ERROR! cant assign to catStr[0]

	// Since they're immutable, using += to concatenate
	//	the string creates a new string with each loop iteration
	//	which is really inefficient. Instead we can use the Builder()
	//	from Go's "strings" package
	var strSlice2 = []string{"p", "o", "t", "a", "t", "o"}
	var strBuilder strings.Builder
	var catStr2 = ""
	for i := range strSlice2 {
		// an array is allocated and appended to in the Builder
		strBuilder.WriteString(strSlice2[i])
	}
	catStr2 = strBuilder.String() // a new string is made only at the end with .String()
	fmt.Printf("%v \n", catStr2)

}

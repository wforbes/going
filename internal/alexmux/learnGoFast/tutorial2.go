package learnGoFast

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func RunT2() {
	var intNum int = 9223372036854775807 // 'int' defaults to the CPU architecture. Mines 64-bit CPU so this is a int64
	// 9.22 billion ^ is the largest int64 value
	fmt.Println("Max default int size:", intNum)
	// also have: int8, int16, int32, int64

	var int16Num int16 = 32767                                                     // largest 16 bit number
	int16Num = int16Num + 1                                                        // doesn't throw runtime overflow error
	fmt.Println("Overflowing int16 gets you the lowest negative int16:", int16Num) // this prints -327678

	// uint types can't be negative but give you twice the size in the positive direction
	var uintNum uint = 18446744073709551615 // largest unsigned uint64 number
	fmt.Println("Max default unsigned int size:", uintNum)
	// also have: uint8, uint16, uint32, uint64

	var float32Num float32 = 12345678.9 // float has no default, must use float32 or float64
	fmt.Println(float32Num)             // prints how this is stored in your machine
	//	^ Windows prints '1.2345679e+07
	var float64Num float64 = 12345678.9
	fmt.Println(float64Num)
	//	^ Windows prints '1.23456789e+07 (more precision than float32)

	// data type size is useful because if you wanted to store
	//	a 256 RGB value for color, a uint8 would be the best fit
	//	and conserve more memory than just using int
	var red int = 255
	var blue int16 = 255
	var green uint8 = 255
	fmt.Println("Size of Red var:", unsafe.Sizeof(red), "bytes")     // 8 bytes
	fmt.Println("Size of Blue var:", unsafe.Sizeof(blue), "bytes")   // 2 bytes
	fmt.Println("Size of Green var:", unsafe.Sizeof(green), "bytes") // 1 byte
	// Since the max value of uint8 is 255 and RGB values go up to 255
	//	there's no reason to use larger ints for it

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// can't perform arithmetic on mixed types
	//	you've got to cast the data types to match
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(floatNum32, "+", intNum32, "=", result)

	// VIDEO TIMESTAMP 9:35 ~

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) //integer division rounds down
	fmt.Println(intNum1 % intNum2) //modulous operator works as expected to get remainder

	var myString string = "Strings: You can use double quotes, but you can't line wrap them without the \\n newline char"
	fmt.Println(myString)
	var myString2 string = `Strings: You can use the backtick symbol
to make it a "raw string literal" and it will
conform to any whitespace within the bounds of the backticks`
	fmt.Println(myString2)
	var myString3 string = "Strings: You" + " " + " can concatinate them " + "with the + operator!"
	fmt.Println(myString3)

	// you can get the length of a string with len(s) but this can be decieving
	//	because its the number of bytes... not the number of characters in the str
	fmt.Println("Length of the string \"test\": ", len("test")) // returns 4
	// GO uses utf-8 encoding, chars outside of the normal 256 ascii charset are stored in
	//	more than a single byte
	fmt.Println("Length of thinking-face emoji: ", len("🤔"))  // returns 4
	fmt.Println("Length of double-daggar symbol: ", len("‡")) // returns 3
	// So to get the true length we can import the "unicode/utf8" package and work with Runes
	fmt.Println("Rune count of double-daggar symbol: ", utf8.RuneCountInString("‡")) // returns 1

	//A rune is an alias for int32 and represents a Unicode code point.
	// Since emojis are part of the Unicode standard, they can be directly assigned to rune variables.
	var myRune rune = '😏'
	fmt.Printf("For emojis we can format strings to output them like this: %c \n", myRune)

	var myBoolean bool = false
	fmt.Println("Boolean types: ", myBoolean)

	// Go sets uninitialized primatives as a default value
	//	Strings is ""
	//	All numbers is 0
	//	Booleans is false

	var myVar = "This type is inferred"
	fmt.Println("Declaring a variable without a type allows type inference: ", myVar)
	myVar2 := "This type is inferred too"
	fmt.Println("We can even drop the 'var' and use the := walrus/shorthand symbol to declare an inferred variable: ", myVar2)
	var1, var2 := 1, 2
	fmt.Printf("We can declare multiple inferred type vars of the same type at the same with commas... %v %v\n", var1, var2)

	const myConst string = "const value"
	fmt.Println("Constants cant be changed, must be declared with a value: ", myConst)

	// VIDEO TIMESTAMP 13:13 ~
	// Continued in tutorial_3 folder of this project
}

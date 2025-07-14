// following along with youtube video
//
//	"Learn GO Fast: Full Tutorial"
//	https://www.youtube.com/watch?v=8uiZC0l4Ajw
package main

import (
	"fmt"
	"unsafe"
)

func main() {
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
}

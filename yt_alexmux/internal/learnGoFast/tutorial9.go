// following along with youtube video
//
//	"Learn GO Fast: Full Tutorial"
//	https://www.youtube.com/watch?v=8uiZC0l4Ajw
//	This file includes content from 47:10 - 52:50
package learnGoFast

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 8

// Channels are a way for goroutines to pass around information
//   - they hold data, like an int or a slice
//   - they're thread safe, we avoid data races when reading and writing from memory
//   - we can listen for when data is added or removed from a channel and we can
//     block code execution until one of these events happen
func RunT9() {
	demoSimpleChannel()
	demoLoopedChannel()
	demoBufferedChannel()
	channelPractice()
}

func demoSimpleChannel() {
	var c = make(chan int) // make a channel that can hold an int

	// if we work with it directly here, it will dead-lock.
	// when we write to an unbuffered channel, the code will block until something else
	// reads from it... see below,
	// the line after writing to the channel will never get hit.
	// c <- 1 // this is how we write the int value 1 to the channel
	// ^ execution stops here, waiting for something else to read from the channel
	//	before continuing
	// var i = <- c // this is how we can pop the value out of the channel into the variable
	// fmt.Println(i)

	go process1(c)   // instead, we do the operation in a go routine concurrently
	var number = <-c // so that execution can reach this line, popping
	fmt.Printf("%d was popped from the channel, into 'number' var \n", number)
}

func process1(c chan int) {
	c <- 123 // we write to the channel
	fmt.Println("123 was written to the channel")
}

// Looping on a channel -
//
//	allows for waiting on a channel to get a value then
//	popping values out of a channel as they're available
//	and looping to wait for the next value
//	calling close() on the channel signals the loop to stop
//	waiting for more values
func demoLoopedChannel() {
	var c2 = make(chan int)
	go process2(c2)
	for i := range c2 { // waiting for something to be pushed to the channel
		// ^ we pop a value out of the channel into 'i' var (similar to i<-c)
		fmt.Println(i) // then we print out the value
	}
}

func process2(c chan int) {
	defer close(c) // without closing, we get a dead-lock
	// ^ so using defer here ensure that happens when this
	//	function ends execution
	//close(c) notifies any other process that we're done with the channel
	//	so it'll stop waiting for any more values from it
	for i := 0; i < 5; i++ {
		c <- i // we push a value into the channel and wait for something to read it
		// when its read from the channel elsewhere, execution continues here
		//	and iterates another loop
	}
	fmt.Println("Exiting process2")
}

// Buffered Channels
//
//	allows for storing multiple values in a channel
//	this is similar to the previous example except for one big difference
//	process3 can write everything to the channel on its own and be done
//	it doesnt have to hang around waiting for something to read from it
//	(notice the "Exiting process" message prints before any values are printed)
func demoBufferedChannel() {
	var c = make(chan int, 5)
	go process3(c)
	for i := range c { // waiting for something to be pushed to the channel
		// but since all the values are pushed to the channel, now it can
		// iterate through them on its own schedule
		time.Sleep(time.Second * 1) // <- to demonstrate, we sleep for a second on each
		fmt.Println(i)              // then we print out the value
	}
}

func process3(c chan int) {
	defer close(c) // without closing, we get a dead-lock
	for i := 5; i <= 10; i++ {
		c <- i
	}
	fmt.Println("Exiting process3")
}

// Channel Practice - listening for a good deal on food
//
//	this is an example of listening to two channels and doing something
//	some point in the future when the channel gets a value
//	we run two goroutines that might push a string to one of two channels in the future
//	we listen to those channels in a select statement that will notify us that a deal was found
func channelPractice() {

	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		fmt.Printf("new price on chicken! ($%.2f)\n", chickenPrice)
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		fmt.Printf("new price on tofu! ($%.2f)\n", tofuPrice)
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case chickenSite := <-chickenChannel:
		fmt.Printf("Text: Found a deal on chicken at %s", chickenSite)
	case tofuSite := <-tofuChannel:
		fmt.Printf("Email: Found a deal on tofu at %s", tofuSite)
	}
}

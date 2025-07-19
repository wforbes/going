// following along with youtube video
//
//	"Learn GO Fast: Full Tutorial"
//	https://www.youtube.com/watch?v=8uiZC0l4Ajw
//	This file includes content from 47:10 - 52:50
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 8

func main() {
	demoSimpleChannel()
	demoLoopedChannel()
	demoBufferedChannel()
	channelPractice()
}

func demoSimpleChannel() {
	var c = make(chan int) // make a channel that can hold an int
	/* This will dead-lock. when we write to an unbuffered channel, the code will block until something else
	reads from it... so the line after writing to the channel below will never get hit.
	c <- 1 // add the value '1' to the channel
	var i = <- c // pop the value out of the channel, into the variable
	//	now the channel is empty
	fmt.Println(i)
	*/
	go process1(c)
	fmt.Println(<-c)
}

func process1(c chan int) {
	c <- 123
}

func demoLoopedChannel() {
	var c2 = make(chan int)
	go process2(c2)
	for i := range c2 {
		fmt.Println(i)
	}
}

func process2(c chan int) {
	defer close(c) // without closing, we get a dead-lock
	for i := 0; i < 5; i++ {
		c <- i
	}
}

func demoBufferedChannel() {
	var c = make(chan int, 5)
	go process3(c)
	for i := range c {
		time.Sleep(time.Second * 1)
		fmt.Println(i)
	}
}

func process3(c chan int) {
	defer close(c) // without closing, we get a dead-lock
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Exiting process")
}

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

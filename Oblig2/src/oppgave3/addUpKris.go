package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
	"sync"
)

var wg = sync.WaitGroup{}
var wg2 = sync.WaitGroup{}
var c1 = make(chan int, 1)
var c2 = make(chan int, 1)
var c3 = make(chan int, 1)

func main() {
wg2.Add(1)
	go func() {

		inputOne := os.Args[1]
		inputTwo := os.Args[2]

		intOne, err := strconv.Atoi(inputOne)
		intTwo, err := strconv.Atoi(inputTwo)

		if err != nil {
			log.Fatal(err)
		}

		c1 <- intOne
		c2 <- intTwo
		wg.Add(1)
		go addUp()
		wg.Wait()
		sum := <-c3
		fmt.Println(sum)
	wg2.Done()
	}()
	wg2.Wait()
	}


 func addUp() {
	add1 := <-c1
	add2 := <-c2
	sum := add1 + add2
	c3 <- sum
	wg.Done()

}

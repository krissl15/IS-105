package main

import (
	"fmt"
	"os"
	"strconv"
	"log"
	"sync"
	"os/signal"
	"syscall"
	"time"
)

var wg = sync.WaitGroup{} //oppretter 2 waitgroups, forklarer dem underveis
var wg2 = sync.WaitGroup{}
var c1 = make(chan int, 1) // c1 skal lagre et tall
var c2 = make(chan int, 1)// c2 lagrer et annet tall
var c3 = make(chan int, 1) //c3 lagrer summen

func main() {
	sigAdd() //goroutine som tar opp SIGINT-signal.

	wg2.Add(1) //wg2 vil vente i main funksjonen til delta=0
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
		wg.Add(1) //wg blir done når den har puttet summen i c3 på linke 56
		go addUp()
		wg.Wait() // venter på at summen skal bli puttet i c3
		sum := <-c3
		time.Sleep(900 * time.Millisecond) //så en rekker ctrl+C...
		fmt.Println(sum)
		wg2.Done() //wg2 delta = 0 når denne gofunksjonen er ferdig
	}()
	wg2.Wait() //wg2 venter på at den skal bli done, slik at ikke main-funksjonen avslutter før goroutinen er ferdig.
}


func addUp() {
	add1 := <-c1
	add2 := <-c2
	sum := add1 + add2
	c3 <- sum
	wg.Done() // wg venter på linje 40, den venter på at sum skal bli puttet i c3.
}

func sigAdd() { // gjør hele goroutinen til egen funksjon. Ryddigere.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		fmt.Printf("Motatt SIGINT signal før fullførelse")
		os.Exit(1)
	}()
}

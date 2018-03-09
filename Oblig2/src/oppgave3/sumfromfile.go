package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"os/signal"
	"syscall"
	"strconv"
	"strings"
)

func main() {
	sigz() //goroutine som tar opp SIGINT.

	file, err := ioutil.ReadFile("numbers.txt.lock")
	checke(err)

		sum := sumBytes(file)
		ioutil.WriteFile("numbers.txt.lock", sum, 0777) // skriv summen inn i tekstfilen
		fmt.Printf("%s", sum)
	}

func sumBytes(b []byte) []byte {
	strSlice := string(b)                   //gjør teksten i filen "numbers" til string
	numbers := strings.Split(strSlice, " ") //tallene i "numbers.txt" har mellomrom mellom seg.
	nrOne := numbers[0]                     //hadde jeg ikke splittet, men brukt "[]"-plassene, ville dette skapt
	nrTwo := numbers[1]                     //problem om tallene har flere siffer, da hvert siffer lagres.

	intOne, err := strconv.Atoi(nrOne) //convert tallene til ints
	intTwo, err := strconv.Atoi(nrTwo)

	checke(err)

	sum := intOne + intTwo //summer tallene
	s := strconv.Itoa(sum) //gjør summen til string
	bySum := []byte(s)     //gjør stringen til []byte
	return bySum
}

func checke(err error) { //så jeg slipper den gad damn if-setninga overalt
	if err != nil {
		log.Println(err)
		return
	}
}

func sigz() { // gjør hele goroutinen til egen funksjon. Ryddigere.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		fmt.Printf("Motatt SIGINT signal før fullførelse")
		os.Exit(1)
	}()

}

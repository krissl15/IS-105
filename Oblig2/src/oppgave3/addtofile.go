package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"os/signal"
	"syscall"

	"time"
)

func main() {

	sig() // Kjør goroutine som tar opp SIGINT-signal



	if len(os.Args) == 3 { //sjekker om det er 3 argument når programmet kjøres ([0]filnavn, [1]tall 1, [2]tall2)
		time.Sleep(900 * time.Millisecond)
		nrOne := os.Args[1]      //os.args blir en slice, plass [1] blir første tallet som skrives (etter filnavn)
		nrTwo := os.Args[2]      // [2] blir tall 2
		createFile(nrOne, nrTwo) // kjør createfile med tallene skrevet i cmd som parametre

		fmt.Println("created file: 'numbers.txt.lock' with numbers:", nrOne,"and", nrTwo)
	} else if len(os.Args) == 1 { //om det ikke er skrevet noen tall når programmet kjøres, vil det som står i tekstfilen printes til kommando
		file, err := ioutil.ReadFile("numbers.txt.lock") //åpner filen for lesing, variabelen "file" blir []byte
		fileText := string(file)                    //gjør om file (som er []byte) til string
		check(err)
		time.Sleep(900 * time.Millisecond)
		fmt.Printf("%q", fileText)
	}else{ // Programmet kjører kun om det er 0 tall eller 2 tall som parametre. alt annet er feil.
		fmt.Println("Oops, you did something wrong, this code is perfect tho!")
		fmt.Println("If you havent already, write 'addtofile' + number1 + number2 ")
		fmt.Println("If you have created a file with numbers, then write 'sumfromfile'")
		fmt.Println("If you have already summarized, make sure you only write 'addtofile' to print sum")
	}

}

func createFile(one, two string) {
	f, err := os.Create("numbers.txt.lock") //lag en fil som heter "numbers.txt"
	check(err)
	defer f.Close()
	nums := []byte(one + " " + two) //lag en []byte med parametrene
	f.Write(nums)                   //skriver parametrene inn i "numbers.txt", bruker mellomrom for string.Split.
}

func check(err error) { //så jeg slipper den gad damn if-setninga overalt
	if err != nil {
		log.Println(err)
		return
	}
}

func sig() { // gjør hele goroutinen til egen funksjon. Ryddigere.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		fmt.Printf("Motatt SIGINT signal før fullførelse")
		os.Exit(1)
	}()

}

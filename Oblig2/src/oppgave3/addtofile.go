package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
	"os/signal"
	"syscall"
	"os/exec"
)

func main() {

	sig() // Kjør goroutine som tar opp SIGINT-signal

	if len(os.Args) == 3 { //sjekker om det er 3 argument når programmet kjøres ([0]filnavn, [1]tall 1, [2]tall2)

		nrOne := os.Args[1]      //os.args blir en slice, plass [1] blir første tallet som skrives (etter filnavn)
		nrTwo := os.Args[2]      // [2] blir tall 2
		createFile(nrOne, nrTwo) // kjør createfile med tallene skrevet i cmd som parametre

		runSumFromFile()

		file, err := ioutil.ReadFile("numbers.txt.lock")
		check(err)
		fmt.Println(string(file))

	} else { // Programmet kjører kun om det er 0 tall eller 2 tall som parametre. alt annet er feil.
		fmt.Println("Oops, you did something wrong, this code is perfect tho!")
		fmt.Println("Make sure you write like this: addtofile number number")
	}

}

func createFile(one, two string) {
	f, err := os.Create("numbers.txt.lock") //lag en fil som heter "numbers.txt"
	check(err)
	defer f.Close()
	nums := []byte(one + " " + two) //lag en []byte med parametrene
	f.Write(nums)                   //skriver parametrene inn i "numbers.txt", bruker mellomrom for string.Split.
}

func runSumFromFile() { //Kjør programmet sumfromfile.exe
	path, err := exec.LookPath("sumfromfile.exe")
	check(err)
	cmd := exec.Command(path)
	cmd.Run()

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

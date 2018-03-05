package main

import (
	"os"
	"log"
	"io/ioutil"
	"fmt"
)



func main() {


    file, err :=  ioutil.ReadFile("numbers.txt") //åpner filen for lesing, variabelen "file" blir []byte
    fileText := string(file) //gjør om file (som er []byte) til string
	check(err)//quickie

    if len(os.Args) == 3 { //sjekker om det er 3 argument når programmet kjøres (filnavn, tall 1, tall2)

		nrOne := os.Args[1] //os.args blir en slice, plass [1] blir første tallet som skrives (etter filnavn)
		nrTwo := os.Args[2]// [2] blir tall 2
		createFile(nrOne, nrTwo) // kjør createfile med tallene skrevet i cmd som parametre
	}else{ //om det ikke er skrevet noen tall når programmet kjøres, vil det som står i tekstfilen printes til kommando
		fmt.Printf("%q",fileText)

	}

}

func createFile(one, two string) {
	f, err := os.Create("numbers.txt") //lag en fil som heter "numbers.txt"
	check(err)
	defer f.Close()
	nums := []byte(one + " " + two) //lag en []byte med parametrene
	f.Write(nums) //skriver parametrene inn i "numbers.txt"

}

func check(err error){ //så jeg slipper den gad damn if-setninga overalt
	if err != nil {
		log.Println(err)
		return
	}


}



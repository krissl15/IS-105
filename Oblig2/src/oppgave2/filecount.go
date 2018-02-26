package main

import "fmt"

func main() {
	fmt.Println("Programmet skal lese en tekst-fil, skrive ut totalt antall linjer og antall for fem runes som forekommer hyppigst i filen")
}
package Oblig2

import (
	"os"
	"log"
	"bufio"
	"fmt"
)

var (
	fileInfo os.FileInfo
	err      error
)

func main() {
	// Stat returns file info. It will return
	// an error if there is no file.
	fileInfo, err = os.Stat("Textcounter.rtf")
	if err != nil {
		log.Fatal(err)


file, _ := os.Open("/path/to/desktop/Textcounter")
fileScanner := bufio.NewScanner(file)
lineCount := 0
for fileScanner.Scan() {
lineCount++
}
fmt.Println("number of lines:", lineCount)

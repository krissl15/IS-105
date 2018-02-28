package main

import (
	"fmt"
	"os"
	"bytes"
	"log"
)

func main() {
	filerc, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer filerc.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	contents := buf.String()
	runeCont := []rune(contents)

	for i := 0; i < len(runeCont); i++ {

	}




	// iterate over the slice of words and populate
	// the map with word:frequency pairs
	// (where word is the key and frequency is the value)
	word_freq := make(map[rune]int)
	// range string slice gives index, word pairs
	// index is not needed, so use blank identifier _
	for _, word := range runeCont {
		// check if word (the key) is already in the map
		_, ok := word_freq[word]
		// if true add 1 to frequency (value of map)
		// else start frequency at 1
		if ok == true {
			word_freq[word] += 1
		} else {
			word_freq[word] = 1
		}
	}

	for key, value := range word_freq {
		fmt.Println("rune:", key, "antall:", value)
	}


}







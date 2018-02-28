package main

import (
	"fmt"
	"os"
	"bytes"
	"log"
	"sort"
)

func main() {
	filerc, err := os.Open("text.txt") //Hent filen her, linjene under er dritkul kode som gjør txt til string.
	if err != nil {
		log.Fatal(err)
	}
	defer filerc.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	contents := buf.String()
	runeCont := []rune(contents) //OG TIL RUNE WÆÆÆÆÆÆÆÆÆ

	word_freq := make(map[rune]int)//Hallgeir hadde digga dette

	for _, word := range runeCont { //legger til runer som value, med antall ganger brukt som value SHAZAM
		_, ok := word_freq[word]
		if ok == true {
			word_freq[word] += 1
		} else {
			word_freq[word] = 1
		}
	}

	type kv struct {//aner ikke hvordan jeg sorterer maps, tror de er inconsise, mekker pairs og sorterer dem istedet.
		Key   rune
		Value int
	}

	var ss []kv
	for k, v := range word_freq {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	fmt.Println("Information about", filerc.Name())
	fmt.Println("number of lines", word_freq[10]+1) // Finn element i mappet med key 10 (10 er unicode til newline)
	fmt.Println("Most common runes:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%v %d %v %d\n", "rune:", ss[i].Key, "antall:", ss[i].Value) //output: hvilken rune og hvor mange ganger anvendt
	}
}

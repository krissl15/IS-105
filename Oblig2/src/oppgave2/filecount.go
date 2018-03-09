package main

import (
	"fmt"
	"os"
	"bytes"
	"log"
	"sort"

)

/**
*Koden leser følgende info om en fil: Hvor mange linjer det er i filen,
*hvilke fem runer som blir mest brukt, og hvor mange ganger de anvendes.
*
*@param filnavn på textfilen som skal anvendes
 */
func filecount(s string) {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := new(bytes.Buffer) //linjene under er dritkul kode som gjør txt til string.
	buf.ReadFrom(file)
	contents := buf.String()
	runeCont := []rune(contents) //OG TIL RUNE

	//lag et map som tar individuelle runer som key, antall ganger brukt som value
	runeMap := make(map[rune]int)

	for _, rune := range runeCont {
		_, ok := runeMap[rune]
		if ok == true {
			runeMap[rune] += 1
		} else {
			runeMap[rune] = 1
		}
	}
	//Maps er ikke addresserbare, så oppretter struct og sorterer dem istedet.
	type ri struct {
		Key   rune
		Value int
	}
	var ss []ri
	for k, v := range runeMap {
		ss = append(ss, ri{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	fmt.Println("")
	fmt.Println("Information about", file.Name())
	fmt.Println("")
	fmt.Println("number of lines", runeMap[10]) // Finn element i mappet med key 10 (10 er unicode til newline).
	fmt.Println("")
	fmt.Println("Most common runes:")
	for i := 0; i < 5; i++ {
		fmt.Printf("%v %v %d %v %c %v %v %d\n", i+1, "Rune:", ss[i].Key, "(", ss[i].Key, ")", "Counts:", ss[i].Value) //output: hvilken rune og hvor mange ganger anvendt
	}
}

func main() {

	filename := os.Args[1] // Må nå ha et argument for å kjøre i cmd
	filecount(filename) //Kjør filecount med argumentet som parameter
}

package main

import (

	"io/ioutil"
	"log"
	"strconv"
	"os"
	"strings"



)

func main()  {
	content, err := ioutil.ReadFile("numbers.txt") // åpner for å bruke ioutil
	file, err := os.OpenFile("numbers.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666, )//åpner for å bruke os.Write.

	if err != nil {
		log.Fatal(err)
	}

	strSlice := string(content) //gjør teksten i filen "numbers" til string
	numbers := strings.Split(strSlice, " ") //tallene i "numbers.txt" har mellomrom mellom seg.
	nrOne := numbers[0]							//hadde jeg ikke splittet, men brukt []posisjonene, ville dette skapt
	nrTwo := numbers[1]							//problem om tallene har flere siffer

	intOne, err := strconv.Atoi(nrOne) //convert tallene til ints
	intTwo, err := strconv.Atoi(nrTwo)

   sum := intOne+intTwo //summer tallene
	s := strconv.Itoa(sum) //gjør summen til string
	b := []byte(s)//gjør stringen til []byte

	file.Write(b) //skriv inn []byten til filen







}

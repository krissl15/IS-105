package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	filename := os.Args[1] // Må nå ha et argument for å kjøre i cmd
	fileInfo(filename)
}

/**
*Funksjon som sjekker detaljert informasjon om en fil
*@param filnavn på fil som skal testes
 */
func fileInfo(s string) {
	file, err := os.Lstat(s)

	bytes := float64(file.Size())//gjør filstørrelsen om til float64 for å kunne regne sammen KB MB og GB
	KB := bytes / 1024
	MB := KB / 1024
	GB := MB / 1024

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("##########################################################")
	fmt.Println("Information about", file.Name())
	fmt.Println("")
	fmt.Printf("%v, %f, %v, %f, %v %f %v %.15g %v\n ", "Size :", bytes, "bytes,", KB, "KB,", MB, "MB,", GB, "GB") //fant ingen bedre løsning
	fmt.Println("")

	//If-else setninger som sjekker informasjon om filen. Hva de forskjellige setningene gjør er selvforklarende.
	if file.Mode().IsDir() == true {
		fmt.Println("Is a directory")
	} else if file.Mode().IsDir() == false {
		fmt.Println("Is not a directory")
	}

	if file.Mode().IsRegular() { //
		fmt.Println("Is a regular file")
	} else {
		fmt.Println("Is not a regular file")
	}

	fmt.Println("Has Unix permission bits:", file.Mode().Perm())

	if file.Mode()&os.ModeAppend == os.ModeAppend {
		fmt.Println("Is append only")
	} else {
		fmt.Println("Is not append only")
	}

	if file.Mode()&os.ModeDevice == os.ModeDevice {
		fmt.Println("Is a device file")
	} else {
		fmt.Println("Is a device file")
	}

	if file.Mode()&os.ModeCharDevice == os.ModeCharDevice {
		fmt.Println("Is a Unix character device")
	} else {
		fmt.Println("Is not Unix character device")
	}

	if file.Mode()&os.ModeSocket == os.ModeSocket {
		fmt.Println("Is a unix block device")
	} else {
		fmt.Println("Is not a unix block device")
	}

	if file.Mode()&os.ModeSymlink == os.ModeSymlink {
		fmt.Println("Is a symbolic link")
	} else {
		fmt.Println("Is not a symbolic link")
	}

}

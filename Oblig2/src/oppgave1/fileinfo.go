package main

import ("fmt"
	"os"
	"log"
)


func main() {
	filename := os.Args[1] // Må nå ha et argument for å kjøre i cmd
	fileInfo(filename)

}

func fileInfo(s string) {
	file, err := os.Lstat(s)



	bytes := float64(file.Size())
	KB := bytes / 1024
	MB := KB / 1024
	GB := MB / 1024


	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Information about", file.Name())
	fmt.Printf("%v, %f, %v, %f, %v %f %v %.15g %v\n ", "Size :", bytes, "bytes,", KB, "KB,", MB, "MB,", GB, "GB")

	if file.Mode().IsDir() == true {
		fmt.Println("Is a directory")
	} else if file.Mode().IsDir() == false {
		fmt.Println("Is not a directory")
	}

	if file.Mode().IsRegular() { //
		fmt.Println("Is a regular file") // hvis det er en file, bli det true
	} else {
		fmt.Println("Is not a regular file") // hvis det er bare en mappe bli det false
	}

	fmt.Println("Has Unix permission bits:", file.Mode().Perm()) // Hvilke retigheter man har på filen

	if file.Mode()&os.ModeAppend == os.ModeAppend {          // Deler filen data/info med andre filer
		fmt.Println("Is append only")
	} else {
		fmt.Println("Is not append only")
	}

	if file.Mode()&os.ModeDevice == os.ModeDevice { // En fil som lar software kommunisere med "device driver" ved hjelp av standar input/output system calls
		fmt.Println("Is a device file")
	} else {
		fmt.Println("Is a device file")
	}

	if file.Mode()&os.ModeCharDevice == os.ModeCharDevice {
		fmt.Println("Is a Unix character device")
	} else {
		fmt.Println("Is not Unix character device")
	}

	if file.Mode()&os.ModeSocket == os.ModeSocket{
		fmt.Println("Is a unix block device")
	}else {
		fmt.Println("Is not a unix block device")
	}

	if file.Mode()&os.ModeSymlink == os.ModeSymlink { // Det er en symlink
		fmt.Println("Is a symbolic link")
	} else {
		fmt.Println("Is not a symbolic link")
	}

}
/**
Is/Is not a directory

Is/Is not a regular file

Has Unix permission bits: -rwxrwxrwx

Is/Is not append only

Is/Is not a device file

Is/Is not a Unix character device

Is/Is not a Unix block device

Is/Is not a symbolic link
*/

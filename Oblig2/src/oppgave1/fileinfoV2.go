
package main

import (
"fmt"
"log"
"os"
)

func mode(s string) {

	fileInfo, err := os.Lstat(s)
	fileSize := float32(fileInfo.Size())

	if err != nil {
		log.Fatal(err)
	}

	bytes := float64(fileSize)
	KB := bytes / 1024
	MB := KB / 1024
	GB := MB / 1024

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("#########################################")
	fmt.Println("Information about", fileInfo.Name())
	fmt.Println(" ")
	fmt.Printf("%v %f %v %f %v %f %v %.15g %v\n ", "Size :", bytes, "bytes,", KB, "KB,", MB, "MB,", GB, "GB") //fant ingen bedre løsning
	fmt.Println(" ")

	switch mode := fileInfo.Mode(); {
	case mode.IsDir():
		fmt.Println("Is a directory")
	case !mode.IsDir():
		fmt.Println("Is not a directory")
	}

	switch mode := fileInfo.Mode(); {
	case mode.IsRegular():
		fmt.Println("Is a regular file")
	case !mode.IsRegular():
		fmt.Println("Is not regular file")
	}

	fmt.Println("Has Unix permission bits : ", fileInfo.Mode().Perm())

	switch mode := fileInfo.Mode(); {
	case mode&os.ModeAppend != 0:
		fmt.Println("Is append only")
	case mode&os.ModeAppend == 0:
		fmt.Println("Is not append only")
	}

	switch mode := fileInfo.Mode(); {
	case mode&os.ModeDevice != 0:
		fmt.Println("Is a device file")
	case mode&os.ModeAppend == 0:
		fmt.Println("Is not a device file")
	}

	//her skal det stå noe om unix block device, aka moby dick.

	switch mode := fileInfo.Mode(); {
	case mode&os.ModeSymlink != 0:
		fmt.Println("Is a symbolic link")
	case mode&os.ModeSymlink == 0:
		fmt.Println("Is not symbolic link")
	}

}

func main() {

	filename := os.Args[1] // Må nå ha et argument for å kjøre i cmd
	mode(filename)
}
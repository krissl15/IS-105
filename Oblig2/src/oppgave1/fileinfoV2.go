package main

import (
	"fmt"
	"log"
	"os"
)

func mode(s string) {
	fi, err := os.Lstat("hello.txt")
	fileInfo, err := os.Lstat("hello.txt")
	fileSize := float32(fileInfo.Size())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Information about", fileInfo.Name())
	fmt.Println("Size : ", fileSize, "bytes,", fileSize/1000, "KB,", fileSize/1000000, "MB,", fileSize/1000000000, "GB")
	switch mode := fi.Mode(); {
	case mode.IsDir():
		fmt.Println("directory")
	case !mode.IsDir():
		fmt.Println("not a directory")
	}

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")
	case !mode.IsRegular():
		fmt.Println("not regular file")
	}

	fmt.Println("Has Unix permission bits : ", fileInfo.Mode().Perm())

	switch mode := fi.Mode(); {
	case mode&os.ModeAppend != 0:
		fmt.Println("append only")
	case mode&os.ModeAppend == 0:
		fmt.Println("not append only")
	}

	switch mode := fi.Mode(); {
	case mode&os.ModeDevice != 0:
		fmt.Println("Is a device file")
	case mode&os.ModeAppend == 0:
		fmt.Println("Is not a device file")
	}

	switch mode := fi.Mode(); {
	case mode&os.ModeCharDevice != 0:
		fmt.Println("Is a unix character device")
		fmt.Println("Is not a unix block device")
	case mode&os.ModeAppend == 0:
		fmt.Println("Is not a unix character device")
		fmt.Println("is a unix block device")
	}

	switch mode := fi.Mode(); {
	case mode&os.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode&os.ModeSymlink == 0:
		fmt.Println("not symbolic link")
	}

}

func main() {

	filename := os.Args[1] // Må nå ha et argument for å kjøre i cmd
	mode(filename)
}

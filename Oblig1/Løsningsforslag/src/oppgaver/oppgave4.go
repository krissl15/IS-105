package oppgaver

import "fmt"

func Oppgave4a(){
	var extascii []byte
	for i := 0x80 ; i <= 0xFF ; i++{
		extascii = append(extascii, byte(i))
	}
	iterateOverExtendedASCIIStringLiteral(string(extascii))
}
func iterateOverExtendedASCIIStringLiteral(tekst string){
	for i := 0; i < len(tekst); i++ {
		fmt.Printf("%X %c %b\n", tekst[i], tekst[i], tekst[i])
	}
}

func Oppgave4b(){
	fmt.Print(ExtendedASCIIText())
}

func ExtendedASCIIText() string{
	tekst := []byte("\x22\x20\x80\x20\xF7\x20\xBE\x20\x64\x6F\x6C\x6C\x61\x72\x20\x22")
	return string(tekst)
}

func Oppgave4c(){
	/*
	i mappen Oblig1\src\oppgaver, kjør: go test iso_test.go oppgave4.go
	testen vil feile da den inneholder ikke extended ascii.
	 */

}
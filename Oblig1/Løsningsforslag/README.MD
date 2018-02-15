Fjern kommentar på oppgave som skal kjøres "//->"

## Programmeringsoppgaver:
Oppgave 2 A
   
    En modifisert bubble sort bassert på wikipedia eksemplet.

Oppgave 2 B

    Kopiert testen og endret target til bubble sort i oppgave2a
    fra src mappen: go test -bench=.

Oppgave 2 C

    BigO
    Bubblesort sammenligner første element med andre element, og swapper de hvis første element er større enn andre.
    Slik fortsetter den loopen til hele listen er sortert
    Bubblesort har to looper og har en eksponentiell vekstratesom gir WORST CASE / BIG O = n^2

     Eksemplet teller antall looper for hver iterasjon og swap i oppgaver.Oppgave2Counts()

Oppgave 3
	
	Her kan man lytte til system kall og definere egen print.
	Her brukes go func() for å kjøre funksjonen synkront


Oppgave 4

	Det er mange måter å løse denne oppgaven på.
	a) er løst ved å fylle en byte slice med "DEC" fra ascii tabellen
	b) er løst ved å fylle en byte slice med "HEX" fra ascii tabellen
    c) tester om det finnes symboler utom extended ascii. I mappen Oblig1\src\oppgaver, kjør: go test iso_test.go oppgave4.go
    
    

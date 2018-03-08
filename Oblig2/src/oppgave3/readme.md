<h1>Oppgave3</h1>
a)<br>
Jeg har valgt å bruke tre channels i denne oppgaven, dette kun for å litt grundigere ilustrere bruken av dem. 
C1 og C2 lagrer hvert sitt tall fra cmd input, C3 lagrer summen når C1 og C2 har blitt summert. <br>
Bruker plassene i os.Args ([]) for å lese inn tallene som blir skrivet i cmd. konverterer dem til ints.<br>
funksjonA bruker tall fra stdin som argument, og dytter dem inn i hver sin channel. Her kalles og en .wait for å forsikre om at ikke programmet går videre før funksjonB har summert og dyttet summen i C3, hvor den så setter wg.done. Når tallet er summert i funksjonB, henter funksjonA ut tallet fra C3 og printer det ut. Det er og en annen waitgroup i funksjonA som venter på at tallet blir printet ut. 



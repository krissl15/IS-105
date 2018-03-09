<h1>Oppgave3</h1>
Nei, pcen er ikke treg. Jeg har lagt til time.Sleep() noen steder slik at det blir lettere å teste SIGINT i d). 
<h2>a)</h2>

<h3>addup.go</h3>
Jeg har valgt å bruke tre channels i denne oppgaven, dette kun for å litt grundigere ilustrere bruken av dem. 
C1 og C2 lagrer hvert sitt tall fra cmd input, C3 lagrer summen når C1 og C2 har blitt summert. <br>
Bruker plassene i os.Args (som er en slice) for å lese inn tallene som blir skrivet i cmd. konverterer dem til ints.<br>
FunksjonA kjøres som goroutine, der en waitgroup har blitt etablert. <br>
wg2.Add(1) gjør at der wg2.Wait() er kalt, vil funksjonen vente med å gå videre til wg2Done() blir kalt. Dette er for å passe på at mainfunksjonen ikke vil stanse før funksjonA har printet ut summen. <br>
"go funksjonB" blir kalt inni funksjonA, med samme waitgroup-system som forklart tidligere. funksjonA vil vente til funksjonB har summert tallet før den prøver å skrive ut noe. 
<br>

<br>
<h2>b)</h2>
Slik jeg har tolket oppgaven skal addtofile lese to tall fra cmd og legge disse inn i en tekstfil. sumfromfile skal tallene fra samme tekstfilen, og skrive summen inn i tekstfilen. addtofile skal også printe ut summen etter den har blitt summert.
<h5>Så dette er recap av løsningen min:</h5>
addtofile må kjøres med to tall som parametre eks. addtofile 2 2. sumfromfile leser tallene, summerer dem, og skriver summen i tekstfilen. addtofile vil lese tekstfilen når sumtofile er ferdig, og printe ut summen. sumfromfile blir kjørt i addtofile ved hjelp av exec.LookPath() og exec.Command(). Dette gjør at en kun trenger å kjøre addtofile, og sumfromfile blir kjørt automatisk. 
<h3>addtofile.go</h3>
For å sjekke at inputen i cmd er med to tall, bruker jeg en if-else med len(os.Args) som condition, dette sjekker lengden på slicen som blir opprettet når en bruker os.Args. <br>
om len(os.Args)==3 vil en vil med navn "numbers.txt.lock" bli opprettet med tallene som innhold. Om filen allerede finnes vil den bli overridet til å inneholde de nye tallene. <br>
Om en skriver mer eller mindre en to tall som parameter, vil en fmt.Println() skrive en påminnelse om å skrive to tall.<br>
Etter filen numbers.txt.lock har blitt opprettet, vil programmet kjøre sumfromfile.exe. Etter sumfromfile.exe er kjørt, vil addtofile igjen lese innholdet i tekstfilen, og så printe ut (det oppdaterte tallet) til stdout. <br>

<h3>sumfromfile.go</h3>
Dette programmet leser "numbers.txt.lock", gjør innholdet til string (fra []byte), og splitter stringsa. Jeg brukte direkte plassene i slicen til å få tak i tallene først, men uten å splitte ender dette opp med å kun få tak i enkle siffre. Når jeg bruker string.Split blir tallene lagret som de er. <br>
Det er ikke noe poeng å sjekke om filen inneholder to tall, ta sumfromfile ikke vil bli kjørt i addtofile.exe om ikke det blir brukt to parametre. <br>
Mer detaljer om koden står i koden. 
<h2>c) </h2>

<h4>Error handling i addup.go</h4>
I addup.go er første errorhandling i strconv.Atoi(linje24). Her vil det komme en error-melding dersom en skriver noe annet en tall. Det vil da komme en "invalid syntax" melding, da funksjonen ikke godtar annet en integer. <br>
Om en skriver mindre enn to tall som argument når en kjører programmet, vil det komme en "index out of range" melding, da koden prøver å lagre os.Args[1] og os.Args[2], som da ikke finnes. <br>
Om en skriver flere en to tall vil kun de to første tallene summeres. Her kunne det kanskje vært en form for errorhandling (Om jeg hadde hatt tid til å implementere det.)

<h4> error handling i sumfromfile.go og addtofile.go</h4>
I denne filen laget jeg en egen funksjon for error-handlingen, så jeg slipper å skrive if-setningen hver gang jeg skriver noe med en error type.<br>
I addtofile.go og sumfromfile.go har ioutil.ReadFile en errortype. Denne erroren blir trigget om programmet ikke finner filen som blir spesifisert (the system cannot fint the file specified.)<br>
os.Create i addtofile.go har også en errortype. Ifølge librariet vil denne funksjonen oppdage PathErrors. Altså om en prøver å lage en fil på et sted som ikke er muligt. Dette ville vært et større problem om en spesifiserte pathen til filen som skal kjøres fra et annet directory(noe jeg ikke har gjort.) <br>
exec.Lookpath() i addtofile.go har også en errortype. Denne blir trigget om .exe filen spesifisert ikke blir funnet. Dette igjen er en PathError (executable file not found in %PATH%.)<br>
If-setningen som sjekker om det er angitt to tall som argument kan vel og kalles en form for errorhandling. Den blir brukt for å kontrollere brukeren, og gi beskjed om noe er gjort feil. 



<h2>d)</h2>
SIGINT håndtering er lagt til i egne funksjoner som kjøres i mainen til alle programmene i a) og b).<br>

<h2>e)</h2>
addup.exe, addtofile.exe, og sumfromfile.exe ligger i bin mappen. (Opprettet på windows)









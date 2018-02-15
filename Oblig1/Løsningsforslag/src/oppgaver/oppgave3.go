package oppgaver

import (
	"os"
	"fmt"
	"os/signal"
	"syscall"
)

func Oppgave3() {
	c := make(chan os.Signal, 0x2)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		fmt.Printf("CTRL+C endte loopen")
		os.Exit(1)
	}()
	for i:= 0; ; i++ {
		fmt.Println(i)
	}

}


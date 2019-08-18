package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	a "github.com/l3njo/remorse/app"

	"github.com/eiannone/keyboard"
)

var signals chan os.Signal

// cleanup handles application shutdown.
func cleanup() {
	keyboard.Close()
	fmt.Println("\nGoodbye!")
}

func init() {
	signals = make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signals
		cleanup()
		os.Exit(1)
	}()
}

func main() {
	app := &a.Application{}
	app.Init()
	if err := app.Run(); err != nil {
		log.Fatalf("\n%s", err)
	}

	cleanup()
}

package util

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func WaitSignal(tearDownFunc func()) {
	var sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan)
	for sig := range sigChan {
		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			log.Printf("terminated by signal %v\n", sig)
			log.Printf("Tearing down...")
			tearDownFunc()
			return
		} else {
			log.Printf("received signal: %v, ignore\n", sig)
		}
	}
}

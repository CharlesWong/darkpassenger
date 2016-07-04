package main

import (
	"flag"
	"github.com/charleswong/darkpassenger/account"
	db "github.com/charleswong/darkpassenger/db/sqlite"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
	"time"
)

func Start() {
	go account.StartAccountService()
	go account.CleanAccountRoutine()
}

func tearDown() {
	// TODO: Close DB connections.
}

func main() {
	configFile := flag.String("config", "./dp.conf", "Config file.")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to file")
	flag.Parse()

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	err := account.InitConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = db.InitDB(account.GetConfig().DataFile)
	if err != nil {
		log.Fatal(err)
	}
	Start()

	defer tearDown()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// Handle exiting signals and process.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-sigChan:
			return
		default:
			time.Sleep(time.Second)
		}
	}
}

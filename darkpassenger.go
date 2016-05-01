package main

import (
	"flag"
	"github.com/CharlesWong/darkpassenger/config"
	"github.com/CharlesWong/darkpassenger/regiservice"
	"github.com/CharlesWong/darkpassenger/transproxy"
	"github.com/CharlesWong/darkpassenger/tunnel"
	"github.com/CharlesWong/darkpassenger/util"
	"log"
)

func tearDown() {

}

func main() {
	// log.SetFlags(log.Lshortfile)
	var secret string
	flag.StringVar(&config.LogTo, "logto", "stdout", "stdout or syslog")
	flag.StringVar(&config.FrontEndAddr, "listen", "", "host:port listen on")
	// flag.StringVar(&config.BackEndAddr, "backend", "", "host:port of the backend")
	flag.StringVar(&config.RegiServiceAddr, "regiservice", "", "host:port of the backend")
	flag.StringVar(&config.CryptoMethod, "crypto", "aes-128-cfb", "encryption method")
	flag.StringVar(&secret, "secret", "", "password used to encrypt the data")
	flag.BoolVar(&config.ClientMode, "clientmode", true, "if running at client mode")
	flag.BoolVar(&config.VerboseTransproxy, "verbose_transproxy", false, "should every proxy request be logged to stdout")
	flag.Parse()

	util.SetupLogOptions(config.LogTo)

	if !config.ClientMode {
		if config.BackEndAddr == "" {
			config.BackEndAddr = ":9999"
		}
		go func() {
			transproxy.StartTransProxy(config.BackEndAddr, config.VerboseTransproxy)
		}()
		go func() {
			regiservice.StartRegiService()
		}()
	} else {
		remote := regiservice.SelectWorkerAddr()
		if remote.Version > config.Version {
			log.Fatalf("Please update client from %s => %s\n", config.Version, remote.Version)
		}
		config.BackEndAddr = remote.WorkerAddr
		log.Printf("Selected avaiable backend: %s\n", config.BackEndAddr)
	}

	t := tunnel.NewTunnel(config.FrontEndAddr, config.BackEndAddr, config.ClientMode, config.CryptoMethod, secret, 4096)
	go t.Start()

	log.Printf("DarkPassenger %s started.\n\nListening at %s", config.Version, config.FrontEndAddr)

	util.WaitSignal(tearDown)
}

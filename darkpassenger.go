package main

import (
	"flag"
	"github.com/charleswong/darkpassenger/config"
	"github.com/charleswong/darkpassenger/regiservice"
	"github.com/charleswong/darkpassenger/transproxy"
	"github.com/charleswong/darkpassenger/tunnel"
	"github.com/charleswong/darkpassenger/util"
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
	flag.StringVar(&config.ProxyServerAddr, "backend", "", "host:port of the Squid")
	flag.StringVar(&config.RegiServiceAddr, "regiservice", "", "host:port of the backend")
	flag.StringVar(&config.CryptoMethod, "crypto", "aes-128-cfb", "encryption method")
	flag.StringVar(&secret, "secret", "", "password used to encrypt the data")
	flag.BoolVar(&config.ClientMode, "clientmode", true, "if running at client mode")
	flag.Parse()

	util.SetupLogOptions(config.LogTo)

	if !config.ClientMode {

		transproxy.StartHttpProxy()
		if config.BackEndAddr == "" {
			config.BackEndAddr = config.ProxyServerAddr
		}

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

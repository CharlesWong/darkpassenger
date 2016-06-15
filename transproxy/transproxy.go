package transproxy

import (
	"github.com/charleswong/darkpassenger/config"
	"github.com/getlantern/golog"
	"github.com/getlantern/http-proxy/commonfilter"
	"github.com/getlantern/http-proxy/forward"
	"github.com/getlantern/http-proxy/httpconnect"
	"github.com/getlantern/http-proxy/listeners"
	"github.com/getlantern/http-proxy/server"
	"net"
	"os"
	"time"
)

var (
	// testingLocal = false
	log = golog.LoggerFor("http-proxy")

	// keyfile   = "key"
	// certfile  = "cert"
	https = "https"
	// addr      = ":8080"
	maxConns  = 0
	idleClose = 10
)

var (
	httpProxyAddr string
	tlsProxyAddr  string
)

func StartHttpProxy() {
	var err error

	// Set up HTTP chained server
	httpProxyAddr, err = setupNewHTTPServer(0, 30*time.Second)
	if err != nil {
		log.Error("Error starting proxy server")
		os.Exit(1)
	}

	config.ProxyServerAddr = httpProxyAddr
	// Set up HTTPS chained server
	tlsProxyAddr, err = setupNewHTTPSServer(0, 30*time.Second)
	if err != nil {
		log.Error("Error starting proxy server")
		os.Exit(1)
	}
}

func basicServer(maxConns uint64, idleTimeout time.Duration) *server.Server {

	// Middleware: Forward HTTP Messages
	forwarder, err := forward.New(nil, forward.IdleTimeoutSetter(idleTimeout))
	if err != nil {
		log.Error(err)
	}

	// Middleware: Handle HTTP CONNECT
	httpConnect, err := httpconnect.New(forwarder, httpconnect.IdleTimeoutSetter(idleTimeout))
	if err != nil {
		log.Error(err)
	}

	// Middleware: Common request filter
	commonHandler, err := commonfilter.New(httpConnect, false)
	if err != nil {
		log.Error(err)
	}

	// Create server
	srv := server.NewServer(commonHandler)

	// Add net.Listener wrappers for inbound connections
	srv.AddListenerWrappers(
		// Limit max number of simultaneous connections
		func(ls net.Listener) net.Listener {
			return listeners.NewLimitedListener(ls, maxConns)
		},
		// Close connections after 30 seconds of no activity
		func(ls net.Listener) net.Listener {
			return listeners.NewIdleConnListener(ls, idleTimeout)
		},
	)

	return srv
}

func setupNewHTTPServer(maxConns uint64, idleTimeout time.Duration) (string, error) {
	s := basicServer(maxConns, idleTimeout)
	var err error
	ready := make(chan string)
	wait := func(addr string) {
		log.Debugf("Started HTTP proxy server at %s", addr)
		ready <- addr
	}
	go func(err *error) {
		if *err = s.ListenAndServeHTTP("localhost:0", wait); err != nil {
			log.Errorf("Unable to serve: %v", err)
		}
	}(&err)
	return <-ready, err
}

func setupNewHTTPSServer(maxConns uint64, idleTimeout time.Duration) (string, error) {
	s := basicServer(maxConns, idleTimeout)
	var err error
	ready := make(chan string)
	wait := func(addr string) {
		log.Debugf("Started HTTPS proxy server at %s", addr)

		ready <- addr
	}
	go func(err *error) {
		if *err = s.ListenAndServeHTTPS("localhost:0", "key.pem", "cert.pem", wait); err != nil {
			log.Errorf("Unable to serve: %v", err)
		}
	}(&err)
	addr := <-ready
	if err != nil {
		return "", err
	}
	// serverCertificate, err = keyman.LoadCertificateFromFile("cert.pem")
	return addr, err
}

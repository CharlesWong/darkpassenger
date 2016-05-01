package transproxy

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func StartTransProxy(listen string, verbose bool) {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = verbose
	log.Println(http.ListenAndServe(listen, proxy))
}

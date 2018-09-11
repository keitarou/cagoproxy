package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/moul/http2curl"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

var verbose *bool

func main() {

	listen := flag.String("listen", "8000", "listen port")
	proxypath := flag.String("proxy_pass", "http://localhost:8080", "proxy path")
	logfilepath := flag.String("logfile", "/tmp/cagoproxy.log", "log file path")
	verbose = flag.Bool("verbose", false, "verbose options")
	flag.Parse()

	vlog("listen: %s, proxy_path: %s, log file path: %s, verbose: %s", *listen, *proxypath, *logfilepath, *verbose)

	logger := logrus.New()
	logger.Formatter = new(logrus.JSONFormatter)
	f, _ := os.OpenFile(*logfilepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	logger.Out = f

	director := func(request *http.Request) {
		schemeAndHost := strings.SplitN(*proxypath, "://", 2)
		url := *request.URL
		url.Scheme = schemeAndHost[0]
		url.Host = schemeAndHost[1]
		header := request.Header

		body := []byte("")
		if request.Body != nil {
			rbody, err := ioutil.ReadAll(request.Body)
			if err != nil {
				vlog("Error reading body: %v", err)
				return
			}
			body = rbody
		}
		buffer := bytes.NewBuffer(body)

		req, err := http.NewRequest(request.Method, url.String(), buffer)
		req.Header = header
		if err != nil {
			log.Fatal(err.Error())
		}
		*request = *req
		curlDump(request, logger)
	}
	rp := &httputil.ReverseProxy{Director: director}
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", *listen),
		Handler: rp,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}

func curlDump(request *http.Request, logger *logrus.Logger) error {
	// url := *request.URL
	// log.Printf("METHOD: %s", request.Method)
	// log.Printf("SCHEME: %s", url.Scheme)
	// log.Printf("HOST: %s", url.Host)
	// log.Printf("PATH: %s", url.String())
	// for hkey, _ := range request.Header {
	// 	log.Printf("hkey: %s, hvalue: %s", hkey, request.Header.Get(hkey))
	// }
	command, err := http2curl.GetCurlCommand(request)
	if err != nil {
		return err
	}
	vlog("curl: %s", command)
	logger.Info(command)
	return nil
}

func vlog(format string, v ...interface{}) {
	if !*verbose {
		return
	}
	log.Printf(format, v)
}

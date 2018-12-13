package core

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

// CallbackHandler handle oauth callback
type CallbackHandler func(state, code string) (int, error)

// Web default web server
type Web struct {
	Server   *http.Server
	Callback CallbackHandler
}

// OpenBrowser opens the URL to authentication
func (*Web) OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func (web *Web) callback(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	state := values.Get("state")
	code := values.Get("code")

	if state != "1234" {
		log.Fatal(errors.New("state code did not match expected"))
	}

	statusCode, err := web.Callback(state, code)
	if err != nil {
		w.WriteHeader(statusCode)
		log.Fatal(err)
	} else {
		w.Write([]byte("Login Successful, you may now close this tab."))
	}
}

// Start start the web server
func (web *Web) Start(c CallbackHandler) {
	web.Server = &http.Server{Addr: ":8080"}
	web.Callback = c

	http.HandleFunc("/oauth2", web.callback)

	err := web.Server.ListenAndServe()
	if err != nil {
		log.Fatal()
	}
}

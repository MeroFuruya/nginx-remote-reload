package main

import (
	"log"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/alexflint/go-arg"
)

var Version string

type nginxSignal string

const (
	reload nginxSignal = "reload"
	stop   nginxSignal = "stop"
	start  nginxSignal = "start"
)

type Args struct {
	Port uint16 `arg:"required,-p,help:port to listen on,env:NRS_PORT"`
}

func (Args) Description() string {
	return "A simple HTTP server that sends signals to an nginx process.\nCall /signal/<start|stop|reload> to send a signal to nginx."
}
func (Args) Version() string {
	return "nginx-remote-reload " + Version
}

func nginxSendSignal(signal nginxSignal) error {
	cmd := exec.Command("nginx", "-s", string(signal))
	return cmd.Run()
}

func reloadHandlerFactory(signal nginxSignal) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err = nginxSendSignal(signal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	var args Args
	arg.MustParse(&args)
	http.HandleFunc("/signal/reload", reloadHandlerFactory(reload))
	http.HandleFunc("/signal/stop", reloadHandlerFactory(stop))
	http.HandleFunc("/signal/start", reloadHandlerFactory(start))

	log.Printf("Listening on port %d\n", args.Port)

	http.ListenAndServe(":"+strconv.Itoa(int(args.Port)), nil)
}

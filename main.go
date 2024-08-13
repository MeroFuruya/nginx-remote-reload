package main

import (
	"net/http"
	"os/exec"
	"strconv"

	"github.com/alexflint/go-arg"
)

type nginxSignal string

const (
	reload nginxSignal = "reload"
	stop   nginxSignal = "stop"
	start  nginxSignal = "start"
)

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

func init() {
	http.HandleFunc("/signal/reload", reloadHandlerFactory(reload))
	http.HandleFunc("/signal/stop", reloadHandlerFactory(stop))
	http.HandleFunc("/signal/start", reloadHandlerFactory(start))
}

func main() {
	var args struct {
		port uint16 `arg:"-p,help:port to listen on"`
	}
	arg.MustParse(&args)
	http.ListenAndServe(":"+strconv.Itoa(int(args.port)), nil)
}

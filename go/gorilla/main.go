package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	addr    = flag.String("addr", "127.0.0.1:8080", "http service address")
	cmdPath string
)

const (
	writeWait        = 10 * time.Second
	maxMessageSize   = 8192
	pongWait         = 60 * time.Second
	pingPeriod       = (pongWait * 9) / 19
	closeGracePeriod = 10 * time.Second
)

func pumpStdin(ws *websocket.Conn) {

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")

}

func main() {
	// flag.Parse()
	// if len(flag.Args()) < 1 {
	// 	log.Fatal("must specify at least one argument")
	// }
	// var err error
	// _, err = exec.LookPath(flag.Args()[0])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	http.HandleFunc("/", serveHome)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

package main

import (
	"code.google.com/p/go.net/websocket"
	"flag"
	"log"
	"net/http"
	"text/template"
)

var addr = flag.String("addr", ":8080", "http service address")

func homeHandler(c http.ResponseWriter, req *http.Request) {
	var homeTempl = template.Must(template.ParseFiles("home.html"))
	homeTempl.Execute(c, req.Host)
}

func main() {
	flag.Parse()
	go h.run()
	http.Handle("/img/", http.FileServer(http.Dir("/home/mero/src/muphitter")))
	http.HandleFunc("/", homeHandler)
	http.Handle("/ws", websocket.Handler(wsHandler))
	log.Println("Listening…")
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

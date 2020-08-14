package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

var tpl = template.Must(template.ParseFiles("templates/index.gotpl"))

func statServer(ws *websocket.Conn) {
	fmt.Println("handle /ws")
	defer func() {
		log.Println("connection handler exits")
	}()
	fmt.Println("local address:", ws.LocalAddr())
	fmt.Println("remote address:", ws.RemoteAddr())
	buf := make([]byte, 4096)
	n, _ := ws.Read(buf)
	log.Printf("Received %d bytes: %q\n", n, buf[:n])
	ws.Write([]byte("B"))
	n, _ = ws.Read(buf)
	log.Printf("Received %d bytes: %q\n", n, buf[:n])
	time.Sleep(2 * time.Second)
	ws.Write([]byte("C"))
	n, _ = ws.Read(buf)
	log.Printf("Received %d bytes: %q\n", n, buf[:n])
	time.Sleep(2 * time.Second)
	ws.Write([]byte("D"))
	n, _ = ws.Read(buf)
	log.Printf("Received %d bytes: %q\n", n, buf[:n])
	time.Sleep(2 * time.Second)
	ws.Close()
}

func handleMainRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle /")

	err := tpl.ExecuteTemplate(w, "index.gotpl", nil)
	if err != nil {
		log.Fatal(err)
	}
}

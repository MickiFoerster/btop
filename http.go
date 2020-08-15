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

func loop(ch chan<- float64) {
	stat1, err := getCPUSample()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(200 * time.Millisecond)
	stat2, err := getCPUSample()
	if err != nil {
		log.Fatal(err)
	}

	cpuUsage := getCpuUsage(stat2, stat1)

	stat1 = stat2
	ch <- cpuUsage

	for {
		select {
		case <-time.After(500 * time.Millisecond):
			stat2, err := getCPUSample()
			if err != nil {
				log.Fatal(err)
			}

			cpuUsage = getCpuUsage(stat2, stat1)

			stat1 = stat2
			ch <- cpuUsage
		}
	}
}

func statServer(ws *websocket.Conn) {
	fmt.Println("handle /ws")
	defer func() {
		log.Println("connection handler exits")
	}()

	ch := make(chan float64)
	go loop(ch)

	for {
		select {
		case cpuUsage := <-ch:
			ws.Write([]byte(fmt.Sprintf("%0.2f%%", cpuUsage)))
		}
	}

	ws.Close()
}

func handleMainRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle /")

	err := tpl.ExecuteTemplate(w, "index.gotpl", nil)
	if err != nil {
		log.Fatal(err)
	}
}

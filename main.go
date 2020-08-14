// Measure CPU usage
// Hints can be found in https://www.idnt.net/en-US/kb/941772

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

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

func main() {
	ch := make(chan float64)
	go loop(ch)

	http.HandleFunc("/", handleMainRoute)
	http.Handle("/ws", websocket.Handler(statServer))

	go func() {
		for {
			select {
			case cpuUsage := <-ch:
				fmt.Printf("%f%%\n", cpuUsage)
			}
		}
	}()

	socket := ":12345"
	fmt.Println("Listening on ", socket)
	log.Fatal(http.ListenAndServe(socket, nil))
}

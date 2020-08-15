// Measure CPU usage
// Hints can be found in https://www.idnt.net/en-US/kb/941772

package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	http.HandleFunc("/", handleMainRoute)
	http.Handle("/ws", websocket.Handler(statServer))

	socket := ":12345"
	fmt.Println("Listening on ", socket)
	log.Fatal(http.ListenAndServe(socket, nil))
}

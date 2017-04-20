package main

import (
	"log"
	"net/http"
	"time"
)

// once this is running, spy on the connection to make sure it's chunking as expected
// sudo ngrep -W byline -d lo0 port 6666
// and then request http://localhost:6666 from a separate program
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		timeChan := time.NewTimer(30 * time.Second).C
		ticker := time.NewTicker(5 * time.Second).C

		for {
			select {
			case <-timeChan:
				log.Println("Timer expired")
				w.Write([]byte("done.\n"))
				return
			case <-ticker:
				log.Println("tick")
				_, err := w.Write([]byte("fake data\n"))
				if err != nil {
					log.Printf("failed to write to client: %v", err)
					return
				}
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
			}
		}
	})

	log.Println("Started up...")

	log.Fatal(http.ListenAndServe(":6666", nil))
}

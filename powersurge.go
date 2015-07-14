package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (
	strip    = flag.String("strip", "10.1.1.213", "IP address of the power strip")
	host     = flag.String("host", "http://www.google.com", "Address of site to use to test connectivity")
	socket   = flag.Int("socket", 1, "Number of the socket the router is connected to")
	interval = flag.String("interval", "5m", "Interval in which to run the test. Example: 10s, 5m, 1h")
)

func TestConnection() bool {
	log.Printf("Testing connection to %s..", *host)
	_, err := http.Get(*host)

	if err != nil {
		log.Printf("Error: %v", err)
		return false
	}

	return true
}

func RestartRouter() {
	vals := []string{"0", "1"}

	for _, i := range vals {
		http.PostForm(fmt.Sprintf("http://%s/", *strip),
			url.Values{fmt.Sprintf("cte%d", *socket): {i}})
		time.Sleep(3 * time.Second)
	}
	log.Print("Router is powering back up.")
}

func main() {
	flag.Parse()

	for {
		if !TestConnection() {
			log.Printf("Connection failed. Restarting router.")
			RestartRouter()
		}

		log.Printf("Sleeping for %s..", *interval)
		t, _ := time.ParseDuration(*interval)
		time.Sleep(t)
	}
}

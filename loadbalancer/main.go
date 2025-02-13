package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	loadbalancer "github.com/learn-golang/loadbalancer/core"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <port>")
		os.Exit(1)
	}

	port := os.Args[1]
	lb := loadbalancer.NewLoadBalancer()
	lb.AddServer(loadbalancer.NewServer("https://www.google.com"))
	lb.AddServer(loadbalancer.NewServer("https://www.yahoo.com"))
	lb.AddServer(loadbalancer.NewServer("https://www.facebook.com"))
	lb.CheckServersHealth()

	http.HandleFunc("/", lb.ServeHTTP)

	log.Println("Starting load balancer on :" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

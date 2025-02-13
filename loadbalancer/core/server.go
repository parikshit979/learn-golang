package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Server represents a server in the load balancer pool
type Server struct {
	URL          *url.URL
	Alive        bool
	ReverseProxy *httputil.ReverseProxy
}

// NewServer creates a new Server instance
func NewServer(rawurl string) *Server {
	target, err := url.Parse(rawurl)
	if err != nil {
		log.Fatalf("Failed to parse URL: %v", err)
	}
	return &Server{
		URL:          target,
		Alive:        false,
		ReverseProxy: httputil.NewSingleHostReverseProxy(target),
	}
}

// MarkAsDead marks the server as dead
func (s *Server) MarkAsDead() {
	s.Alive = false
}

// MarkAsAlive marks the server as alive
func (s *Server) MarkAsAlive() {
	s.Alive = true
}

// IsAlive pings the server to check if it's alive
func (s *Server) IsAlive() bool {
	resp, err := http.Head(s.URL.String())
	if err != nil || resp.StatusCode != http.StatusOK {
		s.Alive = false
		return s.Alive
	}
	s.Alive = true
	return s.Alive
}

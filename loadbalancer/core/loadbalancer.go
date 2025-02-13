package loadbalancer

import (
	"net/http"
	"sync/atomic"
)

type LoadBalancer struct {
	servers     []*Server
	serverCount uint32
}

// NewLoadBalancer creates a new LoadBalancer instance
func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		servers:     make([]*Server, 0),
		serverCount: 0,
	}
}

// AddServer adds a server to the load balancer pool
func (lb *LoadBalancer) AddServer(server *Server) {
	lb.servers = append(lb.servers, server)
	lb.serverCount++
}

// RemoveServer removes a server from the load balancer pool
func (lb *LoadBalancer) RemoveServer(server *Server) {
	for i, s := range lb.servers {
		if s == server {
			lb.servers = append(lb.servers[:i], lb.servers[i+1:]...)
			lb.serverCount--
			break
		}
	}
}

// NextServer returns the next server in the pool.
// It uses a round-robin algorithm to select the next server.
func (lb *LoadBalancer) nextServer() *Server {
	next := atomic.AddUint32(&lb.serverCount, 1)
	server := lb.servers[next%uint32(len(lb.servers))]
	for !server.IsAlive() {
		server = lb.servers[next%uint32(len(lb.servers))]
	}
	return server
}

// CheckServersHealth checks the health of all servers in the pool
func (lb *LoadBalancer) CheckServersHealth() {
	for _, s := range lb.servers {
		if s.IsAlive() {
			s.MarkAsAlive()
		} else {
			s.MarkAsDead()
		}
	}
}

// ServeHTTP routes the request to the next server in the pool
func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server := lb.nextServer()
	if server == nil {
		http.Error(w, "No servers available", http.StatusServiceUnavailable)
		return
	}
	server.ReverseProxy.ServeHTTP(w, r)
}

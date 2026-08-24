package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/tomasen/realip"
	"golang.org/x/time/rate"
)

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

func rateLimiter(next http.Handler, rps, burst int) http.Handler {
	var (
		mut     sync.Mutex
		clients = make(map[string]*client)
	)
	go func() {
		time.Sleep(time.Minute)
		mut.Lock()
		defer mut.Unlock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 5*time.Minute {
				delete(clients, ip)
			}
		}
	}()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mut.Lock()

		defer mut.Unlock()
		ipAddr := realip.FromRequest(r)
		if _, ok := clients[ipAddr]; !ok {
			clients[ipAddr].limiter = rate.NewLimiter(rate.Limit(rps), burst)
			clients[ipAddr].lastSeen = time.Now()
		}
		if !clients[ipAddr].limiter.Allow() {
			http.Error(w, "rate limite exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

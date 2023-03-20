package main

import (
	"sync"

	"golang.org/x/time/rate"
)

type RateLimitedClient struct {
	limiter *rate.Limiter
}

func NewRateLimitedClient(r rate.Limit, b int) *RateLimitedClient {
	return &RateLimitedClient{
		limiter: rate.NewLimiter(r, b),
	}
}

func (rlc *RateLimitedClient) Allow() bool {
	return rlc.limiter.Allow()
}

type ClientRateLimiter struct {
	mu        sync.Mutex
	rateLimit rate.Limit
	burst     int
	clients   map[string]*RateLimitedClient
}

func NewClientRateLimiter(r rate.Limit, b int) *ClientRateLimiter {
	return &ClientRateLimiter{
		rateLimit: r,
		burst:     b,
		clients:   make(map[string]*RateLimitedClient),
	}
}

func (crl *ClientRateLimiter) getClientLimiter(client string) *RateLimitedClient {
	crl.mu.Lock()
	defer crl.mu.Unlock()

	limiter, exists := crl.clients[client]
	if !exists {
		limiter = NewRateLimitedClient(crl.rateLimit, crl.burst)
		crl.clients[client] = limiter
	}

	return limiter
}

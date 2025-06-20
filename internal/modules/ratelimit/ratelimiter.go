package ratelimit

import (
	"flash/internal/core"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	maxRequest = 5
	timeWindow = time.Minute
)

type FlashRateLimit struct {
	requests    map[string][]time.Time
	maxRequests int
	mu          sync.Mutex
	Interval    time.Duration
}

type Client struct {
	Ip string
}

func NewFlashRateLimit(maxRequests int, interval time.Duration) *FlashRateLimit {
	return &FlashRateLimit{
		requests:    make(map[string][]time.Time),
		maxRequests: maxRequests,
		Interval:    interval,
	}
}

func init() {
	core.RegisterModule(NewFlashRateLimit(1, time.Minute))
}

func (p *FlashRateLimit) Name() string {
	return "rate limit"
}

func (p *FlashRateLimit) Match(r *http.Request) bool {
	return true
}

func (p *FlashRateLimit) Init(config map[string]interface{}) error {
	if mr, ok := config["max_requests"].(int); ok {
		p.maxRequests = mr
	}
	if iv, ok := config["interval_seconds"].(int); ok {
		p.Interval = time.Duration(iv) * time.Second
	}
	return nil
}

func (rl *FlashRateLimit) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	ip = strings.Split(ip, ":")[0]

	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.requests[ip] = append(rl.requests[ip], time.Now())

	rl.requests[ip] = dropOldRequests(rl.requests[ip])

	if len(rl.requests[ip]) > maxRequest {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
	}
}

func dropOldRequests(req []time.Time) []time.Time {
	var validReq []time.Time
	for _, v := range req {
		if time.Since(v) <= timeWindow {
			validReq = append(validReq, v)
		}
	}
	return validReq
}

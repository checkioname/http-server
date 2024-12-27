package ratelimit

import (
	"net/http"
	"strings"
	"sync"
	"time"
)


type RateLimiter struct {
  requests map[string] []time.Time
  maxRequests int
  mu sync.Mutex
  Interval time.Duration
}

const (
  maxRequest = 5
  timeWindow = time.Minute
)


type Client struct {
  Ip string
}


// func (rl *RateLimiter) Apply(next http.Handler) http.Handler {
//   return http.HandlerFunc(rl.limit())
// }

func (rl *RateLimiter) limit(w http.ResponseWriter, r *http.Response) {
  ip := r.Request.RemoteAddr
  ip = strings.Split(ip, ":")[0]

  rl.mu.Lock()
  defer rl.mu.Unlock()
  
  rl.requests[ip] = append(rl.requests[ip],time.Now())

  rl.requests[ip] = dropOldRequests(rl.requests[ip])
  
  if len(rl.requests[ip]) > maxRequest {
    http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
  }
  
}


func dropOldRequests(req []time.Time) []time.Time{
  var validReq []time.Time
  for _, v := range req {
    if time.Since(v) <= timeWindow {
      validReq = append(validReq, v)
    }
  }
  return validReq
}

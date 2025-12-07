package middleware

import (
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Visitor struct {
	tokens     int
	lastUpdate time.Time
}

type APIKeyRateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*Visitor
	limit    int           // dk başına üretilecek token sayısı
	burst    int           // maksimum token kapasitesi
	window   time.Duration // refill dönemi
}

func NewAPIKeyRateLimiter(limit int, burst int) *APIKeyRateLimiter {
	return &APIKeyRateLimiter{
		visitors: make(map[string]*Visitor),
		limit:    limit,
		burst:    burst,
		window:   time.Minute,
	}
}

// API KEY → Visitor bucket
func (rl *APIKeyRateLimiter) getVisitor(key string) *Visitor {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	v, exists := rl.visitors[key]
	if !exists {
		v = &Visitor{
			tokens:     rl.burst,
			lastUpdate: time.Now(),
		}
		rl.visitors[key] = v
		return v
	}

	// Token refill
	elapsed := time.Since(v.lastUpdate)
	refill := int(float64(rl.limit) * elapsed.Minutes())

	if refill > 0 {
		v.tokens += refill
		if v.tokens > rl.burst {
			v.tokens = rl.burst
		}
		v.lastUpdate = time.Now()
	}

	return v
}

func (rl *APIKeyRateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// -----------------------
		// 1) Authorization header kontrolü
		// -----------------------
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.JSON(401, gin.H{"error": "missing Authorization header"})
			c.Abort()
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "invalid Authorization format"})
			c.Abort()
			return
		}

		apiKey := parts[1]

		// -----------------------
		// 2) API key için bucket al
		// -----------------------
		visitor := rl.getVisitor(apiKey)

		// -----------------------
		// 3) rate limit enforcement
		// -----------------------
		if visitor.tokens <= 0 {
			c.JSON(429, gin.H{
				"error": "rate limit exceeded for this API key",
			})
			c.Abort()
			return
		}

		visitor.tokens-- // 1 token harca

		c.Next()
	}
}

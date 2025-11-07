package middlewares

import (
	"fmt"
	"go-api/internal/cache"
	"go-api/pkg/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware(CacheRepository cache.CacheRepository, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		key := fmt.Sprintf("rate_limit:%s", clientIp)

		//check
		val, err := CacheRepository.GetOne(key)
		count := 0
		if val != "" {
			count, _ = strconv.Atoi(val)
		}
		if err != nil && err.Error() != "redis: nil" {
			utils.Error(c, http.StatusInternalServerError, "Redis connection error")
			c.Abort()
		}

		if count >= limit {
			utils.Error(c, http.StatusTooManyRequests, "Too many requests,Please try again later ")
			c.Abort()
		}
		//increment
		newCount, _ := CacheRepository.Incr(key)
		if newCount == 1 {
			CacheRepository.Expire(key, window)
		}

		c.Next()
	}
}

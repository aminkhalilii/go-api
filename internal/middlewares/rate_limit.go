package middlewares

import (
	"fmt"
	"go-api/config"
	"go-api/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		key := fmt.Sprintf("rate_limit:%s", clientIp)

		//check
		count, err := config.RedisClient.Get(c, key).Int()
		if err != nil && err.Error() != "redis: nil" {
			utils.Error(c, http.StatusInternalServerError, "Redis connection error")
			c.Abort()
		}

		if count > limit {
			utils.Error(c, http.StatusTooManyRequests, "Too many requests,Please try again later ")
			c.Abort()
		}
		//increment
		pipe := config.RedisClient.TxPipeline()
		pipe.Incr(c, key)
		if count == 0 {
			pipe.Expire(c, key, window)
		}
		_, _ = pipe.Exec(c)

		c.Next()
	}
}

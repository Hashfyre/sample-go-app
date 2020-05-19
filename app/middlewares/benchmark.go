package middleware

import (
	"fmt"
	"math"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/hashfyre/sample-go-app/app/context"
)

// Benchmark defines the time taken for request to complete
func Benchmark() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set(context.RequestStartKey, start.Format(time.RFC3339))
		c.Next()
		elapsed := time.Since(start)
		fmt.Printf("Request took %v milliseconds\n", float64(elapsed.Nanoseconds())/math.Pow(float64(10), float64(6)))
	}
}

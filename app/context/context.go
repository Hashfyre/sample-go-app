package context

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 	serviceNameKey and others that define contents of gin.Context.Keys
const (
	ServiceNameKey   = "service_name"   // middleware.Context
	HostNameKey      = "host_name"      // middleware.Context
	RequestIDKey     = "request_id"     // middleware.Context
	RequestPathKey   = "request_path"   // middleware.Context
	RequestMethodKey = "request_method" // middleware.Context
	RequestStartKey  = "request_start"  // middleware.Benchmark
	UserIDKey        = "user_id"        // middleware.BasicAuth
)

func getCtxID(c *gin.Context, key string) (uuid.UUID, error) {
	ctxID, ok := c.Get(key)
	if !ok {
		log.Println(errMissingContextKey)
		return uuid.UUID{}, errMissingContextKey
	}

	id, ok := ctxID.(uuid.UUID)
	if !ok {
		log.Println(errBadContextKey)
		return uuid.UUID{}, errBadContextKey
	}

	return id, nil
}

// GetCtxRequestID returns the request_id from gin.Context.Keys
func GetCtxRequestID(c *gin.Context) (uuid.UUID, error) {
	return getCtxID(c, RequestIDKey)
}

// GetCtxUserID returns the user_id from gin.Context.Keys
func GetCtxUserID(c *gin.Context) (uuid.UUID, error) {
	return getCtxID(c, UserIDKey)
}

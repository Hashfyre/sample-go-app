package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/hashfyre/sample-go-app/app/context"
	"github.com/hashfyre/sample-go-app/app/types"
	p_uuid "github.com/hashfyre/sample-go-app/pkg/uuid"
)

const serviceName = "sample-go-app"

// AppContext adds metadata used for logging, tracing and debugging to the context object
func AppContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		hostName, ok := os.LookupEnv("HOSTNAME")
		if !ok {
			log.Println(errHostNameUnset)
			c.AbortWithStatusJSON(http.StatusInternalServerError, types.ResponseError{
				Code:    "ERR_INTERNAL_SERVER_ERROR",
				Message: errHostNameUnset.Error(),
			})
			return
		}

		c.Set(context.ServiceNameKey, serviceName)
		c.Set(context.HostNameKey, hostName)
		c.Set(context.RequestIDKey, p_uuid.MustUUID())
		c.Set(context.RequestMethodKey, c.Request.Method)
		c.Set(context.RequestPathKey, c.Request.URL.String())

		c.Next()
	}
}

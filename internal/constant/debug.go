package constant

import (
	"github.com/gin-gonic/gin"
	"github.com/metadiv-io/env"
)

var (
	// IS_DEBUG is the flag to enable/disable debug logs
	// By default, it is enabled in development mode and disabled in production mode based on GIN_MODE
	IS_DEBUG = env.Bool("IS_DEBUG", env.String("GIN_MODE", "") != gin.ReleaseMode)
)

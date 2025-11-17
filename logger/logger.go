package logger

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func init() {
	zerolog.TimeFieldFormat = time.RFC3339
	// Initialize the global logger once at package import
	Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func JsonLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		logEntry := map[string]interface{}{
			"time":          params.TimeStamp.Format(time.RFC3339),
			"status":        params.StatusCode,
			"latency_ms":    params.Latency.Milliseconds(),
			"client_ip":     params.ClientIP,
			"method":        params.Method,
			"path":          params.Path,
			"user_agent":    params.Request.UserAgent(),
			"error_message": params.ErrorMessage,
		}
		logJSON, err := json.Marshal(logEntry)
		if err != nil {
			return ""
		}
		return string(logJSON) + "\n"
	})
}

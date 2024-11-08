package middleware

import "net/http"

// Middleware type for chaining HTTP handlers
type Middleware func(http.HandlerFunc) http.HandlerFunc

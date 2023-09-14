package httpx

import "net/http"

// Doer represents a HTTP client.
type Doer interface {
	Do(req *http.Request) (resp *http.Response, error error)
}

// Is1xx reports whether code is in range [100, 200).
func Is1xx(code int) bool { return code >= 100 && code < 200 }

// Is2xx reports whether code is in range [200, 300).
func Is2xx(code int) bool { return code >= 200 && code < 300 }

// Is3xx reports whether code is in range [300, 400).
func Is3xx(code int) bool { return code >= 300 && code < 400 }

// Is4xx reports whether code is in range [400, 500).
func Is4xx(code int) bool { return code >= 400 && code < 500 }

// Is5xx reports whether code is in range [500, 600).
func Is5xx(code int) bool { return code >= 500 && code < 600 }

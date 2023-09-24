package httpx

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// NewRequest returns a new http.Request.
func NewRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, body)
}

// NewGetRequest returns a new http.Request with GET method.
func NewGetRequest(ctx context.Context, url string) (*http.Request, error) {
	return NewRequest(ctx, http.MethodGet, url, http.NoBody)
}

// NewHeadRequest returns a new http.Request with HEAD method.
func NewHeadRequest(ctx context.Context, url string) (*http.Request, error) {
	return NewRequest(ctx, http.MethodHead, url, http.NoBody)
}

// NewPostRequest returns a new http.Request with POST method.
func NewPostRequest(ctx context.Context, url string, body io.Reader) (*http.Request, error) {
	return NewRequest(ctx, http.MethodPost, url, body)
}

// NewPutRequest returns a new http.Request with PUT method.
func NewPutRequest(ctx context.Context, url string, body io.Reader) (*http.Request, error) {
	return NewRequest(ctx, http.MethodPut, url, body)
}

// NewPatchRequest returns a new http.Request with PATCH method.
func NewPatchRequest(ctx context.Context, url string, body io.Reader) (*http.Request, error) {
	return NewRequest(ctx, http.MethodPatch, url, body)
}

// NewDeleteRequest returns a new http.Request with DELETE method.
func NewDeleteRequest(ctx context.Context, url string, body io.Reader) (*http.Request, error) {
	return NewRequest(ctx, http.MethodDelete, url, body)
}

// NewConnectRequest returns a new http.Request with CONNECT method.
func NewConnectRequest(ctx context.Context, url string, body io.Reader) (*http.Request, error) {
	return NewRequest(ctx, http.MethodConnect, url, body)
}

// NewOptionsRequest returns a new http.Request with OPTIONS method.
func NewOptionsRequest(ctx context.Context, url string, body io.Reader) (*http.Request, error) {
	return NewRequest(ctx, http.MethodOptions, url, body)
}

// NewTraceRequest returns a new http.Request with TRACE method.
func NewTraceRequest(ctx context.Context, url string, body io.Reader) (*http.Request, error) {
	return NewRequest(ctx, http.MethodTrace, url, body)
}

// MustGetRequest returns a new http.Request with GET method.
func MustGetRequest(ctx context.Context, url string) *http.Request {
	req, err := NewRequest(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		panic(fmt.Sprintf("httpx: must make GET request: %v", err))
	}
	return req
}

// MustHeadRequest returns a new http.Request with HEAD method.
func MustHeadRequest(ctx context.Context, url string) *http.Request {
	req, err := NewRequest(ctx, http.MethodHead, url, http.NoBody)
	if err != nil {
		panic(fmt.Sprintf("httpx: must make HEAD request: %v", err))
	}
	return req
}

// MustPostRequest returns a new http.Request with POST method.
func MustPostRequest(ctx context.Context, url string, body io.Reader) *http.Request {
	req, err := NewRequest(ctx, http.MethodPost, url, body)
	if err != nil {
		panic(fmt.Sprintf("httpx: must make POST request: %v", err))
	}
	return req
}

// MustPutRequest returns a new http.Request with PUT method.
func MustPutRequest(ctx context.Context, url string, body io.Reader) *http.Request {
	req, err := NewRequest(ctx, http.MethodPut, url, body)
	if err != nil {
		panic(fmt.Sprintf("httpx: must make PUT request: %v", err))
	}
	return req
}

// MustPatchRequest returns a new http.Request with PATCH method.
func MustPatchRequest(ctx context.Context, url string, body io.Reader) *http.Request {
	req, err := NewRequest(ctx, http.MethodPatch, url, body)
	if err != nil {
		panic(fmt.Sprintf("httpx: must PATCH Patch request%v", err))
	}
	return req
}

// MustDeleteRequest returns a new http.Request with DELETE method.
func MustDeleteRequest(ctx context.Context, url string, body io.Reader) *http.Request {
	req, err := NewRequest(ctx, http.MethodDelete, url, body)
	if err != nil {
		panic(fmt.Sprintf("httpx: must DELETE Delete request%v", err))
	}
	return req
}

// MustConnectRequest returns a new http.Request with CONNECT method.
func MustConnectRequest(ctx context.Context, url string, body io.Reader) *http.Request {
	req, err := NewRequest(ctx, http.MethodConnect, url, body)
	if err != nil {
		panic(fmt.Sprintf("httpx: must CONNECT Connect request%v", err))
	}
	return req
}

// MustOptionsRequest returns a new http.Request with OPTIONS method.
func MustOptionsRequest(ctx context.Context, url string, body io.Reader) *http.Request {
	req, err := NewRequest(ctx, http.MethodOptions, url, body)
	if err != nil {
		panic(fmt.Sprintf("httpx: must OPTIONS Options request%v", err))
	}
	return req
}

// MustTraceRequest returns a new http.Request with TRACE method.
func MustTraceRequest(ctx context.Context, url string, body io.Reader) *http.Request {
	req, err := NewRequest(ctx, http.MethodTrace, url, body)
	if err != nil {
		panic(fmt.Sprintf("httpx: must TRACE Trace request%v", err))
	}
	return req
}

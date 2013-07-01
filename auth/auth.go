// Package auth provides an interface to check for authentication in an
// HTTP request.
package auth

import (
	"net/http"
)

type Authenticator interface {
	// Checks that the request is authenticated.
	// Returns `true` for authenticated requests, `false` otherwise.
	CheckAuth(*http.Request) bool
}

package down

import "fmt"

// Error is a download error.
type Error struct {
	url  string
	code int
}

func (e *Error) Error() string {
	return fmt.Sprintf("cannot download file %s, reason: HTTP %d", e.url, e.code)
}

func newError(url string, code int) error {
	return &Error{url, code}
}

package down

import (
	"io"
	"net/http"
)

// Downloader downloads an URL into a Blob.
type Downloader interface {
	Download(string) (Blob, error)
}

// Blob is a downloaded file. Either stored on a regular filesystem, remotely
// or in-memory.
type Blob interface {
	io.ReadCloser
	Location() string
	Release() error
}

// NewFileDownloader downloads an URL to a temporary file.
func NewFileDownloader(dir string, option ...*http.Client) Downloader {
	return &fileDownloader{dir, getOptionalHTTPClient(option...)}
}

// NewPipeDownloader downloads an URL on demand.
func NewPipeDownloader(option ...*http.Client) Downloader {
	return &pipeDownloader{getOptionalHTTPClient(option...)}
}

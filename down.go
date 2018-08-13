package down

import (
	"io"
	"net/http"
	"os"
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

// NewTempfileDownloader downloads an URL to a temporary file.
func NewTempfileDownloader(option ...*http.Client) Downloader {
	return &tempfileDownloader{os.TempDir(), getOptionalHTTPClient(option...)}
}

// NewPipeDownloader downloads an URL on demand.
func NewPipeDownloader(option ...*http.Client) Downloader {
	return &pipeDownloader{getOptionalHTTPClient(option...)}
}

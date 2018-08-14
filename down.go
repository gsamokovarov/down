package down

import (
	"io"
	"net/http"
)

// Downloader downloads an URL into a Blob.
type Downloader interface {
	Download(string) (Blob, error)
}

// Blob is a downloaded file. It can be physically stored on disk or streamed
// remotely as being read.
//
// Once done Read-ing the blob make sure to either Close or Release it. Close
// will close the file, while Release will clean it up (delete it) if its
// already stored.
type Blob interface {
	io.ReadCloser

	// Location is the location of the file. Can be an absolute path, can be a
	// remote URL or can be empty if it's not representable.
	Location() string

	// Release frees up the file. If it's stored on disk, this will mean removing
	// the file from it.
	Release() error
}

// NewFileDownloader downloads an URL to a file. A directory has to be
// specified for the file to be stored in. Can be used for temporary file
// downloads if dir is os.TempDir.
func NewFileDownloader(dir string, client ...*http.Client) Downloader {
	return &fileDownloader{dir, getOptionalHTTPClient(client...)}
}

// NewStreamDownloader downloads an URL on demand. This downloader won't save the
// file on disk, but will stream it directly as the blob is read. Can be used
// to download a remote file from one location and upload it to another remote
// one.
func NewStreamDownloader(client ...*http.Client) Downloader {
	return &streamDownloader{getOptionalHTTPClient(client...)}
}

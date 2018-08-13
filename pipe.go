package down

import (
	"fmt"
	"io"
	"net/http"

	"github.com/SuperLinearity/rhyme-app/be/shared/httputil/status"
)

type pipeDownloader struct {
	client *http.Client
}

func (t *pipeDownloader) Download(url string) (Blob, error) {
	response, err := t.client.Get(url)
	if err != nil {
		return nil, err
	}

	if code := response.StatusCode; code != status.OK {
		response.Body.Close()
		return nil, fmt.Errorf("cannot download file %s, reason: HTTP %d", url, code)
	}

	return &pipeBlob{url, response.Body}, nil
}

type pipeBlob struct {
	url string
	rc  io.ReadCloser
}

func (t *pipeBlob) Read(p []byte) (int, error) { return t.rc.Read(p) }
func (t *pipeBlob) Close() error               { return t.rc.Close() }
func (t *pipeBlob) Location() string           { return t.url }
func (t *pipeBlob) Release() error             { return t.rc.Close() }

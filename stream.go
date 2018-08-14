package down

import (
	"io"
	"net/http"
)

type streamDownloader struct {
	client *http.Client
}

func (t *streamDownloader) Download(url string) (Blob, error) {
	response, err := t.client.Get(url)
	if err != nil {
		return nil, err
	}

	if code := response.StatusCode; !successfulHTTP(code) {
		response.Body.Close()
		return nil, newError(url, code)
	}

	return &streamBlob{url, response.Body}, nil
}

type streamBlob struct {
	url string
	rc  io.ReadCloser
}

func (t *streamBlob) Read(p []byte) (int, error) { return t.rc.Read(p) }
func (t *streamBlob) Close() error               { return t.rc.Close() }
func (t *streamBlob) Location() string           { return t.url }
func (t *streamBlob) Release() error             { return t.rc.Close() }

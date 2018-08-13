package down

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/SuperLinearity/rhyme-app/be/shared/httputil/status"
)

type tempfileDownloader struct {
	tempDir string
	client  *http.Client
}

func (t *tempfileDownloader) Download(url string) (Blob, error) {
	location := path.Join(t.tempDir, path.Base(url))

	file, err := os.Create(location)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	response, err := t.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if code := response.StatusCode; code != status.OK {
		return nil, fmt.Errorf("cannot download file %s, reason: HTTP %d", url, code)
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return nil, err
	}

	return newTempfileBlob(location)
}

func newTempfileBlob(location string) (Blob, error) {
	file, err := os.Open(location)
	if err != nil {
		return nil, err
	}

	return &tempfileBlob{location, file}, nil
}

type tempfileBlob struct {
	path string
	file *os.File
}

func (t *tempfileBlob) Read(p []byte) (int, error) { return t.file.Read(p) }
func (t *tempfileBlob) Close() error               { return t.file.Close() }
func (t *tempfileBlob) Location() string           { return t.path }
func (t *tempfileBlob) Release() error {
	if err := t.file.Close(); err != nil {
		return err
	}

	return os.Remove(t.path)
}

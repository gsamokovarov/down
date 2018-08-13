package down

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/SuperLinearity/rhyme-app/be/shared/httputil/status"
)

type fileDownloader struct {
	dir    string
	client *http.Client
}

func (t *fileDownloader) Download(url string) (Blob, error) {
	location := path.Join(t.dir, path.Base(url))

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

	return newFileBlob(location)
}

func newFileBlob(location string) (Blob, error) {
	file, err := os.Open(location)
	if err != nil {
		return nil, err
	}

	return &fileBlob{location, file}, nil
}

type fileBlob struct {
	path string
	file *os.File
}

func (t *fileBlob) Read(p []byte) (int, error) { return t.file.Read(p) }
func (t *fileBlob) Close() error               { return t.file.Close() }
func (t *fileBlob) Location() string           { return t.path }
func (t *fileBlob) Release() error {
	if err := t.file.Close(); err != nil {
		return err
	}

	return os.Remove(t.path)
}

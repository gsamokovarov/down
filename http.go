package down

import "net/http"

func getOptionalHTTPClient(option ...*http.Client) *http.Client {
	if len(option) > 0 {
		return option[0]
	}

	return http.DefaultClient
}

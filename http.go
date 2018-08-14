package down

import "net/http"

func successfulHTTP(code int) bool {
	return code >= 200 && code < 300
}

func getOptionalHTTPClient(option ...*http.Client) *http.Client {
	if len(option) > 0 {
		return option[0]
	}

	return http.DefaultClient
}

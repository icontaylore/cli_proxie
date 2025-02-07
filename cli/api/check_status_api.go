package api

import "net/http"

func СheckStatus(resp *http.Response) bool {
	if resp.StatusCode == 200 {
		return true
	}
	return false
}

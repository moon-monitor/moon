package httpx

import (
	"fmt"
	"net/url"
)

func ParseQuery(qr map[string]any) string {
	if len(qr) == 0 {
		return ""
	}
	query := url.Values{}
	for k, v := range qr {
		query.Add(k, fmt.Sprintf("%v", v))
	}
	return query.Encode()
}

package client

import (
	"strings"
	"fmt"
)



// Query helps represent key value pairs used in url query parametes
type Query struct {
	Key   string
	Value string
}

// WithQuery lets you create a Query from key value pairs
func WithQuery(key string, value string) Query {
	return Query{
		Key:   key,
		Value: value,
	}
}

// AddQueryParamsToUrl lets you add query parameters to a url
func AddQueryParamsToUrl(url string, queries ...Query) string {
	for _, query := range queries {
		if strings.Contains(url, "?") {
			url += fmt.Sprintf("&%s=%s", query.Key, query.Value)
		} else {
			url += fmt.Sprintf("?%s=%s", query.Key, query.Value)
		}
	}
	return url
}
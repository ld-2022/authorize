package encoding

import (
	"net/http"
	"net/url"
)

func GetHeaderVal(header http.Header, name string) string {
	val := header.Get(name)
	if str, err := url.PathUnescape(val); err == nil {
		return str
	}
	return val
}

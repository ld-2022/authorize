package encoding

import (
	"net/http"
	"net/url"
	"testing"
)

func TestGetHeaderVal(t *testing.T) {
	header := http.Header{}
	header.Add("a", url.PathEscape("你好啊"))
	header.Add("b", url.PathEscape("你好啊/123"))
	header.Add("c", url.PathEscape("IT/123456"))
	header.Add("d", url.PathEscape("d:/123/世界你 好/aa"))
	//if header.Get("a") != "你好啊" {
	//	t.Error(header.Get("a"))
	//}
	if GetHeaderVal(header, "a") != "你好啊" {
		t.Error(GetHeaderVal(header, "a"))
	}
	if GetHeaderVal(header, "b") != "你好啊/123" {
		t.Error(GetHeaderVal(header, "b"))
	}
	if GetHeaderVal(header, "c") != "IT/123456" {
		t.Error(GetHeaderVal(header, "c"))
	}
	if GetHeaderVal(header, "d") != "d:/123/世界你 好/aa" {
		t.Error(GetHeaderVal(header, "d"))
	}

}

package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpMock(t *testing.T) {
	f := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("content-type", "application/json")
		w.Write([]byte("{\"ret\":true, \"data\":\"hello mock\"}"))
	}

	mockServer := httptest.NewServer(http.HandlerFunc(f))

	defer mockServer.Close()

	resp, _ := http.Get(mockServer.URL)

	fmt.Println(string(readManually(resp.Body)))
}

func readManually(r io.ReadCloser) []byte {
	result := []byte{}
	tmp := [10]byte{}
	sl := tmp[0:10]

	for {
		read, err := r.Read(sl)
		result = append(result, sl[0:read]...)
		if err == io.EOF {
			break
		}
	}

	return result
}

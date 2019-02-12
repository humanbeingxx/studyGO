package test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(os.Stdout, resp.Body)
}

func TestBuffer(t *testing.T) {
	var buffer bytes.Buffer
	buffer.Write([]byte("test for buffer\n"))
	writer := io.MultiWriter(os.Stdout, os.Stdout)
	io.Copy(writer, &buffer)
}

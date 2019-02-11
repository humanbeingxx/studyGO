package test

import (
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

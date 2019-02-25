package test

import (
	"io"
	"os"
	"testing"
)

func TestAppendFile(t *testing.T) {
	file, err := os.OpenFile("/Users/cxs/temp/test_write.sql", os.O_RDWR|os.O_CREATE|os.O_APPEND, 755)
	if err != nil {
		panic(err.Error)
	}
	// ioutil.WriteFile("/Users/cxs/temp/test_write.sql", ([]byte)("test\n"), 666)
	io.WriteString(file, "test\n")
	file.Close()
}

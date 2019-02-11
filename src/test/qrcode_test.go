package test

import (
	"log"
	"testing"

	qrcode "github.com/skip2/go-qrcode"
)

func TestGenerate(t *testing.T) {
	err := qrcode.WriteFile("www.baidu.com", qrcode.Medium, 256, "/tmp/qr.png")
	if err != nil {
		log.Fatal("生成二维码失败", err)
	}
}

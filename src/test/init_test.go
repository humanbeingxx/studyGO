package test

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())
	runtime.GOMAXPROCS(1)

	fmt.Print("init finished\n\n")

	m.Run()
}

package test

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

var counter int32
var waitGroup sync.WaitGroup
var mutex sync.Mutex

func TestNoAtomic(t *testing.T) {
	waitGroup.Add(2)

	go incrementUnsafe()
	go incrementUnsafe()
	waitGroup.Wait()

	convey.Convey("TestNoAtomic", t, func() {
		convey.So(counter, convey.ShouldEqual, 3)
	})
}

func incrementUnsafe() {
	defer waitGroup.Done()

	time.Sleep(10 * time.Millisecond)

	for index := 0; index < 3; index++ {
		value := counter

		runtime.Gosched()

		value++
		counter = value
	}
}

func TestAtomic(t *testing.T) {
	waitGroup.Add(2)

	go incrementWithAtomic()
	go incrementWithAtomic()
	waitGroup.Wait()

	convey.Convey("TestAtomic", t, func() {
		convey.So(counter, convey.ShouldEqual, 6)
	})
}

func incrementWithAtomic() {
	defer waitGroup.Done()

	for index := 0; index < 3; index++ {
		atomic.AddInt32(&counter, 1)
		runtime.Gosched()
	}
}

func TestMutex(t *testing.T) {
	waitGroup.Add(2)

	go incrementWithMutex()
	go incrementWithMutex()
	waitGroup.Wait()

	convey.Convey("TestMutex", t, func() {
		convey.So(counter, convey.ShouldEqual, 6)
	})
}

func incrementWithMutex() {
	defer waitGroup.Done()

	for index := 0; index < 3; index++ {

		mutex.Lock()
		value := counter

		runtime.Gosched()

		value++
		counter = value
		mutex.Unlock()
	}
}

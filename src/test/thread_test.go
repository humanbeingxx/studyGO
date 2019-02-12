package test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

func TestSleep(t *testing.T) {
	fmt.Println(time.Now())
	time.Sleep(time.Second)
	fmt.Println(time.Now())
}

func TestMulti(t *testing.T) {
	go sleep()
	go sleep()
	go sleep()
	time.Sleep(2 * time.Second)
}

func sleep() {
	fmt.Println("ready to sleep")
	time.Sleep(time.Second)
}

func TestRoutineFunc(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}

	for _, value := range data {
		go func(value int) {
			fmt.Println(value)
		}(value)
	}
}

func TestWait(t *testing.T) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	results := make(chan *string)

	go func() {
		fmt.Println("worker1 working......")
		time.Sleep(time.Second)
		result := "result1"
		results <- &result
		waitGroup.Done()
	}()
	go func() {
		fmt.Println("worker2 working......")
		time.Sleep(time.Second)
		result := "result2"
		results <- &result
		waitGroup.Done()
	}()
	go func() {
		fmt.Println("worker3 working......")
		time.Sleep(time.Second)
		result := "result3"
		results <- &result
		waitGroup.Done()
	}()

	go func() {
		waitGroup.Wait()
		fmt.Println("all finished")
	}()

	for value := range results {
		fmt.Println(*value)
	}

}

func TestNoBufferChannel(t *testing.T) {
	noBufferChan := make(chan int)

	sent := false

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		noBufferChan <- 1
		sent = true
	}()

	go func() {
		defer wg.Done()

		convey.Convey("before receive", t, func() {
			convey.So(sent, convey.ShouldNotEqual, true)
		})
		value := <-noBufferChan
		fmt.Printf("received %d\n", value)

		time.Sleep(100 * time.Millisecond)

		convey.Convey("after receive", t, func() {
			convey.So(sent, convey.ShouldEqual, true)
		})
	}()

	wg.Wait()
}

func TestBufferChannel(t *testing.T) {
	bufferChan := make(chan int, 3)

	var wg sync.WaitGroup
	sent := false

	wg.Add(2)

	go func() {
		defer wg.Done()

		bufferChan <- 1
		bufferChan <- 2
		bufferChan <- 3
		sent = true
		close(bufferChan)
	}()

	time.Sleep(100 * time.Millisecond)

	go func() {
		defer wg.Done()

		convey.Convey("after receive", t, func() {
			convey.So(sent, convey.ShouldEqual, true)
		})

		for value, ok := <-bufferChan; ok; value, ok = <-bufferChan {
			fmt.Printf("received %d\n", value)
		}
	}()

	wg.Wait()
}

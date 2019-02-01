package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
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

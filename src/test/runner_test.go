package test

import (
	"fmt"
	"studyGO/src/service/runner"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	myRunner := runner.New(3 * time.Second)
	myRunner.AddTask(func() {
		for index := 0; index < 5; index++ {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("worker A is working")
		}
	})

	myRunner.AddTask(func() {
		for index := 0; index < 5; index++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("worker B is working")
		}
	})

	myRunner.AddTask(func() {
		for index := 0; index < 5; index++ {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("worker C is working")
		}
	})

	if err := myRunner.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			fmt.Println("Oh, timeout")
			return
		// ctrl + c,只支持在启动task时响应中断，启动后的任务无法中断
		case runner.ErrInterrupt:
			fmt.Println("Oh, interrupted")
			return
		}
	}

	time.Sleep(3 * time.Second)

	fmt.Println("work complete")
}

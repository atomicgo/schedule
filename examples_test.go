package schedule_test

import (
	"fmt"
	"time"

	"atomicgo.dev/schedule"
)

func ExampleAfter() {
	task := schedule.After(5*time.Second, func() {
		fmt.Println("5 seconds are over!")
	})

	fmt.Println("Some stuff happening...")

	task.Wait()
}

func ExampleAt() {
	task := schedule.At(time.Now().Add(5*time.Second), func() {
		fmt.Println("5 seconds are over!")
	})

	fmt.Println("Some stuff happening...")

	task.Wait()
}

func ExampleEvery() {
	task := schedule.Every(time.Second, func() bool {
		fmt.Println("1 second is over!")

		return true // return false to stop the task
	})

	fmt.Println("Some stuff happening...")

	time.Sleep(10 * time.Second)

	task.Stop()
}

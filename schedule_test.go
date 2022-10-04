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
	task := schedule.Every(5*time.Second, func() {
		fmt.Println("5 seconds are over!")
	})

	fmt.Println("Some stuff happening...")

	time.Sleep(10 * time.Second)

	task.Stop()
}

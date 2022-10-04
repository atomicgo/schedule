package schedule

import (
	"fmt"
	"time"
)

func ExampleAfter() {
	task := After(5*time.Second, func() {
		fmt.Println("5 seconds are over!")
	})

	fmt.Println("Some stuff happening...")

	task.Wait()
}

func ExampleAt() {
	task := At(time.Now().Add(5*time.Second), func() {
		fmt.Println("5 seconds are over!")
	})

	fmt.Println("Some stuff happening...")

	task.Wait()
}

func ExampleEvery() {
	task := Every(5*time.Second, func() {
		fmt.Println("5 seconds are over!")
	})

	fmt.Println("Some stuff happening...")

	time.Sleep(10 * time.Second)

	task.Stop()
}

package schedule

import "testing"

func TestHelloWorld(t *testing.T) {
	if HelloWorld() != "Hello, World!" {
		t.Fatal("Not equal")
	}
}

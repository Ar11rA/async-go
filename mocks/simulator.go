package mocks

import (
	"time"
)

func AsyncA(done chan string) {
	time.Sleep(20000)
	done <- "A"
}

func AsyncB(done chan string) {
	time.Sleep(20000)
	done <- "B"
}

func SyncA() string {
	time.Sleep(20000)
	return "A"
}

func SyncB() string {
	time.Sleep(20000)
	return "B"
}
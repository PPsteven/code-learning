package main

import (
	"testing"
	"time"
)

// understanding t.Parallel
// TestA and TestB will run the same time
// go test -parallel 2
func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
}

func TestB(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
}
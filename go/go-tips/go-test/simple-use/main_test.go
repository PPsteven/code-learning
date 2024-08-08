package main

import (
	"fmt"
	"testing"
	"time"
)

// Simple Testing
func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Log("log: add function success") // 仅失败 或 -v 输出
		t.FailNow()
		//t.Error() // Log -> Fail
		//t.Fatal() // Log -> FailNow
	}
}

// table driven test
func TestAdd_table_driven(t *testing.T) {
	fmt.Printf("start test TestAdd2")
	var test = []struct{
		x, y, expect int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{5, 3, 8},
	}

	for _, tt := range test {
		if actual := Add(tt.x, tt.y); actual != tt.expect {
			t.Errorf("add(%d, %d), expect: %d, actual:%d", tt.x, tt.y, tt.expect, actual)
		}
	}
}

// skip long runing test
// go test -short
func TestAdd_longtest(t *testing.T) {
	fmt.Printf("start test TestAdd3")
	if testing.Short() {
		t.Skip("Skip long runing test")
	}
	time.Sleep(100 * time.Second)
}
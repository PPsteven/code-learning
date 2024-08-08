package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestAdd_assert
func TestAdd_assert(t *testing.T) {
	//if Add(1, 2) != 3 {
	//	t.Log("log: add function success") // 仅失败 或 -v 输出
	//	t.FailNow()
	//}
	assert.Equal(t, Add(1, 2), 3)
}
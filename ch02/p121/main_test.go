package main

import (
	"testing"

	"github.com/youyouryu/ant-book/common"
)

func TestSolve1(t *testing.T) {
	common.Check(solve, "test/q1.txt", "test/a1.txt")
}

func TestSolve2(t *testing.T) {
	common.Check(solve, "test/q2.txt", "test/a2.txt")
}
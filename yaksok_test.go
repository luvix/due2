package main

import (
	"os"
	"testing"
)

// main 함수를 테스트 할 수 있는 함수는 이것 뿐이다.
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestShakeFlag(t *testing.T) {
	yaksok := YaksokFactory()
	yaksok.ShakeFlag([]string{"h"})
	yaksok.ShakeFlag([]string{"v"})
}

func TestShakeSubsetFlag(t *testing.T) {
	yaksok := YaksokFactory()
	yaksok.ShakeSubsetFlag([]string{"secondly", "echo \"hello world\""})
}

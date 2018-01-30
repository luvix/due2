package main

import (
	"os"
	"testing"
)

// main 함수를 테스트 할 수 있는 함수는 이것 뿐이다.
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestYaksokFactory(t *testing.T) {
	ys := YaksokFactory()

	if ys.once == nil {
		t.Error()
	}
	if ys.daily == nil {
		t.Error()
	}
	if ys.weekly == nil {
		t.Error()
	}
	if ys.monthly == nil {
		t.Error()
	}
	if ys.yearly == nil {
		t.Error()
	}
	if ys.secondly == nil {
		t.Error()
	}
	if ys.minutely == nil {
		t.Error()
	}
	if ys.hourly == nil {
		t.Error()
	}
	if ys.list == nil {
		t.Error()
	}
	if ys.minutely == nil {
		t.Error()
	}
	if ys.setting == nil {
		t.Error()
	}

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

package main

import (
	"os"
	"testing"
)

// main 함수를 테스트 할 수 있는 함수는 이것 뿐이다.
//그런데 놀랍게도 이 함수는 flag.Parse를 호출하지 않기때문에 쓸모없다.
func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestMainFlagBoxFactory(t *testing.T) {
	box := MainFlagBoxFactory()

	if box == nil {
		t.Error("New box fail")
	}
}

func TestSubFlagBoxFactory(t *testing.T) {
	box := SubFlagBoxFactory()

	if box == nil {
		t.Error("New box fail")
	}
}

func TestFlagBoxFactory(t *testing.T) {
	box := FlagBoxFactory()

	if box == nil {
		t.Error("New box fail")
	}
}

func TestFlagBoxPickup(t *testing.T) {
	box := FlagBoxFactory()
	box.Pickup(nil)
}

func TestMainFlagBoxPickup(t *testing.T) {
	box := MainFlagBoxFactory()
	box.Pickup()
}

func TestSubFlagBoxPickup(t *testing.T) {
	box := SubFlagBoxFactory()

	box.Pickup(nil)
}

package main

import (
	"testing"
)

// main 함수를 테스트 할 수 있는 함수는 이것 뿐이다.
//그런데 놀랍게도 이 함수는 flag.Parse를 호출하지 않기때문에 쓸모없다.
// func TestMain(m *testing.M) {
// 	Usage() //Print usage of testing
// 	os.Exit(m.Run())
// }

func TestNewFlagSet(t *testing.T) {
	fs1 := NewFlagSet("once")
	fs2 := NewFlagSet("twice")
	fs1.JobName()
	fs2.JobName()
}

func TestNewYaksokFlagBox(t *testing.T) {
	box := NewYaksokFlagBox()

	if box == nil {
		t.Error("New box fail")
	} else {
		if box.help != nil {
			t.Error("YaksokFlagSet::help is not empty be4 alloc")
		}
		if box.version != nil {
			t.Error("YaksokFlagSet::version is not empty be4 alloc")
		}

		flagBox := Ready2FlagBox()
		box := flagBox.yaksokBox
		if *box.help {
			t.Error("YaksokFlagSet::help is empty in box after Ready2FlagBox")
		}
		if *box.version {
			t.Error("YaksokFlagSet::version is empty in box after Ready2FlagBox")
		}
	}
}

func TestNewSubFlagBox(t *testing.T) {
	box := NewSubFlagBox()

	if box == nil {
		t.Error("New box fail")
	}

	if box.once == nil {
		t.Error("No once")
	}
}

func TestNewFlagBox(t *testing.T) {
	box := NewFlagBox()

	if box == nil {
		t.Error("New box fail")
	}
}

func TestFlagBoxPickup(t *testing.T) {
	box := NewFlagBox()
	box.Pickup(nil)
}

func TestMainFlagBoxPickup(t *testing.T) {
	box := NewYaksokFlagBox()
	box.Pickup()
}

func TestSubFlagBoxPickup(t *testing.T) {
	box := NewSubFlagBox()
	box.Pickup(nil)
}

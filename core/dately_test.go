package core

import (
	"flag"
	"testing"
)

func TestNewTimelyFlagSet(t *testing.T) {
	tfs := NewTimelyFlagSet("test", flag.ExitOnError)

	if tfs == nil {
		t.Error()
	}
}

func TestTimelyYaksokTags(T *testing.T) {

}

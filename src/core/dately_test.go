package core

import (
	"flag"
	"testing"
)

func TestNewTimelyFlagSet(t *testing.T) {
	tfs := NewTimelyFlagSet("test", flag.ExitOnError)
}

// func TestTimelyYaksokTags(T *testing.T) {
	
// }

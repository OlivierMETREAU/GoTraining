package orm

import (
	"testing"
)

func TestReflexion(t *testing.T) {
	s := New("John", 42)
	s.Inspect()
}

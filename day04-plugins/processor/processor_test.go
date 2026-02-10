package processor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUppercase(t *testing.T) {
	assert.Equal(t, "TO UPPER CASE", Uppercase{}.Process("To uPper caSe"))
}

func TestLowercase(t *testing.T) {
	assert.Equal(t, "to lower case", Lowercase{}.Process("To LowEr caSe"))
}

func TestRot13(t *testing.T) {
	assert.Equal(t, "Pack My Box With Five Dozen Liquor Jugs", Rot13{}.Process("Cnpx Zl Obk Jvgu Svir Qbmra Yvdhbe Whtf"))
}

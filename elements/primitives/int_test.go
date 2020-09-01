package primitives_test

import (
	"github.com/christiangelone/collgo/elements/primitives"
	"github.com/realistschuckle/testify/assert"
	"testing"
)

func TestConvertingIntIntoString(t *testing.T) {
	anInt := primitives.NewInt(1)
	aString := anInt.ToString()
	assert.Equal(t, "1", aString.Value())
}
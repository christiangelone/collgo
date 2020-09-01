package primitives_test

import (
	"github.com/christiangelone/collgo/elements/primitives"
	"github.com/realistschuckle/testify/assert"
	"testing"
)

func TestConvertingStringIntoInt(t *testing.T) {
	aString := primitives.NewString("1")
	anInt := aString.ToInt()
	assert.Equal(t, 1, anInt.Value())
}
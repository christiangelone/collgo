package primitives_test

import (
	"github.com/christiangelone/collgo/elements/primitives"
	"github.com/stretchr/testify/assert"
	"testing"
)

type _Struct struct {
	value int
}

func TestConvertingStructIntoString(t *testing.T) {
	aStruct := primitives.NewStruct(_Struct{ value: 1 })
	aString := aStruct.ToString()
	assert.NotEqual(t, "", aString)
}

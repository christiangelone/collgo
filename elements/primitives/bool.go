package primitives

import (
	"strconv"
)

type Bool struct {
	value bool
}

func NewBool(aBool bool) *Bool {
	return &Bool{
		value: aBool,
	}
}

func (b *Bool) Value() Any {
	return b.value
}

func (b *Bool) ToString() *String {
	aString := strconv.FormatBool(b.value)
	return NewString(aString)
}

func (b *Bool) ToInt() *Int {
	if b.value {
		return NewInt(1)
	}
	return NewInt(0)
}

func (b *Bool) ToDouble() *Double {
	if b.value {
		return NewDouble(1.0)
	}
	return NewDouble(0.0)
}

func (b *Bool) ToFloat() *Float {
	if b.value {
		return NewFloat(1.0)
	}
	return NewFloat(0.0)
}

func (b *Bool) ToBool() *Bool {
	return b
}

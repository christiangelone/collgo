package primitives

import (
	"strconv"
)

type Int struct {
	value int
}

func NewInt(anInt int) *Int {
	return &Int{
		value: anInt,
	}
}

func (i *Int) Value() Any {
	return i.value
}

func (i *Int) ToString() *String {
	aString := strconv.Itoa(i.value)
	return NewString(aString)
}

func (i *Int) ToInt() *Int {
	return i
}

func (i *Int) ToDouble() *Double {
	return NewDouble(float64(i.value))
}

func (i *Int) ToFloat() *Float {
	return NewFloat(float32(i.value))
}

func (i *Int) ToBool() *Bool {
	if i.value > 0 {
		return NewBool(true)
	}
	return NewBool(false)
}

package primitives

import (
	"strconv"
)

type Float struct {
	value float32
}

func NewFloat(aFloat float32) *Float {
	return &Float{
		value: aFloat,
	}
}

func (f *Float) Value() Any {
	return f.value
}

func (f *Float) ToString() *String {
	aString := strconv.FormatFloat(float64(f.value), 'E', -1, 32)
	return NewString(aString)
}

func (f *Float) ToInt() *Int {
	return NewInt(int(f.value))
}

func (f *Float) ToDouble() *Double {
	return NewDouble(float64(f.value))
}

func (f *Float) ToFloat() *Float {
	return f
}

func (f *Float) ToBool() *Bool {
	if f.value > 0 {
		return NewBool(true)
	}
	return NewBool(false)
}

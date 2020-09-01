package primitives

import (
	"strconv"
)

type Double struct {
	value float64
}

func NewDouble(aDouble float64) *Double {
	return &Double{
		value: aDouble,
	}
}

func (d *Double) Value() Any {
	return d.value
}

func (d *Double) ToString() *String {
	aString := strconv.FormatFloat(d.value, 'E', -1, 64)
	return NewString(aString)
}

func (d *Double) ToInt() *Int {
	return NewInt(int(d.value))
}

func (d *Double) ToDouble() *Double {
	return d
}

func (d *Double) ToFloat() *Float {
	return NewFloat(float32(d.value))
}

func (d *Double) ToBool() *Bool {
	if d.value > 0 {
		return NewBool(true)
	}
	return NewBool(false)
}

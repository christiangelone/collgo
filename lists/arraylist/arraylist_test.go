package arraylist_test

import (
	"github.com/christiangelone/collgo/elements"
	"github.com/christiangelone/collgo/elements/primitives"
	"github.com/christiangelone/collgo/lists/arraylist"
	"testing"

	"github.com/realistschuckle/testify/assert"
)

func TestRemovingElementsInArrayList(t *testing.T) {
	list := arraylist.New(0)
	list.Add(primitives.NewInt(1))
	assert.False(t, list.IsEmpty())
	list.RemoveAt(0)
	assert.True(t, list.IsEmpty())
}

func TestAddingElementsInArrayList(t *testing.T) {
	list := arraylist.New(0)
	assert.Equal(t, uint64(0), list.Size())
	list.Add(primitives.NewInt(1))
	assert.Equal(t, uint64(1), list.Size())
	element, err := list.GetAt(0)
	assert.Nil(t, err)
	assert.Equal(t, 1, element.Value())
}

func TestPanicWhenAddingADifferentTypeInArrayList(t *testing.T) {
	list := arraylist.New(0)
	list.Add(primitives.NewInt(1))

	defer func() { recover() }()
	list.Add(primitives.NewString("asd"))
}

func TestIterateWithForEachInArrayList(t *testing.T) {
	list := arraylist.New(0)
	list.Add(primitives.NewInt(0))
	list.Add(primitives.NewInt(1))
	list.Add(primitives.NewInt(2))

	i := 0
	list.ForEach(func(element elements.IElement) {
		assert.Equal(t, i, element.Value())
		i++
	})
}

func TestHasElementInArrayList(t *testing.T) {
	list := arraylist.New(0)
	list.Add(primitives.NewStruct(struct{ Value int}{ Value: 0 }))

	hasElement := list.Has(primitives.NewStruct(struct{ Value int}{ Value: 0 }))
	assert.True(t, hasElement)
}
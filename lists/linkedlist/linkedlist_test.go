package linkedlist_test

import (
	"github.com/christiangelone/collgo/elements"
	"github.com/christiangelone/collgo/elements/primitives"
	"github.com/christiangelone/collgo/lists"
	"github.com/christiangelone/collgo/lists/linkedlist"
	"strconv"
	"testing"

	"github.com/realistschuckle/testify/assert"
)

func TestAddingInLinkedList(t *testing.T) {
	list := linkedlist.New()
	assert.Equal(t, uint64(0), list.Size())

	list.Add(primitives.NewInt(1))
	assert.Equal(t, uint64(1), list.Size())
	element1, _ :=  list.Last()
	assert.Equal(t, 1, element1.Value())

	list.Add(primitives.NewInt(2))
	assert.Equal(t, uint64(2), list.Size())
	element2, _ :=  list.Last()
	assert.Equal(t, 2, element2.Value())

	list.Add(primitives.NewInt(3))
	assert.Equal(t, uint64(3), list.Size())
	element3, _ := list.Last()
	assert.Equal(t, 3, element3.Value())
}

func TestPushAndPopInLinkedList(t *testing.T) {
	var stack lists.IStack = linkedlist.New()
	stack.Push(primitives.NewInt(0))
	stack.Push(primitives.NewInt(1))
	stack.Push(primitives.NewInt(2))

	elem1, _ := stack.Pop()
	elem2, _ := stack.Pop()
	elem3, _ := stack.Pop()

	assert.Equal(t, uint64(0), stack.Size())
	assert.Equal(t, 2, elem1.Value())
	assert.Equal(t, 1, elem2.Value())
	assert.Equal(t, 0, elem3.Value())
}

func TestEnqueueDequeueInLinkedList(t *testing.T) {
	var queue lists.IQueue = linkedlist.New()
	queue.Enqueue(primitives.NewInt(0))
	queue.Enqueue(primitives.NewInt(1))
	queue.Enqueue(primitives.NewInt(2))

	elem1, _ := queue.Dequeue()
	elem2, _ := queue.Dequeue()
	elem3, _ := queue.Dequeue()

	assert.Equal(t, uint64(0), queue.Size())
	assert.Equal(t, 0, elem1.Value())
	assert.Equal(t, 1, elem2.Value())
	assert.Equal(t, 2, elem3.Value())
}

func TestIterateWithIteratorInLinkedList(t *testing.T) {
	list := linkedlist.New()
	list.Add(primitives.NewInt(0))
	list.Add(primitives.NewInt(1))
	list.Add(primitives.NewInt(2))

	iterator := list.GetRightIterator()

	i := 0
	for element, end := iterator.First(); !end; element, end = iterator.Next() {
		assert.Equal(t, i, element.Value())
		i++
	}
}

func TestIterateWithForEachInLinkedList(t *testing.T) {
	list := linkedlist.New()
	list.Add(primitives.NewInt(0))
	list.Add(primitives.NewInt(1))
	list.Add(primitives.NewInt(2))
	
	i := 0
	list.ForEach(func(element elements.IElement) {
		assert.Equal(t, i, element.Value())
		i++
	})
}

func TestFilterInLinkedList(t *testing.T) {
	list := linkedlist.New()
	list.Add(primitives.NewInt(1))
	list.Add(primitives.NewInt(2))
	list.Add(primitives.NewInt(3))

	filteredList := list.Filter(func(element elements.IElement) bool {
		return element.Value() != 3
	})

	assert.Equal(t, uint64(2), filteredList.Size())
}

func TestMapInLinkedList(t *testing.T) {
	list := linkedlist.New()
	list.Add(primitives.NewInt(1))
	list.Add(primitives.NewInt(2))
	list.Add(primitives.NewInt(3))

	mappedList := list.Map(func(element elements.IElement) elements.IElement {
		return primitives.NewString(strconv.Itoa(element.Value().(int)))
	})

	element1 , _ := mappedList.(lists.IQueue).Dequeue()
	element2, _ := mappedList.(lists.IQueue).Dequeue()
	element3, _ := mappedList.(lists.IQueue).Dequeue()
	assert.Equal(t, "1", element1.Value())
	assert.Equal(t, "2", element2.Value())
	assert.Equal(t, "3", element3.Value())
}

func TestRightReduceInLinkedList(t *testing.T) {
	list := linkedlist.New()
	list.Add(primitives.NewString("A"))
	list.Add(primitives.NewString("B"))
	list.Add(primitives.NewString("C"))

	element := list.RightReduce(func(reduced elements.IElement, element elements.IElement) elements.IElement {
		return primitives.NewString(reduced.Value().(string) + element.Value().(string))
	})
	assert.Equal(t, "ABC", element.Value())
}

func TestLeftReduceInLinkedList(t *testing.T) {
	list := linkedlist.New()
	list.Add(primitives.NewString("A"))
	list.Add(primitives.NewString("B"))
	list.Add(primitives.NewString("C"))

	element := list.LeftReduce(func(reduced elements.IElement, element elements.IElement) elements.IElement {
		return primitives.NewString(reduced.Value().(string) + element.Value().(string))
	})
	assert.Equal(t, "CBA", element.Value())
}

func TestZipInLinkedList(t *testing.T) {
	listA := linkedlist.New()
	listA.Add(primitives.NewString("A"))
	listA.Add(primitives.NewString("B"))
	listA.Add(primitives.NewString("C"))

	listB := linkedlist.New()
	listB.Add(primitives.NewString("D"))
	listB.Add(primitives.NewString("E"))
	listB.Add(primitives.NewString("F"))

	listC := listA.Zip(listB)
	listC.ForEach(func(element elements.IElement) {
		elementA, _ := listA.Dequeue()
		elementB, _ := listB.Dequeue()
		assert.Equal(
			t,
			[]interface{}{elementA.Value(), elementB.Value()},
			element.Value(),
		)
	})
}

func TestZipWithDifferentSizesInLinkedList(t *testing.T) {
	listA := linkedlist.New()
	listA.Add(primitives.NewString("A"))
	listA.Add(primitives.NewString("B"))
	listA.Add(primitives.NewString("C"))

	listB := linkedlist.New()
	listB.Add(primitives.NewString("D"))
	listB.Add(primitives.NewString("E"))

	defer func() { recover() }()
	listA.Zip(listB)
}

func TestAdjustedZipInLinkedList(t *testing.T) {
	listA := linkedlist.New()
	listA.Add(primitives.NewString("A"))
	listA.Add(primitives.NewString("B"))
	listA.Add(primitives.NewString("C"))

	listB := linkedlist.New()
	listB.Add(primitives.NewString("D"))
	listB.Add(primitives.NewString("E"))

	listC := listA.AdjustedZip(listB, nil)
	elementA, _ := listC.(lists.IQueue).Dequeue()
	assert.Equal(t, []interface{}{"A", "D"}, elementA.Value())
	elementB, _ := listC.(lists.IQueue).Dequeue()
	assert.Equal(t, []interface{}{"B", "E"}, elementB.Value())
	elementC, _ := listC.(lists.IQueue).Dequeue()
	assert.Equal(t, []interface{}{"C", nil}, elementC.Value())
}

func TestIsSortedInLinkedList(t *testing.T) {
	list := linkedlist.New()
	list.Add(primitives.NewInt(1))
	list.Add(primitives.NewInt(2))
	list.Add(primitives.NewInt(2))
	list.Add(primitives.NewInt(3))

	assert.True(t, list.IsSorted(func(a elements.IElement, b elements.IElement) bool {
		return a.Value().(int) <= b.Value().(int)
	}))
}

func TestSortInLinkedList(t *testing.T) {
	list := linkedlist.New()
	list.Add(primitives.NewInt(16))
	list.Add(primitives.NewInt(14))
	list.Add(primitives.NewInt(12))
	list.Add(primitives.NewInt(10))

	sortedList := list.Sort(func(a elements.IElement, b elements.IElement) bool {
		return a.Value().(int) <= b.Value().(int)
	})

	i := 10
	sortedList.ForEach(func(element elements.IElement) {
		assert.Equal(t, i, element.Value())
		i+=2
	})
}
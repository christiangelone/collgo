package arraylist

import (
	"github.com/christiangelone/collgo/elements"
	"github.com/christiangelone/collgo/lists"
	"reflect"
)

func New(capacity uint64) *ArrayList {
	return &ArrayList{
		size:     0,
		threshold: 0.7 * float64(capacity),
		elements: make([]elements.IElement, 0, capacity),
	}
}

type ArrayList struct {
	size       uint64
	threshold  float64
	elements   []elements.IElement
}

func (l *ArrayList) AdjustedZip(list lists.IList, fillValue interface{}) lists.IList {
	panic("implement me")
}

func (l *ArrayList) Zip(collection lists.IList) lists.IList {
	panic("implement me")
}

func (l *ArrayList) Sort(fn lists.CompareFn) lists.IList {
	panic("implement me")
}

func (l *ArrayList) IsSorted(lists.CompareFn) bool {
	panic("implement me")
}

func (l *ArrayList) Map(mapFn lists.MapFn) lists.IList {
	newList := New(uint64(cap(l.elements)))
	l.ForEach(func(element elements.IElement) {
		newList.Add(mapFn(element))
	})
	return newList
}

func (l *ArrayList) Filter(filterFn lists.FilterFn) lists.IList {
	newList := New(uint64(cap(l.elements)))
	l.ForEach(func(element elements.IElement) {
		if filterFn(element) {
			newList.Add(element)
		}
	})
	return newList
}

func (l *ArrayList) LeftReduce(fn lists.ReduceFn) elements.IElement {
	panic("implement me")
}

func (l *ArrayList) LeftReduceWith(fn lists.ReduceFn, element elements.IElement) elements.IElement {
	panic("implement me")
}

func (l *ArrayList) RightReduce(fn lists.ReduceFn) elements.IElement {
	panic("implement me")
}

func (l *ArrayList) RightReduceWith(fn lists.ReduceFn, element elements.IElement) elements.IElement {
	panic("implement me")
}

func (l *ArrayList) Has(element elements.IElement) bool {
	iterator := l.GetRightIterator()
	for currentElement, end := iterator.First(); !end; currentElement, end = iterator.Next() {
		if currentElement.Value() == element.Value() {
			return true
		}
	}
	return false
}

func (l *ArrayList) ForEach(eachFn lists.EachFn) {
	iterator := l.GetRightIterator()
	for elem, end := iterator.First(); !end; elem, end = iterator.Next() {
		eachFn(elem)
	}
}

func (l *ArrayList) GetRightIterator() lists.IIterator {
	return &rightIterator{
		firstIndex: 0,
		currentIndex: 0,
		list:  l,
	}
}

func (l *ArrayList) GetLeftIterator() lists.IIterator {
	return &leftIterator{
		firstIndex: l.Size() - 1,
		currentIndex: l.Size() - 1,
		list:  l,
	}
}

func (l *ArrayList) IsEmpty() bool {
	return l.size == 0
}

func (l *ArrayList) Size() uint64 {
	return l.size
}

func (l *ArrayList) GetAt(index uint64) (elements.IElement, error) {
	if l.IsEmpty() {
		return nil, &lists.FetchingInEmptinessError{}
	}
	if index >= 0 && index < l.size {
		return l.elements[index], nil
	}
	return nil, &OutOfBoundError{index: index}
}

func (l *ArrayList) Add(element elements.IElement) {
	first, _ := l.First()
	if l.IsEmpty() || reflect.TypeOf(first.Value()) == reflect.TypeOf(element.Value()) {
		l.size++
		l.elements = append(l.elements, element)
		if float64(l.Size()) >= l.threshold {
			newElements := make([]elements.IElement, l.Size(), 2 * cap(l.elements))
			copy(newElements, l.elements)
			l.elements = newElements
		}
	} else {
		first, _ := l.First()
		panic(&lists.ElementTypeError{
			Expected: first,
			Actual:   element,
		})
	}
}

func (l *ArrayList) RemoveAt(index uint64) {
	if !l.IsEmpty() && index >= 0 && index < l.size {
		l.elements[index] = l.elements[l.size-1]
		l.elements[l.size-1] = nil
		l.elements = l.elements[:l.size-1]
		l.size--
	}
}

func (l *ArrayList) First() (elements.IElement, error) {
	if l.IsEmpty() {
		return nil, &lists.FetchingInEmptinessError{}
	}
	return l.elements[0], nil
}

func (l *ArrayList) Last() (elements.IElement, error) {
	if l.IsEmpty() {
		return nil, &lists.FetchingInEmptinessError{}
	}
	return l.elements[l.size-1], nil
}

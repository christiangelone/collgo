package arraylist

import (
	"github.com/christiangelone/collgo/elements"
	"github.com/christiangelone/collgo/lists"
)

type rightIterator struct {
	firstIndex uint64
	currentIndex uint64
	list *ArrayList
}

func (iter *rightIterator) WithOffset(offset uint64) lists.IIterator {
	if offset > 0 {
		if (iter.currentIndex + offset) < iter.list.Size() {
			iter.currentIndex = iter.currentIndex + offset
		}else{
			iter.currentIndex = iter.list.Size() - 1
		}
		iter.firstIndex = iter.currentIndex
	}
	return iter
}

func (iter *rightIterator) First() (element elements.IElement, end bool) {
	element, err := iter.list.GetAt(iter.firstIndex)
	return element, err != nil
}

func (iter *rightIterator) Next() (elements.IElement, bool) {
	if !iter.list.IsEmpty() && iter.currentIndex < iter.list.Size() {
		iter.currentIndex++
		element, err := iter.list.GetAt(iter.currentIndex)
		return element, err != nil
	}
	return nil, true
}

type leftIterator struct {
	firstIndex uint64
	currentIndex uint64
	list *ArrayList
}

func (iter *leftIterator) First() (element elements.IElement, end bool) {
	element, err := iter.list.GetAt(iter.firstIndex)
	return element, err != nil
}

func (iter *leftIterator) WithOffset(offset uint64) lists.IIterator {
	if offset > 0 {
		if (iter.currentIndex - offset) > 0 {
			iter.currentIndex = iter.currentIndex - offset
		}else{
			iter.currentIndex = 0
		}
		iter.firstIndex = iter.currentIndex
	}
	return iter
}

func (iter *leftIterator) Next() (elements.IElement, bool) {
	if !iter.list.IsEmpty() && iter.currentIndex >= 0 {
		iter.currentIndex--
		element, err := iter.list.GetAt(iter.currentIndex)
		return element, err != nil
	}
	return nil, true
}
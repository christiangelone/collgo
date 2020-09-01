package linkedlist

import (
	"github.com/christiangelone/collgo/elements"
	"github.com/christiangelone/collgo/lists"
)

type rightIterator struct {
	firstNode   *node
	currentNode *node
}

func (iter *rightIterator) WithOffset(offset uint64) lists.IIterator {
	_, end := iter.First()
	for i := 0; uint64(i) < offset && !end; i++ {
		_, end = iter.Next()
	}
	iter.firstNode = iter.currentNode
	return iter
}

func (iter *rightIterator) First() (element elements.IElement, end bool) {
	if iter.firstNode != nil {
		return iter.firstNode.element, false
	}
	return nil, true
}

func (iter *rightIterator) Next() (elements.IElement, bool) {
	if iter.currentNode != nil {
		if iter.currentNode = iter.currentNode.next; iter.currentNode != nil {
			return iter.currentNode.element, false
		}
		return nil, true
	}
	return nil, true
}

type leftIterator struct {
	firstNode   *node
	rootNode    *node
	currentNode *node
}

func (iter *leftIterator) First() (element elements.IElement, end bool) {
	if iter.currentNode == nil {
		return nil, true
	}
	return iter.currentNode.element, false
}

func (iter *leftIterator) WithOffset(offset uint64) lists.IIterator {
	_, end := iter.First()
	for i := 0; uint64(i) < offset && !end; i++ {
		_, end = iter.Next()
	}
	iter.firstNode = iter.currentNode
	return iter
}

func (iter *leftIterator) Next() (elements.IElement, bool) {
	if iter.currentNode != nil {
		if iter.currentNode = iter.currentNode.previous; iter.currentNode != nil {
			return iter.currentNode.element, false
		}
		return nil, true
	}
	return nil, true
}

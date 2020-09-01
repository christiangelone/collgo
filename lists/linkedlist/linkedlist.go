package linkedlist

import (
	"github.com/christiangelone/collgo/elements"
	. "github.com/christiangelone/collgo/elements/primitives"
	"github.com/christiangelone/collgo/lists"
)

type node struct {
	element  elements.IElement
	next     *node
	previous *node
}

func New() *LinkedList {
	return &LinkedList{
		size:     0,
		rootNode: nil,
		lastNode: nil,
	}
}

func NewFrom(iterable lists.IIterable) *LinkedList {
	list := New()
	iterable.ForEach(func(element elements.IElement) {
		list.Add(element)
	})
	return list
}

type LinkedList struct {
	size     uint64
	rootNode *node
	lastNode *node
}

func swap(nodeA *node, nodeB *node) {
	nodeA.element, nodeB.element = nodeB.element, nodeA.element
}

func quickSort(startNode *node, endNode *node, compareFn lists.CompareFn){
	if endNode != nil && startNode != endNode && startNode != endNode.next {
		pivotElement := endNode.element
		i := startNode.previous
		for j := startNode; j != endNode; j = j.next {
			if compareFn(j.element, pivotElement) {
				if i == nil { i = startNode } else { i = i.next }
				swap(i, j)
			}
		}
		if i == nil { i = startNode } else { i = i.next }
		swap(i, endNode)
		quickSort(startNode, i.previous, compareFn)
		quickSort(i.next, endNode, compareFn)
	}
}

func (l *LinkedList) Sort(compareFn lists.CompareFn) lists.IList {
	if l.IsSorted(compareFn)  {
		return NewFrom(l)
	}
	listToSort := NewFrom(l)
	quickSort(listToSort.rootNode, listToSort.lastNode, compareFn)
	return listToSort
}

func (l *LinkedList) IsSorted(compareFn lists.CompareFn) bool {
	if l.Size() < 2 {
		return true
	}
	iterator := l.GetRightIterator()
	seenElement, noMore := iterator.First()
	for currentElement, end := iterator.Next(); !noMore && !end; currentElement, end = iterator.Next() {
		if compareFn(seenElement, currentElement) {
			seenElement = currentElement
		}else{
			return false
		}
	}
	return true
}

func (l *LinkedList) Has(element elements.IElement) bool {
	iterator := l.GetRightIterator()
	for currentElement, end := iterator.First(); !end; currentElement, end = iterator.Next() {
		if currentElement.Value() == element.Value() {
			return true
		}
	}
	return false
}

func (l *LinkedList) Zip(collection lists.IList) lists.IList {
	if l.Size() == collection.Size() {
		newList := New()
		selfIterator := l.GetRightIterator()
		collectionIterator := collection.GetRightIterator()
		selfElement, selfEnd := selfIterator.First()
		collectionElement, collectionEnd := collectionIterator.First()
		for {
			if selfEnd || collectionEnd { break }
			newList.Add(NewStruct([]interface{}{
				selfElement.Value(),
				collectionElement.Value(),
			}))
			selfElement, selfEnd = selfIterator.Next()
			collectionElement, collectionEnd = collectionIterator.Next()
		}
		return newList
	}
	panic(&lists.DifferentSizeError{
		Expected: l.Size(),
		Actual: collection.Size(),
	})
}

func (l *LinkedList) AdjustedZip(collection lists.IList, fillValue interface{}) lists.IList {
	newList := New()
	selfIterator := l.GetRightIterator()
	collectionIterator := collection.GetRightIterator()
	selfElement, selfEnd := selfIterator.First()
	collectionElement, collectionEnd := collectionIterator.First()
	for {
		if !selfEnd && !collectionEnd {
			newList.Add(NewStruct([]interface{}{
				selfElement.Value(),
				collectionElement.Value(),
			}))
		} else if selfEnd && !collectionEnd {
			newList.Add(NewStruct([]interface{}{
				fillValue,
				collectionElement.Value(),
			}))
		} else if !selfEnd && collectionEnd {
			newList.Add(NewStruct([]interface{}{
				selfElement.Value(),
				fillValue,
			}))
		} else {
			break
		}
		selfElement, selfEnd = selfIterator.Next()
		collectionElement, collectionEnd = collectionIterator.Next()
	}
	return newList
}

func (l *LinkedList) LeftReduce(reduceFn lists.ReduceFn) elements.IElement {
	iterator := l.GetLeftIterator()
	reducedElement, noMore := iterator.First()
	for element, end := iterator.Next(); !noMore && !end; element, end = iterator.Next() {
		reducedElement = reduceFn(reducedElement, element)
	}
	return reducedElement
}

func (l *LinkedList) LeftReduceWith(reduceFn lists.ReduceFn, initialElement elements.IElement) elements.IElement {
	iterator := l.GetLeftIterator()
	reducedElement := initialElement
	for element, end := iterator.Next(); !end; element, end = iterator.Next() {
		reducedElement = reduceFn(reducedElement, element)
	}
	return reducedElement
}

func (l *LinkedList) RightReduce(reduceFn lists.ReduceFn) elements.IElement {
	iterator := l.GetRightIterator()
	reducedElement, noMore := iterator.First()
	for element, end := iterator.Next(); !noMore && !end; element, end = iterator.Next() {
		reducedElement = reduceFn(reducedElement, element)
	}
	return reducedElement
}

func (l *LinkedList) RightReduceWith(reduceFn lists.ReduceFn, initialElement elements.IElement) elements.IElement {
	reducedElement := initialElement
	l.ForEach(func(element elements.IElement) {
		reducedElement = reduceFn(reducedElement, element)
	})
	return reducedElement
}

func (l *LinkedList) Map(mapFn lists.MapFn) lists.IList {
	newList := New()
	l.ForEach(func(element elements.IElement) {
		newList.Add(mapFn(element))
	})
	return newList
}

func (l *LinkedList) Filter(filterFn lists.FilterFn) lists.IList {
	newList := New()
	l.ForEach(func(element elements.IElement) {
		if filterFn(element) {
			newList.Add(element)
		}
	})
	return newList
}

func (l *LinkedList) Top() (elements.IElement, error) {
	return l.Last()
}

func (l *LinkedList) Push(element elements.IElement) {
	l.Add(element)
}

func (l *LinkedList) Pop() (elements.IElement, error) {
	if l.IsEmpty() {
		panic(&lists.RemovingInEmptinessError{})
	}
	element := l.lastNode.element
	if l.rootNode == l.lastNode {
		l.lastNode = nil
		l.rootNode = nil
	} else {
		l.lastNode = l.lastNode.previous
		l.lastNode.next = nil
	}
	l.size--
	return element, nil
}

func (l *LinkedList) Head() (elements.IElement, error) {
	return l.Last()
}

func (l *LinkedList) Enqueue(element elements.IElement) {
	l.Add(element)
}

func (l *LinkedList) Dequeue() (elements.IElement, error) {
	if l.IsEmpty() {
		panic(&lists.RemovingInEmptinessError{})
	}
	element := l.rootNode.element
	if l.rootNode == l.lastNode {
		l.lastNode = nil
		l.rootNode = nil
	} else {
		l.rootNode = l.rootNode.next
	}
	l.size--
	return element, nil
}

func (l *LinkedList) GetRightIterator() lists.IIterator {
	return &rightIterator{
		firstNode: l.rootNode,
		currentNode: l.rootNode,
	}
}

func (l *LinkedList) GetLeftIterator() lists.IIterator {
	return &leftIterator{
		firstNode: l.lastNode,
		rootNode: l.rootNode,
		currentNode: l.lastNode,
	}
}

func (l *LinkedList) ForEach(eachFn lists.EachFn) {
	iterator := l.GetRightIterator()
	for elem, end := iterator.First(); !end; elem, end = iterator.Next() {
		eachFn(elem)
	}
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) Size() uint64 {
	return l.size
}

func (l *LinkedList) Add(element elements.IElement) {
	if l.IsEmpty() {
		l.rootNode = &node{
			element:  element,
			previous: nil,
			next:     nil,
		}
		l.lastNode = l.rootNode
	} else {
		l.lastNode.next = &node{
			element:  element,
			previous: l.lastNode,
			next:     nil,
		}
		l.lastNode = l.lastNode.next
	}
	l.size++
}

func (l *LinkedList) First() (elements.IElement, error) {
	if l.IsEmpty() {
		return nil, &lists.FetchingInEmptinessError{}
	}
	return l.rootNode.element, nil
}

func (l *LinkedList) Last() (elements.IElement, error) {
	if l.IsEmpty() {
		return nil, &lists.FetchingInEmptinessError{}
	}
	return l.lastNode.element, nil
}

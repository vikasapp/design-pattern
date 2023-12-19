package iterator

type IIterable interface {
	getIterator() IIterator
}

package iterator

type IIterator interface {
	hasNext() bool
	next()
	getCurrentItem() *User
}

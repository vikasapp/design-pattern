package iterator

type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}

	return false
}

func (u *UserIterator) next() {
	if u.hasNext() {
		u.index++
	}
}

func (u *UserIterator) getCurrentItem() *User {
	if u.hasNext() {
		return u.users[u.index]
	}

	return nil
}

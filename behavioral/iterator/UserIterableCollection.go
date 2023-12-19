package iterator

type UserIterableCollection struct {
	users []*User
}

func (u *UserIterableCollection) getIterator() IIterator {
	return &UserIterator{
		users: u.users,
	}
}

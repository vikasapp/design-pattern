package iterator

import "fmt"

func TestIteratorPattern() {
	user := &User{name: "Akash", age: 22}
	user2 := &User{name: "Vikas", age: 25}

	userIterableCollection := &UserIterableCollection{users: []*User{user, user2}}
	iterator := userIterableCollection.getIterator()

	for iterator.hasNext() {
		item := iterator.getCurrentItem()
		fmt.Println("Current Item is ", item)
		iterator.next()
	}
}

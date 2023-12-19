package builder

import "fmt"

func TestBuilderPattern() {
	houseBuilder := GetHouseBuilder()
	director := NewDirector(houseBuilder)
	house := director.BuildHouse()
	fmt.Println(house)
}

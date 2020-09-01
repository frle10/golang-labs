package main

import (
	"fmt"

	"fergo.info.tm/homework3/pets/animals"
)

func main() {
	pets := []animals.Animal{
		&animals.Cat{CatName: "Mittens"}, // for help and learning try to remove `&` here and see what happens
		&animals.Cat{CatName: "Smokey"},
		&Parrot{parrotName: "Tiki"},
		// just for fun, try to add dog same way as Cat here, it won't work
	}
	dog := &animals.Dog{}
	dog.NameSet("Axel")
	pets = append(pets, dog)
	result := count_animals(pets)
	fmt.Println(result) // this example will print map[Cat:2 Dog:1 Parrot:1]
	resultCats := count_cats(pets)
	fmt.Println(resultCats) // this example will print 2
}

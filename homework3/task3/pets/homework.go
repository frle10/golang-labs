package main

import (
	"fergo.info.tm/homework3/pets/animals"
)

type Parrot struct {
	parrotName string
}

func (c *Parrot) Color() string { // leave this also here to see if we have a problem with it
	return "multicolor"
}

//add all required functions
func (p *Parrot) NameSet(name string) {
	p.parrotName = name
}

func (p *Parrot) NameGet() string {
	return p.parrotName
}

func (p *Parrot) Species() string {
	return "Parrot"
}

func (p *Parrot) Sound() string {
	return "squawk"
}

func count_animals(pets []animals.Animal) map[string]int {
	countMap := make(map[string]int)

	for _, animal := range pets {
		countMap[animal.Species()] += 1
	}

	return countMap
}

func count_cats(pets []animals.Animal) int {
	catCount := 0

	for _, animal := range pets {
		_, ok := animal.(*animals.Cat)
		if ok {
			catCount++
		}
	}

	return catCount
}

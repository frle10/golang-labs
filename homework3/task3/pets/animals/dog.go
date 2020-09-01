package animals

type Dog struct {
	dogName string
}

func (c *Dog) NameSet(name string) {
	c.dogName = name
}

func (c *Dog) NameGet() string {
	return c.dogName
}

func (c *Dog) Species() string {
	return "Dog"
}

func (c *Dog) Sound() string {
	return "woof"
}

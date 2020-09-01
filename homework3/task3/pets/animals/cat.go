package animals

type Cat struct {
	CatName string // yes this can also be available publicly, but not through interfaces
}

func (c *Cat) NameSet(name string) {
	c.CatName = name
}

func (c *Cat) NameGet() string {
	return c.CatName
}

func (c *Cat) Species() string {
	return "Cat"
}

func (c *Cat) Sound() string {
	return "mijau"
}

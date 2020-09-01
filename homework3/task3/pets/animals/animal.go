package animals

type Animal interface {
	NameGet() string
	NameSet(name string)
	Species() string
}

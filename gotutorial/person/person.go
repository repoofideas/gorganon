package person

type Person struct {
	name string
	age  int
}

// pointer to person structure from the main func
func (p *Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age
}

func (p Person) Name() string {
	return p.name
}

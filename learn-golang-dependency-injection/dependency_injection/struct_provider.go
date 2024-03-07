package dependency_injection

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
}

func NewBar() *Bar {
	return &Bar{}
}

type NewFooBar struct {
	*Foo
	*Bar
}

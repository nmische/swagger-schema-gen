package foo

type Foo struct {
	FooProperty string `json:"fooProperty"`
	Name        string `json:"name"`
}

func (f Foo) GetGreeting(msg string) string {
	return msg + f.Name
}

type Bar struct {
	BarProperty string `json:"barProperty"`
}

type BarExtended struct {
	Bar
	AnotherProperty    string `json:"anotherProperty"`
	YetAnotherProperty int
}

type FooWrapper struct {
	Foo Foo
}

type BarContainer struct {
	Bars []Bar
}

type FooBar struct {
	Foo
	Bar
}

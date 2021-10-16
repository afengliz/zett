package demo

const Key = "zett:demo"
type IDemo interface {
	GetFoo() Foo
}
type Foo struct {
	Name string
}

package main

type Obj1 struct {
}
type Obj2 struct {
}

func (o Obj1) object() {}
func (o Obj2) object() {}

type Obj3 struct {
	Obj1
	Obj2
}

func main() {
}

package main

import "fmt"

func main() {
	fEntity := FragmentImpl{1}
	fmt.Printf("初始化：%v\n", &fEntity)
	fEntity.Exec()
	fmt.Printf("值调用指针方法后：%v\n", fEntity)
	(&fEntity).Exec()
	fmt.Printf("指针调用指针方法后：%v\n", fEntity)

	fEntity.Exec2()
	fmt.Printf("值调用值方法后：%v\n", fEntity)
	(&fEntity).Exec2()
	fmt.Printf("指针调用值方法后：%v\n", fEntity)


}

func (f FragmentImpl) Exec2() error {
	fmt.Printf("%p\n", &f)
	f.i = f.i + 1
	return nil
}

type FragmentImpl struct {
	i int
}

type FragmentImpl2 struct {
	i int
}

func (f *FragmentImpl2) Exec() error {
	fmt.Printf("%p\n", &f)
	f.i = f.i + 1
	return nil
}

func (f *FragmentImpl) Exec() error {
	fmt.Printf("%p\n", &f)
	f.i = f.i + 1
	return nil
}

type Fragment interface {
	Exec() error
}

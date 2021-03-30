package main

import "fmt"

func main(){
	/**
	* 类型断言
	**/
	p1:=NewPerson("张三")
	var s fmt.Stringer
	s = p1
	// p2 := s.(person)传person就在运行时抛出异常，终止允许
	p2 := s.(*person)//小提示：这里返回的 p2 已经是 *person 类型了，也就是在类型断言的时候，同时完成了类型转换。
	fmt.Println(p2)
	// a:=s.(address)
    // fmt.Println(a)//这个代码在编译的时候不会有问题，因为 address 实现了接口 Stringer，但是在运行的时候，会抛出异常信息：panic: interface conversion: fmt.Stringer is *main.person, not main.address

	// 这显然不符合我们的初衷，我们本来想判断一个接口的值是否是某个具体类型，但不能因为判断失败就导致程序异常。考虑到这点，Go 语言为我们提供了类型断言的多值返回，如下所示：
	// 类型断言返回的第二个值“ok”就是断言是否成功的标志，如果为 true 则成功，否则失败。
	a,ok:=s.(address)
    if ok {
        fmt.Println(a)
    }else {
        fmt.Println("s不是一个address")
    }
}

type person struct {
	name string
	age int
	address
}
type address struct {
	province string
	city string
	
}
func NewPerson(name string) *person {
    return &person{name:name}
}
func (p *person) String() string{
	return fmt.Sprintf("the name is %s,age is %d",p.name,p.age)
}
func (addr address) String() string{
	return fmt.Sprintf("the addr is %s%d",addr.province,addr.city)
}

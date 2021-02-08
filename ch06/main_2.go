package main

import "fmt"

func main(){
	/**
	* 类型断言
	**/
	p1:=NewPerson("张三")
	var s fmt.Stringer
	s = p1
	p2 := s.(*person)
	fmt.Println(p2)
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
//练习实现有两个方法的接口
type WalkRun interface {
	Walk()
	Run()
}

func (p *person) Walk()  {
	fmt.Println("%s能走\n",p.name)
}
func (p *person) Run() {
	fmt.Println("%s能跑\n",p.name)
	
}
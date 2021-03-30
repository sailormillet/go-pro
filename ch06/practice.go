package main
import "fmt"
func main()  {
	
	p:=person{name:"amy"}
	p.Walk()
}

//练习实现有两个方法的接口
type person struct {
	name string
	age int
}
type WalkRun interface {
	Walk()
	Run()
}

func (p *person) Walk()  {
	fmt.Printf("%s能走\n",p.name)
}
func (p *person) Run() {
	fmt.Println("%s能跑\n",p.name)
	
}
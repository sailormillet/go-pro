package main
import "fmt"
// import "errors"
//struct 和 interface：结构体与接口都实现了哪些功能
func main(){
    /**
    * 结构体、字段结构体
    **/
    // p:=person{"Amy",20}
    // p:=person{age:20,name:"Amy"}
    p:=person{age:20}
    fmt.Println(p)
    
    //字段结构体
    p2:=person2{
        age:30,
        name:"Amy",
        addr:address{
            province: "beijing",
            city:     "beijing",
        },
    }
    fmt.Println(p2.addr.province)
    /**
    * 接口的定义 Stringer 是 Go SDK 的一个接口，属于 fmt 包。
    **/
    
    printString(p2)
    printString(p2.addr)

    /**
    *  值接收者和指针接收者 
    * -当值类型作为接收者时，person 类型和*person类型都实现了该接口。
    * - 当指针类型作为接收者时，只有*person类型实现了该接口。
    **/

    printString(&p2)//变量 p2 的指针作为实参传给 printString 函数

    /**
    * 工厂函数
    **/
    pNew:=NewPerson("张三")
    fmt.Println(pNew)
    errNew :=New("出错了");
    fmt.Println(errNew)
}
/**
    * 结构体、字段结构体
    **/
type person struct {
    name string
    age uint
}
type address struct {
    province string
    city string
}
type person2 struct {
    name string
    age uint
    addr address
}
/**
* 接口的定义 Stringer 是 Go SDK 的一个接口，属于 fmt 包。
**/
type Stringer interface {
    String() string
}
// 接口的实现
// 当值类型作为接收者时，person 类型和*person类型都实现了该接口。
func (p person2) String()  string{ //方法接收者  //printString(p2) 和 printString(&p2)都可以
    return fmt.Sprintf("the name is %s,age is %d",p.name,p.age)
}

//当指针类型作为接收者时，只有*person类型实现了该接口。（ printString(&p2)才可以）
// func (p *person2) String()  string{
//     return fmt.Sprintf("the name is %s,age is %d",p.name,p.age)
// }
func (addr address) String()  string{
    return fmt.Sprintf("the addr is %s%s",addr.province,addr.city)
}
func printString(s fmt.Stringer){
    fmt.Println(s.String())
}

/**
* 工厂函数
**/
func NewPerson(name string) *person{
    return &person{name:name}
}
//工厂函数，返回一个error接口，其实具体实现是*errorString
func New(text string) error {
    return &errorString{text}
}
//结构体，内部一个字段s，存储错误信息
type errorString struct {
    s string
}
//用于实现error接口
func (e *errorString) Error() string {
    return e.s
}
/**
* 继承和组合
**/
type Reader interface {
    read(p []byte) (n int, err error)
}
type Writer interface{
    Write(p []byte) (n int, err error)
}
//ReadWriter是Reader和Writer的组合
type ReadWriter interface {
    Reader
    Writer
}
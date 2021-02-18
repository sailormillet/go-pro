package main
import "fmt"
// import "time"
// import "sync"
// import "errors"

/**
    参数传递：值、引用及指针之间的区别？
**/
func main(){
    /**
    修改参数
    Go 语言中的函数传参都是值传递。 值传递指的是传递原来数据的一份拷贝，而不是原来的数据本身。
    **/
    p:=person{name:"amy",age:18}
    fmt.Printf("main函数：p的内存地址为%p\n",&p)
    modifyPerson(&p)
    fmt.Println("person name:", p.name,",age:",p.age)
    /**
        值类型：除了 struct 外，还有浮点型、整型、字符串、布尔、数组，这些都是值类型。
    **/
    /**
        指针类型：指针类型的变量保存的值就是数据对应的内存地址，所以在函数参数传递是传值的原则下，拷贝的值也是内存地址。现在对以上示例稍做修改，修改后的代码如下：
    **/
    /**
        引用类型： map 和 chan。
        严格来说，Go 语言没有引用类型，但是我们可以把 map、chan 称为引用类型，这样便于理解。除了 map、chan 之外，Go 语言中的函数、接口、slice 切片都可以称为引用类型。
        指针类型也可以理解为是一种引用类型。
    **/
    //map
    m:=make(map[string]int) //小提示：用字面量或者 make 函数的方式创建 map，并转换成 makemap 函数的调用，这个转换是 Go 语言编译器自动帮我们做的。makemap 函数返回的是一个 *hmap 类型，也就是说返回的是一个指针，所以我们创建的 map 其实就是一个 *hmap。
    m["amy"] = 18
    fmt.Printf("main函数：m的内存地址为%p\n",m)
    fmt.Println("amy's age",m["amy"])
    modifyMap(m)
    fmt.Println("amy's age",m["amy"])
    /**
        类型的零值
        int float       0
        bool            false
        string          ""(空字符串)
        struct          内部字段的零值
        slice           nil
        map             nil
        指针            nil
        函数            nil
        chan            nil
        interface       nil

    **/
}

type person struct {
    name string
    age int
}
func modifyPerson(p *person)  {
    fmt.Printf("modifyPerson函数：p的内存地址为%p\n",p)
    p.name ="jay"
    p.age=10
}
func modifyMap(p map[string]int)  {
    fmt.Printf("modifyMap函数：p的内存地址为%p\n",p)
    p["amy"]=10
}
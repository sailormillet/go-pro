package main
import "fmt"
// import "time"
// import "sync"
// import "errors"

/**
    指针详解：在什么情况下应该使用指针？
**/
func main(){
    /**
        什么是指针
        我们都知道程序运行时的数据是存放在内存中的，而内存会被抽象为一系列具有连续编号的存储空间，那么每一个存储在内存中的数据都会有一个编号，这个编号就是内存地址。有了这个内存地址就可以找到这个内存中存储的数据，而内存地址可以被赋值给一个指针。
        小提示：内存地址通常为 16 进制的数字表示，比如 0x45b876。
        小技巧：你也可以简单地把指针理解为内存地址。
    **/
    //指针的声明和定义：在 Go 语言中，获取一个变量的指针非常容易，使用取地址符 & 就可以，比如下面的例子：
    name:="amy"
    nameP:=&name
    fmt.Println("name变量的值为：",name)
    fmt.Println("name变量的内存地址为：",nameP)
    /**
    小提示：通过 var 声明的指针变量是不能直接赋值和取值的，因为这时候它仅仅是个变量，还没有对应的内存地址，它的值是 nil。
    **/
    // var intP *int
    // intP = &name //指针类型不同，无法赋值
    // fmt.Println("name变量的内存地址为：",intP)
    intP1:=new(int)//返回一个 *int 类型的 intP1
    fmt.Println("int的内存地址为：",intP1)
    /**
    指针的操作
    在 Go 语言中指针的操作无非是两种：一种是获取指针指向的值，一种是修改指针指向的值。
    **/
    nameV:=*nameP
    fmt.Println("nameP指针指向的值为:",nameV)
    //修改指针指向的值
    *nameP = "amy2" //修改指针指向的值
    fmt.Println("nameP指针指向的值为:",*nameP)
    fmt.Println("name变量的值为:",name)
    /**
    指针参数
    **/
    age:=18
    // modifyAgeWrong(age)
    modifyAge(&age)
    fmt.Println("age的值为：",age)
    /**
    指针接收者
    对于是否使用指针类型作为接收者，有以下几点参考：

    如果接收者类型是 map、slice、channel 这类引用类型，不使用指针；

    如果需要修改方法接收者内部的数据或者状态时，需要使用指针；

    如果需要修改参数的值或者内部数据时，也需要使用指针类型的参数；

    如果是比较大的结构体，每次参数传递或者调用方法都要内存拷贝，内存占用多，这时候可以考虑使用指针；

    像 int、bool 这样的小数据类型没必要使用指针；

    如果需要并发安全，则尽可能地不要使用指针，使用指针一定要保证并发安全；

    指针最好不要嵌套，也就是不要使用一个指向指针的指针，虽然 Go 语言允许这么做，但是这会使你的代码变得异常复杂。

    指针的两大好处：

    1、可以修改指向数据的值；

    2、在变量赋值，参数传值的时候可以节省内存。
    **/

    add := address{province:"beijing",city:"beijing"}
    printString(add)
    printString(&add)
    //定义一个指向接口的指针,虽然指向具体类型的指针可以实现一个接口，但是指向接口的指针永远不可能实现该接口。
    // var sh fmt.Stringer = address{province:"shanghai",city:"shanghai"}
    // printString(sh)
    // shp := &sh
    // printString(shp)
}
// func modifyAgeWrong(age int)  {  
//     age = 20 // modifyAgeWrong 中的 age 只是实参 age 的一份拷贝，所以修改它不会改变实参 age 的值。
// }
func modifyAge(age *int)  {
    *age = 20
}
type address struct {
    province string
    city string
}

func (addr address) String() string {
    return fmt.Sprintf("the addr is %s%s",addr.province,addr.city)
}
func printString(s fmt.Stringer)  {
    fmt.Println(s.String())
}
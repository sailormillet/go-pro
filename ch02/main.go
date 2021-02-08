package main
import "fmt"
import "strconv"
import "strings"
func main(){
    //变量类型
    var f32 float32 = 2.2
    var f64 float64 = 10.3456
    fmt.Println(f32,f64)
    var bf bool = false
    var bt bool = true
    fmt.Println("bf is ",bf ,"bt is " ,bt)
    var s1 string = "hello"
    var s2 string = "world"
    fmt.Println("s1 is ",s1 ,"s2 is " ,s2)
    fmt.Println("s1 + s2 ",s1+s2)
    //未赋值的初始值
    var zs string
    var zb bool
    var zi int
    var zf float32
    fmt.Println(zs,zb,zi,zf)
    //变量简短声明
    i :=10
    bf1 :=false
    s11 :="hello"
    fmt.Println(i,bf1,s11)
    //指针
    pi:=&i
    fmt.Println(*pi)
    //赋值
    i=20
    fmt.Println(i)
    //常量
    const name = "amy"
    fmt.Println(name)
    //iota 是一个常量生成器它可以用来初始化相似规则的常量，避免重复的初始化。假设我们要定义 one、two、three 和 four 四个常量，对应的值分别是 1、2、3 和 4，如果不使用 iota，则需要按照如下代码的方式定义：
    // const(
    //     one = 1
    //     two = 2
    //     three =3
    //     four =4
    // )
    const(
        one = iota+1
        two
        three
        four
    )
    fmt.Println(one,two,three,four)
    //字符串和数字互转。Go 语言是强类型的语言，需要先进行类型转换来相互使用和计算的。
    i2s := strconv.Itoa(i)
    s2i, err := strconv.Atoi(i2s)
    i2f:= float64(i)
    f2i:= int(f64)
    fmt.Println(i2s,s2i,err,i2f,f2i)
    //Strings 包https://golang.google.cn/pkg/strings/
    fmt.Println(strings.HasPrefix(s11,"H"))
    fmt.Println(strings.Index(s11,"o"))
    fmt.Println(strings.ToUpper(s11))
}
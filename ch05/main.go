package main
import "fmt"
import "errors"
//Go 语言中的函数和方法
func main(){
    /**
    * 函数
    **/
    result := sum(1,2)
    fmt.Println(result)

    //多值返回
    result1,err1 := sum1(-1, 2)
    // result1,_ := sum1(1, 2)//函数的返回值不需要，可以使用下划线 _ 丢弃它
    if err1!=nil {
        fmt.Println(err1)
    }else {
        fmt.Println(result)
    }
    fmt.Println(result1,err1)

    //命名返回参数
    result2,err2 := sum2(-1, 2)
    fmt.Println(result2,err2)

    //可变参数
    fmt.Println(sum3(1,2,3))
    //包级函数
    // 函数名称首字母小写代表私有函数，只有在同一个包中才可以被调用；

    // 函数名称首字母大写代表公有函数，不同的包也可以调用；

    // 任何一个函数都会从属于一个包。

    //匿名函数和闭包
    //匿名函数
    sum4 := func(a, b int) int {
        return a + b
    }
    fmt.Println(sum4(1, 2))
    //闭包
    cl:=colsure()
    fmt.Println(cl())
    fmt.Println(cl())
    fmt.Println(cl())


    /**
    * 不同于函数的方法 在 Go 语言中，方法和函数是两个概念，但又非常相似，不同点在于方法必须要有一个接收者，这个接收者是一个类型，这样方法就和这个类型绑定在一起，称为这个类型的方法。
    **/
    age:=Age(25)
    age.String()
    // age.Modify()
    (&age).Modify()
    age.String()
    //方法赋值给变量，方法表达式
    sm:=Age.String
    //通过变量，要传一个接收者进行调用也就是age
    sm(age)
    
}
// func sum(a int,b int) int{
//     return a+b
// }
//函数声明
func sum(a, b int) int {
    return a + b
}
// 多值返回 第一个值返回函数的结果，第二个值返回函数出错的信息，这种就是多值返回的经典应用。
func sum1(a ,b int) (int,error){
    if a<0 || b<0 {
        return 0, errors.New("a或者b不能是负数")
    }
    return a + b,nil
}

// 命名返回参数
func sum2(a, b int) (sum int,err error){
    if a<0 || b<0 {
        return 0,errors.New("a或者b不能是负数")
    }
    sum=a+b
    err=nil
    return 
}
//可变参数
func sum3(params ...int) int {
    sum := 0
    for _, i := range params {
        sum += i
    }
    return sum
}
//闭包
func colsure() func() int {
    i:=0
    return func() int {
        i++
        return i
    }
}

//方法
//值类型接收者 => 示例中方法 String() 就是类型 Age 的方法，类型 Age 是方法 String() 的接收者。
type Age uint //type Age uint 表示定义一个新类型 Age，该类型等价于 uint，可以理解为类型 uint 的重命名。
func (age Age) String() {
    fmt.Println("the age is",age)
}
//指针类型接收者
func (age *Age) Modify(){
    *age = Age(30)
}
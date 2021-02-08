package main
import "fmt"
import "strconv"
import "errors"
import (
    "io/ioutil"
    "os"
)
/**
* 错误处理：如何通过 error、deferred、panic 等处理错误？
**/
func main(){
  i,err := strconv.Atoi("a")
  if err !=nil {
      fmt.Println(err)
  }else{
      fmt.Println(i)
  }

  sum,err:=add(-1,2)
  if err != nil  {
    fmt.Println(err)
  }else{
    fmt.Println(sum)
  }
//   error 断言
  if cm,ok := err.(*commonError); ok{
    fmt.Println("错误代码为:",cm.errorCode,"，错误信息为：",cm.errorMsg)
  }else{
    fmt.Println(sum)
  }

  //错误嵌套 
  e := errors.New("原始错误e")
  newErr := MyError{e,"数据上传问题"}
  fmt.Println(e,newErr)

  //go 1.13版本增加 Error Wrapping 功能 替换 错误嵌套
  w := fmt.Errorf("Wrap了一个错误:%w", e)
  fmt.Println(w)
  fmt.Println(errors.Unwrap(w))
  //   errors.Is 函数:判断两个 error 是否是同一个
  fmt.Println(errors.Is(w,e))

  //errors.As 函数 有了 error 嵌套后，error 断言也不能用了，因为你不知道一个 error 是否被嵌套，又嵌套了几层。所以 Go 语言为解决这个问题提供了 errors.As 函数，比如前面 error 断言的例子，可以使用 errors.As 函数重写，效果是一样的
  var cm *commonError
  if errors.As(err,&cm) {
      fmt.Println("错误代码为:",cm.errorCode,"，错误信息为：",cm.errorMsg)
  }else{
      fmt.Println(sum)
  }
  //Deferred 函数 Go 语言为我们提供了 defer 函数，可以保证文件关闭后一定会被执行，不管你自定义的函数出现异常还是错误。
  // Go 语言标准包 ioutil 中的 ReadFile 函数，它需要打开一个文件，然后通过 defer 关键字确保在 ReadFile 函数执行结束后，f.Close() 方法被执行，这样文件的资源才一定会释放。
  fileContent, err := read("D:/test.txt")
  fmt.Println(fileContent,err)
  if err == nil{
    fmt.Println("File Content =", string(fileContent))
    }else{
        fmt.Println("Read file err, err =", err)
    }
    // 多个defer 在一个方法或者函数中，可以有多个 defer 语句；多个 defer 语句的执行顺序依照后进先出的原则。
    moreDefer()
    // Panic 异常 panic 异常是一种非常严重的情况，会让程序中断运行，使程序崩溃，所以如果是不影响程序运行的错误，不要使用 panic，使用普通错误 error 即可。
    // Recover 捕获 Panic 异常
    defer func() {
        if p:=recover();p!=nil{
           fmt.Println(p)
        }
     }()
     connectMySQL("","root","123456")
     
}

type error interface {
    Error() string
}
// error 工厂函数
func add(a,b int) (int,error){
    if a<0 || b<0 {
        // return 0,errors.New("a或者b不能为负数")
        return 0, &commonError{
                    errorCode: 1,
                    errorMsg:  "a或者b不能为负数"}
    }else{
        return a+b,nil
    }
}
//自定义 error
type commonError struct {
    errorCode int//错误码
    errorMsg string //错误信息
}
func (ce * commonError) Error() string{
    return ce.errorMsg
}

//错误嵌套 
type MyError struct {
    err error
    msg string
}
func (e *MyError) Error() string{
    return e.err.Error() + e.msg
}
//Deferred 函数
func read(filename string) ([]byte,error){
    f,err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer f.Close()
    // buf := make([]byte, 1024)
    return  ioutil.ReadAll(f)
}
// 多个defer
func moreDefer(){
    defer  fmt.Println("First defer")
    defer  fmt.Println("Second defer")
    defer  fmt.Println("Three defer")
    fmt.Println("函数自身代码")
 }
// Panic 异常
func connectMySQL(ip,username,password string){
    if ip =="" {
       panic("ip不能为空")
    }
    //省略其他代码
 }

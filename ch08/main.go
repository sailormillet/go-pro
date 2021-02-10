package main
import "fmt"
import "time"
// import "strconv"
// import "errors"
// import (
//     "io/ioutil"
//     "os"
// )

/**
  并发基础：Goroutines 和 Channels 的声明与使用 从使用上讲，Go 语言还是更推荐 channel 这种更高级别的并发控制方式，因为它更简洁，也更容易理解和使用。

  什么是并发 ->并发可以让你编写的程序在同一时刻做多几件事情。
  进程 -> 在操作系统中，进程是一个非常重要的概念。当你启动一个软件（比如浏览器）的时候，操作系统会为这个软件创建一个进程，这个进程是该软件的工作空间，它包含了软件运行所需的所有资源，比如内存空间、文件句柄，还有下面要讲的线程等。
  线程 ->线程是进程的执行空间，一个进程可以有多个线程，线程被操作系统调度执行，比如下载一个文件，发送一个消息等。这种多个线程被操作系统同时调度执行的情况，就是多线程的并发。
一个程序启动，就会有对应的进程被创建，同时进程也会启动一个线程，这个线程叫作主线程。如果主线程结束，那么整个程序就退出了。有了主线程，就可以从主线里启动很多其他线程，也就有了多线程的并发。
  
channel 为什么是并发安全的呢？是因为 channel 内部使用了互斥锁来保证并发的安全。

在 Go 语言中，提倡通过通信来共享内存，而不是通过共享内存来通信，其实就是提倡通过 channel 发送接收消息的方式进行数据传递，而不是通过修改同一个变量。所以在数据流动、传递的场景中要优先使用 channel，它是并发安全的，性能也不错。
**/
func main(){
    /**
    协程（Goroutine）Go 语言中没有线程的概念，只有协程，也称为 goroutine。相比线程来说，协程更加轻量，一个程序可以随意启动成千上万个 goroutine。
    **/
    go fmt.Println("hello goroutine") //go 关键字后跟一个方法或者函数的调用，就可以启动一个 goroutine ，程序是并发的，go 关键字启动的 goroutine 并不阻塞 main goroutine 的执行，所以我们才会看到如下打印结果。
    fmt.Println("main goroutine") 
    // time.Sleep(time.Second)//表示等待一秒，这里是让 main goroutine 等一秒，不然 main goroutine 执行完毕程序就退出了，也就看不到启动的新 goroutine 的hello goroutine打印结果了

    /**
    Channel 如果启动了多个 goroutine，它们之间该如何通信呢？这就是 Go 语言提供的 channel（通道）要解决的问题。
    在 Go 语言中，声明一个 channel 非常简单，使用内置的 make 函数即可。
    chan 是一个关键字，表示是 channel 类型。 后面的 string 表示 channel 里的数据是 string 类型。通过 channel 的声明也可以看到，chan 是一个集合类型。
    定义好 chan 后就可以使用了，一个 chan 的操作只有两种：发送和接收。
    接收：获取 chan 中的值，操作符为 <- chan。
    发送：向 chan 发送值，把值放在 chan 中，操作符为 chan <-。
    **/
    /**
    无缓冲 channel,也可以称为同步 channel.使用 make 创建的 chan 就是一个无缓冲 channel，它的容量是 0，不能存储任何数据。所以无缓冲 channel 只起到传输数据的作用，数据并不会在 channel 中做任何停留。这也意味着，无缓冲 channel 的发送和接收操作是同时进行的
    无缓冲 channel 其实就是一个容量大小为 0 的 channel。比如 make(chan int,0)。
    **/
    ch := make(chan string)
    go func(){
        fmt.Println("hello first goroutine")
        ch <- "goroutine 完成"
    }()
    v:=<-ch
    fmt.Println("接收都的chan种的值为：",v)
    /**
    有缓冲 channel 有缓冲 channel 类似一个可阻塞的队列，内部的元素先进先出。通过 make 函数的第二个参数可以指定 channel 容量的大小，进而创建一个有缓冲 channel.
    一个有缓冲 channel 具备以下特点：

    有缓冲 channel 的内部有一个缓冲队列；
    
    发送操作是向队列的尾部插入元素，如果队列已满，则阻塞等待，直到另一个 goroutine 执行，接收操作释放队列的空间；
    
    接收操作是从队列的头部获取元素并把它从队列中删除，如果队列为空，则阻塞等待，直到另一个 goroutine 执行，发送操作插入新的元素。
    **/
    cacheCh:=make(chan int,5)
    cacheCh <- 2
    cacheCh <- 3
    fmt.Println("cacheCh容量为:",cap(cacheCh),",元素个数为：",len(cacheCh))

    /**
    关闭 channel
    如果一个 channel 被关闭了，就不能向里面发送数据了，如果发送的话，会引起 painc 异常。
    但是还可以接收 channel 里的数据，如果 channel 里没有数据的话，接收的数据是元素类型的零值。
    **/
    // close(cacheCh)
    /**
    单向 channel
    限制一个 channel 只可以接收但是不能发送，或者限制一个 channel 只能发送但不能接收，这种 channel 称为单向 channel。
    **/
    onlySend := make(chan<- int)
    onlyReceive:=make(<-chan int)
    // onlySend <- 2
    // or:= <- onlyReceive 
    fmt.Println("onlySend:",len(onlySend),",onlyReceive:",len(onlyReceive))
/**
    select+channel 示例
    假设要从网上下载一个文件，我启动了 3 个 goroutine 进行下载，并把结果发送到 3 个 channel 中。其中，哪个先下载好，就会使用哪个 channel 的结果。
    在这种情况下，如果我们尝试获取第一个 channel 的结果，程序就会被阻塞，无法获取剩下两个 channel 的结果，也无法判断哪个先下载好。这个时候就需要用到多路复用操作了，在 Go 语言中，通过 select 语句可以实现多路复用，
**/
//多路复用可以简单地理解为，N 个 channel 中，任意一个 channel 有数据产生，select 都可以监听到，然后执行相应的分支，接收数据并处理。
firstCh := make(chan string)
secondCh := make(chan string)
threeCh := make(chan string)

go func(){
    firstCh <- downloadFile("firstCh")
}()
go func(){
    secondCh <- downloadFile("secondCh")
}()
go func(){
    threeCh <- downloadFile("threeCh")
}()
select {
case filePath := <-firstCh:
    fmt.Println(filePath)
case filePath := <-secondCh:
    fmt.Println(filePath)
case filePath := <-threeCh:
    fmt.Println(filePath)
}
    
}
func counter(out chan<- int) {
//函数内容使用变量out，只能进行发送操作

}
func downloadFile(chanName string) string {
    //模拟下载文件,可以自己随机time.Sleep点时间试试
    time.Sleep(time.Second)
    return chanName+":filePath"
 }
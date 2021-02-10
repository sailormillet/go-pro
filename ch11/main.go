package main
import "fmt"
import "time"
import "sync"
// import "errors"

/**
    Go 语言中即学即用的高效并发模式
**/
func main(){
    /**
        for select 循环模式
    **/

    //第一种模式是  for 循环 +select 多路复用的并发模式，哪个 case 满足就执行哪个，直到满足一定的条件退出 for 循环（比如发送退出信号）。
    // for { //for无限循环，或者for range循环
    //     select {
    //       //通过一个channel控制
    //     }
    //   }

    //第二种模式是 for range select 有限循环

    // for _,s:=range []int{}{
    //     select {
    //     case <-done:
    //        return
    //     case resultCh <- s:
    //     }
    //  }

    /**
        select timeout 模式
        假如需要访问服务器获取数据，因为网络的不同响应时间不一样，为保证程序的质量，不可能一直等待网络返回，所以需要设置一个超时时间，这时候就可以使用 select timeout 模式
        小提示：如果可以使用 Context 的 WithTimeout 函数超时取消，要优先使用。
    **/
    result := make(chan string)
    go func(){
        //模拟网络访问
        time.Sleep(8 * time.Second)
        result <- "服务端结果"
    }()
    select {
        case v:= <-result:
            fmt.Println(v)
        case <-time.After(5*time.Second):
            fmt.Println("网络访问超时了")

    }
    /**
        Pipeline 模式
        Pipeline 模式也称为流水线模式，模拟的就是现实世界中的流水线生产。以手机组装为例，整条生产流水线可能有成百上千道工序，每道工序只负责自己的事情，最终经过一道道工序组装，就完成了一部手机的生产。

        从技术上看，每一道工序的输出，就是下一道工序的输入，在工序之间传递的东西就是数据，这种模式称为流水线模式，而传递的数据称为数据流。
    **/
    coms := buy(100)//采购100套配件
    //三班人同时组装100部手机
    phones1 := build(coms)
    phones2 := build(coms)
    phones3 := build(coms)
    // phones := build(coms)
    //汇聚三个channel成一个
    /**
        扇出和扇入模式merge
    **/
    phones := merge(phones1,phones2,phones3)
    packs := pack(phones)
    for p := range packs{
        fmt.Println(p)
    }

    /**
        Futures 模式
        Futures 模式下的协程和普通协程最大的区别是可以返回结果，而这个结果会在未来的某个时间点使用。所以在未来获取这个结果的操作必须是一个阻塞的操作，要一直等到获取结果为止。

        如果你的大任务可以拆解为一个个独立并发执行的小任务，并且可以通过这些小任务的结果得出最终大任务的结果，就可以使用 Futures 模式。
    **/
    vegetablesCh := washVegetables() //洗菜
    waterCh := boilWater() //烧水
    fmt.Println("已经安排洗菜和烧水了，我先眯一会")
    time.Sleep(2 * time.Second)
    fmt.Println("要做火锅了，看看菜和水好了吗")
    vegetables := <-vegetablesCh
    water := <-waterCh
    fmt.Println("准备好了，可以做火锅了:",vegetables,water)
}

//工序1采购
func buy(n int) <-chan string{
    out := make(chan string)
    go func ()  {
        defer close(out)
        for i:=1; i<=n; i++{
            out <- fmt.Sprint("配件", i)
        }
    }()
    return out
}
//工序2组装
func build (in<-chan string)<-chan string{
    out := make(chan string)
    go func ()  {
       defer close(out)
       for c := range in {
           out <- "组装("+c+")"
       } 
    }()
    return out
}
//工序3打包
func pack(in <-chan string) <-chan string{
    out := make(chan string)
    go func ()  {
        defer close(out)
        for c:=range in{
            out <- "打包("+c+")"
        }
    }()
    return out 
}
/**
 merge 函数的核心逻辑就是对输入的每个 channel 使用单独的协程处理，并将每个协程处理的结果都发送到变量 out 中，达到扇入的目的。总结起来就是通过多个协程并发，把多个 channel 合成一个。
在整条手机组装流水线中，merge 函数非常小，而且和业务无关，不能当作一道工序，所以我把它叫作组件。该 merge 组件是可以复用的，流水线中的任何工序需要扇入的时候，都可以使用 merge 组件。
小提示：这次的改造新增了 merge 函数，其他函数保持不变，符合开闭原则。开闭原则规定“软件中的对象（类，模块，函数等等）应该对于扩展是开放的，但是对于修改是封闭的”。
 **/
func merge(ins ...<-chan string) <-chan string {
    var wg sync.WaitGroup
    out:= make(chan string)
    p:=func (in <-chan string){
        defer wg.Done()
        for c := range in {
            out <- c
        }
    }
    wg.Add(len(ins))
    //扇入，需要启动多个goroutine
    for _,cs :=range ins{
        go p(cs)
    }
    //等待所有输入的数据ins处理完，再关闭输出out
    go func ()  {
        wg.Wait()
        close(out)
    }()
    return out
}

//洗菜
func washVegetables() <-chan string  {
    vegetables := make(chan string)
    go func ()  {
        time.Sleep(5*time.Second)
        vegetables <- "洗好的菜"
    }()
    return vegetables
}
//烧水
func boilWater() <-chan string  {
    water := make(chan string)
    go func(){
        time.Sleep(5*time.Second)
        water <- "烧开的水"
    }()
    return water
}
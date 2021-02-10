package main
import "fmt"
import "time"
import "sync"
// import "strconv"
// import "errors"
// import (
//     "io/ioutil"
//     "os"
// )

/**
  同步原语：sync 包让你对并发控制得心应手

  在 Go 语言中，不仅有 channel 这类比较易用且高级的同步机制，还有 sync.Mutex、sync.WaitGroup 等比较原始的同步机制。通过它们，我们可以更加灵活地控制数据的同步和多协程的并发。
 **/
 //共享的资源
 var (
    sum int
    // mutex sync.Mutex
    mutex sync.RWMutex
 )
func main(){
    /**
    sync.WaitGroup 可以很好地跟踪协程。在协程执行完毕后，整个 run 函数才能执行完毕，时间不多不少，正好是协程执行的时间。
    sync.WaitGroup 用于最终完成的场景，关键点在于一定要等待所有协程都执行完毕。
    time.Sleep(2 * time.Second) 代码，这是为了防止主函数 main 返回使用，一旦 main 函数返回了，程序也就退出了。
    **/
    run()
    /**
    sync.Once 让代码只执行一次，哪怕是在高并发的情况下，比如创建一个单例。
    time.Sleep(2 * time.Second) 代码，这是为了防止主函数 main 返回使用，一旦 main 函数返回了，程序也就退出了。
    **/
    doOnce()
    /**
    sync.Cond  可以用于发号施令，一声令下所有协程都可以开始执行，关键点在于协程开始的时候是等待的，要等待 sync.Cond 唤醒才能执行。
    sync.Cond 从字面意思看是条件变量，它具有阻塞协程和唤醒协程的功能，所以可以在满足一定条件的情况下唤醒协程，但条件变量只是它的一种使用场景。
    我以 10 个人赛跑为例来演示 sync.Cond 的用法。在这个示例中有一个裁判，裁判要先等这 10 个人准备就绪，然后一声发令枪响，这 10 个人就可以开始跑了。
    Wait，阻塞当前协程，直到被其他协程调用 Broadcast 或者 Signal 方法唤醒，使用的时候需要加锁，使用 sync.Cond 中的锁即可，也就是 L 字段。

    Signal，唤醒一个等待时间最长的协程。

    Broadcast，唤醒所有等待的协程。

    注意：在调用 Signal 或者 Broadcast 之前，要确保目标协程处于 Wait 阻塞状态，不然会出现死锁问题。
    **/
    race()

    /**
    sync.Map的使用和内置的 map 类型一样，只不过它是并发安全的 
    方法:
    Store：存储一对 key-value 值。

    Load：根据 key 获取对应的 value 值，并且可以判断 key 是否存在。

    LoadOrStore：如果 key 对应的 value 存在，则返回该 value；如果不存在，存储相应的 value。

    Delete：删除一个 key-value 键值对。

    Range：循环迭代 sync.Map，效果与 for range 一样。
    **/
}
func run()  {
    var wg sync.WaitGroup
    wg.Add(110)
    for i:=0; i<100;i++{
        go func() {
            //计数器值减1
            defer wg.Done()
            add(10)
         }()
    }
    for i:=0; i<10;i++{
        go func() {
            //计数器值减1
            defer wg.Done()
            fmt.Println("读取和为",readSum())
         }()
        
    }
    // time.Sleep(2 * time.Second)
    //一直等待，只要计数器值为0
    wg.Wait()
}
func add(i int)  {
    mutex.Lock()
    defer mutex.Unlock()
    sum += i;
}

func readSum()  int{
    mutex.RLock()
    defer mutex.RUnlock()
    b:=sum
    return b
    
}

func doOnce()  {
    var once sync.Once
    onceBody := func(){
        fmt.Println("Only once")
    }
    //用于等待携程执行完毕
    done:= make(chan bool)
    for i:= 0; i<10; i++ {
        go func ()  {
            once.Do(onceBody)
            done <- true
        }()
    }
    for i:=0; i<10; i++{
        <- done
    }
}
//10个人赛跑，1个裁判发号施令
func race(){
    cond := sync.NewCond(&sync.Mutex{})
    var wg sync.WaitGroup
    wg.Add(11)
    for i:=0; i<10;i++{
        go func(num int) {
            defer  wg.Done()
            fmt.Println(num,"号已经就位")
            cond.L.Lock()
            cond.Wait()//等待发令枪响
            fmt.Println(num,"号开始跑……")
            cond.L.Unlock()
         }(i)
    }
    //等待所有的goroutine 都进入wait状态
    time.Sleep(2*time.Second)
    go func() {
        defer  wg.Done()
        fmt.Println("裁判已经就位，准备发令枪")
        fmt.Println("比赛开始，大家准备跑")
        cond.Broadcast()//发令枪响
     }()
    //防止函数提前返回退出
    wg.Wait()
}
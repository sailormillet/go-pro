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
  同步原语：sync 包让你对并发控制得心应手 同步原语通常用于更复杂的并发控制，如果追求更灵活的控制方式和性能，你可以使用它们。

  在 Go 语言中，不仅有 channel 这类比较易用且高级的同步机制，还有 sync.Mutex、sync.WaitGroup 等比较原始的同步机制。通过它们，我们可以更加灵活地控制数据的同步和多协程的并发。
 **/
 //共享的资源
 var sum =0;
 var (
    sumMutex int
    // mutex sync.Mutex
    mutex sync.RWMutex
 ) 
func main(){
    /**
    资源竞争

    在一个 goroutine 中，如果分配的内存没有被其他 goroutine 访问，只在该 goroutine 中被使用，那么不存在资源竞争的问题。

    但如果同一块内存被多个 goroutine 同时访问，就会产生不知道谁先访问也无法预料最后结果的情况。这就是资源竞争，这块内存可以称为共享的资源。

    小技巧：使用 go build、go run、go test 这些 Go 语言工具链提供的命令时，添加 -race 标识可以帮你检查 Go 语言代码是否存在资源竞争。
    **/

    for i:=0; i<100; i++ {
        //开启100个协程让sum+10
        go add(10)
    }
    for i:=0; i<100; i++ {
        //开启100个协程让sum+10
        go addMutex(10)
    }
    for i:=0;i<10;i++{
        go fmt.Println("读取和为:",readSum())
    }
   //防止提前退出
   time.Sleep(2*time.Second)
   fmt.Println("和为", sum)
   fmt.Println("mutex和为", sumMutex)
     /**
    同步原语 

    sync.Mutex 互斥锁，顾名思义，指的是在同一时刻只有一个协程执行某段代码，其他协程都要等待该协程执行完毕后才能继续执行。

    互斥锁的使用非常简单，它只有两个方法 Lock 和 Unlock，代表加锁和解锁。当一个协程获得 Mutex 锁后，其他协程只能等到 Mutex 锁释放后才能再次获得锁。

    Mutex 的 Lock 和 Unlock 方法总是成对出现，而且要确保 Lock 获得锁后，一定执行 UnLock 释放锁，所以在函数或者方法中会采用 defer 语句释放锁。

    在下面的示例中，我声明了一个互斥锁 mutex，然后修改 add 函数，对 sum+=i 这段代码加锁保护。这样这段访问共享资源的代码片段就并发安全了，可以得到正确的结果。
    **/

    /**
    sync.RWMutex 读写锁 

    对共享资源 sumMutex 的加法操作进行了加锁，这样可以保证在修改 sumMutex 值的时候是并发安全的。如果读取操作也采用多个协程呢？
    每次读写共享资源都要加锁，所以性能低下，这该怎么解决呢？

    读写这个特殊场景，有以下几种情况：
    1、写的时候不能同时读，因为这个时候读取的话可能读到脏数据（不正确的数据）；

    2、读的时候不能同时写，因为也可能产生不可预料的结果；

    3、读的时候可以同时读，因为数据不会改变，所以不管多少个 goroutine 读都是并发安全的。
    把锁的声明换成读写锁 sync.RWMutex。 把函数 readSum 读取数据的代码换成读锁，也就是 RLock 和 RUnlock。性能就会有很大的提升，因为多个 goroutine 可以同时读数据，不再相互等待。
    **/

}
func add(i int)  {
    sum += i
}
func addMutex(i int)  {
    mutex.Lock()
    defer mutex.Unlock() //这样可以确保锁一定会被释放，不会被遗忘
    sumMutex+=i
    // mutex.Unlock()
    /**
        以上被加锁保护的 sumMutex+=i 代码片段又称为临界区。
        在同步的程序设计中，临界区段指的是一个访问共享资源的程序片段，而这些共享资源又有无法同时被多个协程访问的特性。 
        当有协程进入临界区段时，其他协程必须等待，这样就保证了临界区的并发安全。
    **/
}

func readSum()  int{
    mutex.RLock()
    defer mutex.RUnlock()
    b:=sumMutex
    return b
    
}
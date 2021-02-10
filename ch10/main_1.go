package main
import "fmt"
import "time"
import "sync"
// import "errors"
// import (
//     "io/ioutil"
//     "os"
// )

func main(){
    /**
        协程如何退出 
        一个协程启动后，大部分情况需要等待里面的代码执行完毕，然后协程会自行退出。但是如果有一种情景，需要让协程提前退出怎么办呢？
        如果需要让监控狗停止监控、退出程序，一个办法是定义一个全局变量，其他地方可以通过修改这个变量发出停止监控狗的通知。然后在协程中先检查这个变量，如果发现被通知关闭就停止监控，退出当前协程。
        但是这种方法需要通过加锁来保证多协程下并发的安全，基于这个思路，有个升级版的方案：用 select+channel 做检测，如下面的代码所示：
        select+channel 的方式改造的 watchDog 函数，实现了通过 channel 发送指令让监控狗停止，进而达到协程退出的目的。以上示例主要有两处修改，具体如下：

        为 watchDog 函数增加 stopCh 参数，用于接收停止指令；

        在 main 函数中，声明用于停止的 stopCh，传递给 watchDog 函数，然后通过 stopCh<-true 发送停止指令让协程退出。
    **/
   var wg sync.WaitGroup
   wg.Add(1)
   stopCh := make(chan bool) //用来停止监控狗
   go func (){
       defer wg.Done()
    //    watchDog("【监控1】")
       watchDog(stopCh,"【监控1】")
   }()
   time.Sleep(5 * time.Second) //先让监控狗监控5秒
   stopCh <- true //发停止指令
   wg.Wait()

   /**
    初识 Context
    select+channel 让协程退出的方式比较优雅，但是如果我们希望做到同时取消很多个协程呢？如果是定时取消协程又该怎么办？这时候 select+channel 的局限性就凸现出来了，即使定义了多个 channel 解决问题，代码逻辑也会非常复杂、难以维护。

    要解决这种复杂的协程问题，必须有一种可以跟踪协程的方案，只有跟踪到每个协程，才能更好地控制它们，这种方案就是 Go 语言标准库为我们提供的 Context，也是这节课的主角。
   **/
}
func watchDog(stopCh chan bool,name string){
    //开启for select循环，一直后台监控
    // for {
    //     select{
    //     default:
    //         fmt.Println(name,"正在监控.....")
    //     }
    //     time.Sleep(1*time.Second)
        
    // }
    for{
        select{
            case <-stopCh:
                fmt.Println(name,"停止指令已收到，马上停止")
                return
            default:
                fmt.Println(name,"正在监控.....")
        }
        time.Sleep(1*time.Second)
    }
}

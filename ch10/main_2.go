package main
import "fmt"
import "time"
import "sync"
import "context"

func main(){
  /**
    什么是 Context
    一个任务会有很多个协程协作完成，一次 HTTP 请求也会触发很多个协程的启动，而这些协程有可能会启动更多的子协程，并且无法预知有多少层协程、每一层有多少个协程。
    如果因为某些原因导致任务终止了，HTTP 请求取消了，那么它们启动的协程怎么办？该如何取消呢？因为取消这些协程可以节约内存，提升性能，同时避免不可预料的 Bug。
    Context 就是用来简化解决这些问题的，并且是并发安全的。Context 是一个接口，它具备手动、定时、超时发出取消信号、传值等功能，主要用于控制多个协程之间的协作，尤其是取消操作。一旦取消指令下达，那么被 Context 跟踪的这些协程都会收到取消信号，就可以做清理和退出操作。

    Context 接口只有四个方法，下面进行详细介绍，在开发中你会经常使用它们，你可以结合下面的代码来看。

    Context 是一种非常好的工具，使用它可以很方便地控制取消多个协程。在 Go 语言标准库中也使用了它们，比如 net/http 中使用 Context 取消网络的请求。
    要更好地使用 Context，有一些使用原则需要尽可能地遵守。

    Context 不要放在结构体中，要以参数的方式传递。

    Context 作为函数的参数时，要放在第一位，也就是第一个参数。

    要使用 context.Background 函数生成根节点的 Context，也就是最顶层的 Context。

    Context 传值要传递必须的值，而且要尽可能地少，不要什么都传。

    Context 多协程安全，可以在多个协程中放心使用。

    以上原则是规范类的，Go 语言的编译器并不会做这些检查，要靠自己遵守。
    **/
//    var wg sync.WaitGroup
//    wg.Add(1)
//    ctx,stop:=context.WithCancel(context.Background())
//    go func ()  {
//       defer wg.Done()
//       watchDog(ctx,"【监控1】") 
//    }()
//    time.Sleep(5* time.Second)
//    stop()
//    wg.Wait()
    
    /**
    Context 树
    Go 语言提供了函数可以帮助我们生成不同的 Context，通过这些函数可以生成一颗 Context 树，这样 Context 才可以关联起来，父 Context 发出取消信号的时候，子 Context 也会发出，这样就可以控制不同层级的协程退出。
    从使用功能上分，有四种实现好的 Context。

    空 Context：不可取消，没有截止时间，主要用于 Context 树的根节点。
    可取消的 Context：用于发出取消信号，当取消的时候，它的子 Context 也会取消。
    可定时取消的 Context：多了一个定时的功能。
    值 Context：用于存储一个 key-value 键值对。

    Go 语言提供的四个函数
    WithCancel(parent Context)：生成一个可取消的 Context。
    WithDeadline(parent Context, d time.Time)：生成一个可定时取消的 Context，参数 d 为定时取消的具体时间。
    WithTimeout(parent Context, timeout time.Duration)：生成一个可超时取消的 Context，参数 timeout 用于设置多久后取消
    WithValue(parent Context, key, val interface{})：生成一个可携带 key-value 键值对的 Context。
    **/
    // 使用 Context 取消多个协程
    var wg sync.WaitGroup
    wg.Add(4) //记得这里要改为4，原来是3，因为要多启动一个协程
    ctx,stop:=context.WithCancel(context.Background())
    go func ()  {
       defer wg.Done()
       watchDog(ctx,"【监控1】") 
    }()
    go func ()  {
        defer wg.Done()
        watchDog(ctx,"【监控2】") 
    }()
    go func ()  {
        defer wg.Done()
        watchDog(ctx,"【监控3】") 
    }()
   
    // 以上示例中的 Context 没有子 Context，如果一个 Context 有子 Context，在该 Context 取消时会发生什么呢？
    /**
        Context 传值
        Context 不仅可以取消，还可以传值，通过这个能力，可以把 Context 存储的值供其他协程使用。我通过下面的代码来说明：
    **/
    valCtx:= context.WithValue(ctx, "userId",2)
    go func(){
        defer wg.Done()
        getUser(valCtx)
    }()

    time.Sleep(5* time.Second)
    stop()
    wg.Wait()
}
//Context 接口的四个方法中最常用的就是 Done 方法，它返回一个只读的 channel，用于接收取消信号。当 Context 取消的时候，会关闭这个只读 channel，也就等于发出了取消信号。
type Context interface{
    Deadline() (deadline time.Time,ok bool) //方法可以获取设置的截止时间，第一个返回值 deadline 是截止时间，到了这个时间点，Context 会自动发起取消请求，第二个返回值 ok 代表是否设置了截止时间。
    Done() <-chan struct{}//方法返回一个只读的 channel，类型为 struct{}。在协程中，如果该方法返回的 chan 可以读取，则意味着 Context 已经发起了取消信号。通过 Done 方法收到这个信号后，就可以做清理操作，然后退出协程，释放资源。
    Err() error //方法返回取消的错误原因，即因为什么原因 Context 被取消。
    Value(key interface{})interface{}  //方法获取该 Context 上绑定的值，是一个键值对，所以要通过一个 key 才可以获取对应的值。
}

func watchDog(ctx context.Context,name string){
    //开启for select循环，一直后台监控
    for {
        select{
        case <-ctx.Done():
            fmt.Println(name,"停止指令已收到，马上停止")
            return 
        default:
            fmt.Println(name,"正在监控....")
        }
        time.Sleep(1 * time.Second)
    }
}
func getUser(ctx context.Context){
    for{
        select {
        case <-ctx.Done():
            fmt.Println("【获取用户】","协程退出")
            return
        default:
            userId:=ctx.Value("userId")
            fmt.Println("【获取用户】","用户ID为",userId)
            time.Sleep(1 * time.Second)
        }
        
    }
}
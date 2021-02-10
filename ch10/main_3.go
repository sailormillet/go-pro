package main
import "fmt"
import "time"
import "sync"
import "context"
/**
假如一个用户请求访问我们的网站，如何通过 Context 实现日志跟踪？
要想跟踪一个用户的请求，必须有一个唯一的 ID 来标识这次请求调用了哪些函数、执行了哪些代码，然后通过这个唯一的 ID 把日志信息串联起来。这样就形成了一个日志轨迹，也就实现了用户的跟踪，于是思路就有了。

在用户请求的入口点生成 TraceID。

通过 context.WithValue 保存 TraceID。

然后这个保存着 TraceID 的 Context 就可以作为参数在各个协程或者函数间传递。

在需要记录日志的地方，通过 Context 的 Value 方法获取保存的 TraceID，然后把它和其他日志信息记录下来。

这样具备同样 TraceID 的日志就可以被串联起来，达到日志跟踪的目的。

以上思路实现的核心是 Context 的传值功能。

**/


func main(){
    ctx,stop := context.WithCancel(context.Background())
    valCtx := context.WithValue(ctx,"userId",2)
    
}

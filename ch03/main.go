package main
import "fmt"

func main(){
    //if、for、switch 逻辑语句
   if i:=6; i>10{
    fmt.Println("i>10")
   }else if i>5 && i<=10{
    fmt.Println("5<i<=10")
   }else{
    fmt.Println("i<=5")
   }

   //switch
   switch i:=6; {
    case i>10:
        fmt.Println("i>10")
    case i>5 && i<=10:
        fmt.Println("5<i<=10")
    case i<=5:
        fmt.Println("i<=5")
   }
   switch j:=1;j {
    case 1:
        fallthrough
    case 2:
        fmt.Println("1,2")
    default:
        fmt.Println("not")
    }
    switch 2>1 {
    case true:
        fmt.Println("2>1")
    case false:
        fmt.Println("2<=1")
    }
    //for
    sum:=0
    for i:=1;i<=100;i++{
        sum+=i
    }
    fmt.Println("the sum is",sum)

    sum2:=0
    i:=1
    for i<=100{
        sum2+=i
        i++
    }
    
    fmt.Println("the sum is",sum2)
    //continue
    sumc := 0 
    for i:=1;i<=100;i++{
        if i%2 != 0  {
            continue
        }
        fmt.Println(i)
        sumc += i
    }
    
    fmt.Println("the sumc is:", sumc)
}


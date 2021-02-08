package main
import "fmt"
import "unicode/utf8"
//array,slice,map
func main(){
    //初始化
    array:=[5]string{"a","b","c","d","e"}
    fmt.Println(array[2])
    //初始化
    array1:=[...]string{"a","b","c","d"}
    fmt.Println(array1)
    //特定索引元素初始化
    array2:=[5]string{2:"a",3:"b"}
    fmt.Println(array2)
    //for 循环
    for i:=0;i<5;i++{
        fmt.Printf("数组索引:%d,对应值:%s\n", i, array2[i])
    }
    //range 循环
    for i,v:=range array2{
        fmt.Printf("数组索引:%d,对应值:%s\n", i, v)
    }
    //如果返回的值用不到，可以使用 _ 下划线丢弃:
    for _,v:=range array{
        fmt.Printf("对应值:%s\n", v)
    }

    //Slice（切片） slice:=array[start:end]包含索引start，但是不包含索引end，修改slice的值也会改变原始值
    //array[:4] 等价于 array[0:4]
    //array[1:] 等价于 array[1:5]
    //array[:] 等价于 array[0:5]
    slice:=array[2:5]
    fmt.Println(slice)
    fmt.Println(array)
    slice[1] ="f"
    fmt.Println(array)

    //make 函数切片声明
    // slice1:=make([]string,4)创建的切片 []string 长度是 4
    slice1:=make([]string,4,8)//创建的切片 []string 容量为 8,切片的容量不能比切片的长度小
    fmt.Println(slice1)
    //切片字面量的方式声明和初始化
    slice2:=[]string{"a","b","c","d","e"}
    fmt.Println(len(slice2),cap(slice2))
    //Append 函数对一个切片追加元素
    // slice3:=append(slice,"f")
    // slice3:=append(slice,"f","g")
    slice3:=append(slice,slice2...)
    fmt.Println(slice3)
    //在创建新切片的时候，最好要让新切片的长度和容量一样，这样在追加操作的时候就会生成新的底层数组，从而和原有数组分离，就不会因为共用底层数组导致修改内容的时候影响多个切片。


    //Map 使用make声明初始化 
    nameAgeMap := make(map[string]int)
    nameAgeMap["Jay"]=12
    fmt.Println(nameAgeMap)
    //Map 使用字面量声明初始化 必须有{}
    nameAgeMap1 := map[string]int {}
    nameAgeMap1["Amy"]=28
    fmt.Println(nameAgeMap1)
    // Map 获取和删除
    //Map 获取 如果 Key 不存在，返回的 Value 是该类型的零值
    fmt.Println(nameAgeMap1["Amy"])
    //map 的 [] 操作符可以返回两个值,第一个值是对应的 Value。第二个值标记该 Key 是否存在，如果存在，它的值为 true。
    age,ok:=nameAgeMap1["Amy"]
    if ok{
        fmt.Println("age",age)
    }
    // Map 删除 delete 有两个参数：第一个参数是 map，第二个参数是要删除键值对的 Key。
    delete(nameAgeMap1,"Amy")
    fmt.Println(nameAgeMap1)

    //遍历 Map for range map 的时候，也可以使用一个值返回。使用一个返回值的时候，这个返回值默认是 map 的 Key
    for k,v:=range nameAgeMap{
        fmt.Println("Key is",k,",Value is",v)
    }
    fmt.Println(len(nameAgeMap))


    //String 和 []byte
    s:="abc你好"
    bs:=[]byte(s)
    fmt.Println(bs)
    fmt.Println(s[0],s[1],s[5])
    fmt.Println(len(s))
    fmt.Println(utf8.RuneCountInString(s))
    for i,r:=range s{
        fmt.Println(i,r)
    }



    //创建二位数组
    array2n := [5][4] int {}
    fmt.Println(array2n)
}


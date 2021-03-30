// 工厂函数创建一个接口，并隐藏其内部实现，如下代码所示：

//工厂函数，返回一个error接口，其实具体实现是*errorString
func New(text string) error {
    return &errorString{text}
}
//结构体，内部一个字段s，存储错误信息
type errorString struct {
    s string
}
//用于实现error接口
func (e *errorString) Error() string {
    return e.s
}
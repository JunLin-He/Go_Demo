package main 

import (
	"fmt"
)

type A struct {
	B
	C
	Name string
}

type B struct {
	Name string
}

type C struct {
	Name string
}

func main()  {
	a := A{Name: "A", B: B{Name: "B"}, C: C{Name: "C"}}
	fmt.Println(a.Name, a.B.Name)	// 此处a.Name 输出的是A的name，按照优先级进行
}

// Output: A B
// 当最高级别(即当前结构的字段)有该字段时，则取该字段；若无，才会往下去找下一层的选中字段
// 注意，如果同一个结构内的两个嵌套结构都包含了名字相同的字段，则会报错；因为B、C级别相同，
//     所以编译器无法选择，从而导致出错

/* The general coding sequence of Go program*/

// The current program's package name
// 当前程序包名
package main

// Import the other package, this sentense allow variety of package
// Attentions: if the imported package is not uesed, it will be errored.
// import std "fmt"
// 导入其他包
import (
	std "fmt" // The alias of package "fmt"
)

// The definition of the const variable
//const PI = 3.14
// 常量的定义
const (
	PI = 3.14
)

// Declare the globle variable and assign the value
//var name = "gopher"
// 全局变量的声明与赋值
var (
	name = "gopher"
)

// The declare of the general type
//type newType int
// 一般类型的声明
type (
	newType int
)

// Declare the structure
// 结构的声明
type gopher struct{}

// Declare the interface
// 接口的声明
type golang interface{}

// Boost by main function which is the entrance of the program
// 由main函数作为程序入口点启动
func main() {
	std.Println("Hello world!")
}

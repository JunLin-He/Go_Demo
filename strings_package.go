package main

import (
	std "fmt"
	"strings"
)

func main() {
	var str string = "go_lang"
	std.Printf("T/F? Does the string \"%s\" have prefix \"%s\"? ", str, "go")
	std.Printf("%t\n", strings.HasPrefix(str, "go"))
	std.Printf("T/F? Does the string \"%s\" contains \"%s\"? ", str, "-")
	std.Printf("%t\n", strings.Contains(str, "-"))
	str_new := strings.Replace(str, "go", "python", 1)
	std.Printf("Origin strings: \"%s\", after replace: \"%s\"\n", str, str_new)
	std.Printf("Number of 'n' in \"%s\" is: %d\n", str_new, strings.Count(str_new, "n"))
}

/*
	T/F? Does the string go_lang have prefix go? true
	T/F? Does the string go_lang contains "-"? true
	Origin strings: "go_lang", after replace: "python_lang"
	Number of 'n' in "python_lang" is: 2
*/

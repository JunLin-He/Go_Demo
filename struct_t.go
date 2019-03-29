package main

import "fmt"

type Count int // Create a custom type

func (count *Count) Increment()  { *count++ }
func (count *Count) Decrement()  { *count-- }
func (count Count) IsZero() bool { return count == 0 }

type Part struct { // Create a custom type based on struct type Part
	stat  string
	Count // Anonymous field
}

func (part Part) IsZero() bool {
	return part.Count.IsZero() && part.stat == ""
}

func (part Part) String() string {
	return fmt.Sprintf("<<%s, %d>>", part.stat, part.Count)
}

func main() {
	var i Count = -1
	fmt.Printf("Start \"Count\" test:\nOrigin value of count: %d\n", i)
	i.Increment()
	fmt.Printf("Value of count after increment: %d\n", i)
	fmt.Printf("Count is zoro t/f? : %t\n\n", i.IsZero())
	fmt.Println("Start: \"Part\" test:")
	part := Part{"232", 0}
	fmt.Printf("Part: %v\n", part)
	fmt.Printf("Part is zero t/f? : %t\n", part.IsZero())
	fmt.Printf("Count in Part is zero t/f?: %t\n", part.Count.IsZero())

}

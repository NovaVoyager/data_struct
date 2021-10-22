package main

import "fmt"

func main() {
	linked := Constructor()
	linked.AddAtIndex(0, 10)
	linked.AddAtIndex(0, 20)
	linked.AddAtIndex(1, 30)
	linked.AddAtIndex(1, 40)
	linked.AddAtIndex(1, 50)
	linked.DeleteAtIndex(2)
	linked.DeleteAtIndex(0)
	linked.DeleteAtIndex(2)

	fmt.Println(linked)
}

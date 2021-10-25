package main

import "fmt"

func main() {
	test()
}

func test() {
	size := 3
	que := Constructor(size)

	fmt.Println("1", que.Front())
	fmt.Println("2", que.Rear())
	fmt.Println("3", que.IsEmpty())
	fmt.Println("4", que.IsFull())

	fmt.Println("5", que.EnQueue(1))
	fmt.Println("5-1", que.IsFull())
	fmt.Println("6", que.EnQueue(2))
	fmt.Println("7", que.EnQueue(3))
	fmt.Println("8", que.EnQueue(4))

	fmt.Println("9", que.IsEmpty())
	fmt.Println("10", que.IsFull())

	fmt.Println("11", que.Front())
	fmt.Println("12", que.Rear())

	fmt.Println("13", que.DeQueue())
	fmt.Println("14", que.DeQueue())
	fmt.Println("15", que.DeQueue())
	fmt.Println("16", que.DeQueue())

	fmt.Println("17", que.IsEmpty())
	fmt.Println("18", que.IsFull())

}

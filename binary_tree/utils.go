package main

import (
	"bufio"
	"os"
)

//readIn 读取输入
func readIn() string {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	return string(char)
}

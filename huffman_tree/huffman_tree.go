package main

import "fmt"

//HTNode 哈夫曼树定义
type HTNode struct {
	weight                 int
	parent, lchild, rchild int
}

type HTNodes []HTNode

type ElemType string

func String(this HTNodes) {
	fmt.Printf("结点\tweight\tparent\tlchild\trchild\n")
	for i := 1; i < len(this); i++ {
		fmt.Printf("%d\t\t%d\t\t%d\t\t%d\t\t%d\n", i, this[i].weight, this[i].parent, this[i].lchild, this[i].rchild)
	}
}

func createTree(n int, weights []int) HTNodes {
	if n <= 1 || len(weights) == 0 || len(weights) != n {
		return nil
	}

	//哈夫曼树节点数 2n-1
	m := 2*n - 1
	//第一个元素不使用,需分配 2n 个节点数的空间
	ht := make([]HTNode, 2*n, 2*n)
	//输入n个节点的权值
	for i := 1; i <= n; i++ {
		ht[i].weight = weights[i-1]
	}
	fmt.Println("初态")
	String(ht)
	fmt.Println("-------------------------------------------------------------------")
	//创建哈夫曼树
	for i := n + 1; i <= m; i++ {
		s1, s2 := Select(ht, i-1)
		ht[i].lchild, ht[i].rchild, ht[i].weight = s1, s2, ht[s1].weight+ht[s2].weight
		ht[s1].parent, ht[s2].parent = i, i
	}
	fmt.Println("终态")
	String(ht)
	return ht
}

func Select(ht []HTNode, n int) (s1, s2 int) {
	min := 0
	for i := 1; i <= n; i++ {
		if ht[i].parent == 0 {
			min = i
			i = n + 1
		}
	}
	for i := 1; i <= n; i++ {
		if ht[i].parent == 0 {
			if ht[i].weight < ht[min].weight {
				min = i
			}
		}
	}
	s1 = min

	for i := 1; i <= n; i++ {
		if ht[i].parent == 0 && i != s1 {
			min = i
			i = n + 1
		}
	}
	for i := 1; i <= n; i++ {
		if ht[i].parent == 0 && i != s1 {
			if ht[i].weight < ht[min].weight {
				min = i
			}
		}
	}
	s2 = min
	return s1, s2
}

//huffmanCode 哈夫曼编码
//chars 保存一个顺序与哈夫曼数一致的字符数组，第一个元素不使用，下标从1开始
//返回值，返回每个char对应的哈夫曼编码
func huffmanCode(HT HTNodes, n int, chars []ElemType) map[ElemType][]int {
	charWithCode := make(map[ElemType][]int, n) //保存最终结果
	for i := 1; i <= n; i++ {
		tmp := make([]int, 0) //临时变量保存编码
		f := HT[i].parent
		c := i
		for f != 0 {
			if HT[f].lchild == c {
				tmp = append(tmp, 0)
			} else {
				tmp = append(tmp, 1)
			}
			c = f
			f = HT[f].parent
		}

		//因为是逆向回溯，所以编码需要倒序取出
		tmpLastIndex := len(tmp) - 1
		code := make([]int, 0, len(tmp))
		for j := tmpLastIndex; j >= 0; j-- {
			code = append(code, tmp[j])
		}
		charWithCode[chars[i]] = code
	}
	printHuffCode(charWithCode)
	return charWithCode
}

func printHuffCode(htCode map[ElemType][]int) {
	fmt.Printf("\nhuffman code:\n")
	for k, v := range htCode {
		fmt.Printf("char:%s --> code:%v\n", k, v)
	}
}

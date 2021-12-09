package main

type KeyType int

//ElemType 定义元素类型
type ElemType struct {
	Key KeyType
}

//SSTable 定义顺序表
type SSTable struct {
	R      []ElemType
	Length int
}

//SearchSeq 顺序查找
func (this *SSTable) SearchSeq(key KeyType) int {
	for i, v := range this.R {
		if v.Key == key {
			return i
		}
	}
	return 0
}

//SearchBin 二分查找
func (this *SSTable) SearchBin(key KeyType) int {
	left := 1
	right := len(this.R)
	for left <= right {
		mid := (left + right) / 2
		if this.R[mid].Key == key {
			return mid
		}
		if this.R[mid].Key > key {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0
}

//SearchBinRecursion 递归二分查找
func (this *SSTable) SearchBinRecursion(key KeyType, left, right int) int {
	if left > right {
		return 0
	}
	mid := (left + right) / 2
	if this.R[mid].Key == key {
		return mid
	}
	if this.R[mid].Key > key {
		return this.SearchBinRecursion(key, left, mid-1)
	} else {
		return this.SearchBinRecursion(key, mid+1, right)
	}
}

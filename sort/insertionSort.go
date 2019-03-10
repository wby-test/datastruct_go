package main

import "fmt"

//自己实现1（这好像是个冒泡）

func insertSort1(val []int) {
	if len(val) <= 1 {
		return
	}
	for i := 1; i < len(val); i++ {
		for j := 0; j < len(val); j++ {
			if val[i] < val[j] {
				//temp := val[i]
				//val[i] = val[j]
				//val[j] = temp

				val[i] , val[j] = val[j], val[i]
			}
		}
	}
}

// 网上写法
func InsertionSort(s []int) {
	n := len(s)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && s[j] < s[j - 1]; j-- {
			swap(s, j, j - 1)
		}
	}
}

func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

//标准写法

func stdInsertSort(val []int) {
	if len(val) <= 1 {
		return
	}

	for i := 1; i < len(val); i++ {
		temp := val[i]
		j := i - 1
		for  ; j >= 0; j-- {
			if val[j] > temp {
				val[j + 1] = val[j]
			}else {
				break
			}
		}
		val[j + 1] = temp //插入数据
	}
}

func main() {
	val := []int{3, 2, 4, 7, 5, 1, 10, 6}
	//stdInsertSort(val)
	insertSort1(val)
	fmt.Println(val)
}
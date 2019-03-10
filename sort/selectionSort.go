package main

import "fmt"

//自己实现

func selectionSort(val []int) {
	if len(val) <= 1 {
		return
	}

	for i := 0; i < len(val)-1; i++ {
		var temp int
		temp = val[i]
		j := i
		k := i
		for ; j < len(val)-1; j++ {
			if temp > val[j+1] {
				temp = val[j+1]
				k = j + 1
			}
		}
		if k != i {
			val[k] = val[i]
			val[i] = temp
		}
	}
}

//网上实现
func swap(slice []int, i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func SelectionSort(s []int) {
	l := len(s) //以免每次循环判断都运算
	m := len(s) - 1
	for i := 0; i < m; i++ {
		k := i
		for j := i + 1; j < l; j++ {
			if s[j] < s[k] {
				k = j
			}
		}
		if k != i {
			swap(s, k, i)
		}
	}
}

func main() {
	val := []int{3, 2, 4, 7, 5, 1, 10, 6}
	//stdInsertSort(val)
	selectionSort(val)
	fmt.Println(val)
}

package main

import "fmt"

//自己实现第一次
func bubbleSort(val []int) {
	num := len(val)
	for num != 1 {
		for i := 0; i < num - 1; i++ {
			for j := i + 1; j < num; j++ {
				if val[i] > val[j] {
					temp := val[i]
					val[i] = val[j]
					val[j] = temp
				}
			}
		}
		num--
	}
}

//自己实现第二次
func bubbleSort2(val []int) {
	if len(val) <= 1 {
		return
	}
	for i :=0; i < len(val); i++ {
		for j := 0; j < len(val) - i -1; j++ {
			if val[j] > val[j + 1] {
				temp := val[j]
				val[j] = val[j + 1]
				val[j + 1] = temp
			}
		}
	}
}



//标准实现
// 冒泡排序，a 表示数组，n 表示数组大小
func StdBubbleSort(val []int) {
	if len(val) <= 1 {
		return
	}

	num := len(val)
	for i := 0; i < num; i++ {
		flag := 0
		for j := 0; j < num - i - 1; j++ {
			if val[j] > val[j+1] {
				temp := val[j]
				val[j] = val[j+1]
				val[j+1] = temp
				flag = 1 //冒泡优化
			}
		}
		if flag == 0 {
			break
		}
	}
}


func main() {
	val := []int{3, 2, 4, 7, 5, 1, 10, 6}
	StdBubbleSort(val)
	fmt.Println(val)
}
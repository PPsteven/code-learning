package main

import "fmt"

func CutAndAppend(a []int, elem int) []int {
	for i := 0; i < len(a); i++ {
		if a[i]	== elem {
			a = append(a[:i], a[i+1:]...)
			i-- // NOTE: the length of array decrease
		}
	}
	return a
}

func NewSlice(a []int, elem int) []int {
	newA := make([]int, 0, len(a))
	for _, v := range a {
		if v != elem {
			newA = append(newA, v)
		}
	}
	return newA
}

func ModifyOriginSliceByNewTempSlice(a []int, elem int) []int {
	tmp := a[:0] // NOTE: slice 是共享数组底层, tmp 共享了a的底层数组, 优于上面的操作
	for _, v := range a {
		if v != elem {
			tmp = append(tmp, v)
		}
	}
	return tmp
}

// Best Solution
func ModifyOriginSlice(a []int, elem int) []int {
	i := 0
	for _, v := range a {
		if v != elem {
			a[i] = v
			i++
		}
	}
	return a[:i]
}


func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(ModifyOriginSlice(a, 1))
}

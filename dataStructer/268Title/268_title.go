package main

import "fmt"

func main() {
	cc := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	//dd:=[]int{3,0,1}
	//ee:=[]int{0}
	title(cc)
}
func title(nums []int) int {
	var result int

	//m := make(map[int]int, 0)
	//for i,k := range nums {
	//	m[k]=i
	//}
	//for i := 0;i<=len(nums);i++{
	//	if _,v := m[i];!v{
	//		result=i
	//	}
	//}

	for i, k := range nums {
		result ^= k ^ i
		//fmt.Println(i,"------------",k)
		fmt.Println(result)
	}
	//fmt.Println(result)
	//fmt.Println(0^1)
	fmt.Println(result ^ len(nums))
	return result ^ len(nums)
}

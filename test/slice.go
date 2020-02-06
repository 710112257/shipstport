package main

import "fmt"

func main(){
	var arr1 =[]int{1,2,3,4,5,6}
	fmt.Println(arr1)
	arr2:=arr1[:3]
	fmt.Println(arr2)
	arr2[2]=99
	fmt.Println(arr2)
	fmt.Println(arr1)
	arr2=append(arr2,11)
	fmt.Println(arr2)
	fmt.Println(arr1)

	arr3:=arr2[:]
	arr3[2]=77
	fmt.Println(arr3)
	fmt.Println(arr2)
	fmt.Println(arr1)

}
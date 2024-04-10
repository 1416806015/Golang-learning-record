package main

import "fmt"

func main() {
	var arr [][]int = [][]int{{1, 2, 3}, {8, 9, 7}, {4, 5, 6}}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Println()
	}
	//2
	for key, value := range arr {
		for k, v := range value {
			fmt.Printf("arr[%v][%v]=%v\t", key, k, v)
		}
		fmt.Println()
	}
}

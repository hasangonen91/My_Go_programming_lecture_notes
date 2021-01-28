package main

import "fmt"

func main() {

	var sayilar [2][3]int
	sayilar[0][0] = 10
	sayilar[0][1] = 20
	sayilar[0][2] = 30
	sayilar[1][0] = 40
	sayilar[1][1] = 50
	sayilar[1][2] = 60

	fmt.Println(sayilar)
	fmt.Println(sayilar[1][2])

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(sayilar[i][j])
			fmt.Print(" ")

		}
		fmt.Println()
	}
	//fmt.Println(sayilar[1][1])
}

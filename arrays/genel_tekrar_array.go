package main

import "fmt"

func Demo1() {
	var sayilar [5]int
	sayilar2 := [5]int{1, 2, 3, 4, 5}
	var sehirler [2]string

	sayilar[2] = 50

	sehirler[0] = "HAKKARI"

	fmt.Println(sayilar[2])
	fmt.Println(sehirler[0])

	enbuyuk := sayilar2[0]
	for i := 0; i < len(sayilar2); i++ {
		if sayilar2[i] > enbuyuk {
			enbuyuk = sayilar2[i]
		}
	}
	fmt.Println("en buyuk sayi: ", enbuyuk)

	var boyutlusayilar [2][3]int //satir ve sutun

	boyutlusayilar[0][0] = 10
	boyutlusayilar[0][1] = 20
	boyutlusayilar[0][2] = 30
	boyutlusayilar[1][0] = 40
	boyutlusayilar[1][1] = 50
	boyutlusayilar[1][2] = 60

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(boyutlusayilar[i][j])
			fmt.Print(" ")
		}
		fmt.Println(" ")
	}

}

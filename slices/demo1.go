package slices

import "fmt"

func Demo1() {

	isimler := make([]string, 3)

	fmt.Println(isimler)
	isimler[0] = "hasan"
	isimler[1] = "huseyin"
	isimler[2] = "ahmet"
	//isimler[3]="harun"
	isimler = append(isimler, "irem")//append yeni eleman ekliyor
	fmt.Println(isimler)
	fmt.Println(len(isimler))

}

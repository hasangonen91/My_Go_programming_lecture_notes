package slices

import "fmt"

func Demo2() {

	sehirler := []string{"ankara", "istanbul", "izmir"}
	fmt.Println(sehirler)
	sehirlerKopya := make([]string, len(sehirler)) //bu yepyeni dizidir
	fmt.Println(sehirlerKopya)
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)
	//fmt.Println("EKLE:")
	//var ekle string
	//fmt.Scanln(&ekle)
	//sehirler = append(sehirler, ekle)
	sehirler = append(sehirler, "ekle")
	fmt.Println(sehirler)
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)

	fmt.Println(sehirler[0:5])
	fmt.Println(sehirler[1:3])
	fmt.Println(sehirler[1:])
	fmt.Println(sehirler[0:])
	fmt.Println(sehirler[:2])
}

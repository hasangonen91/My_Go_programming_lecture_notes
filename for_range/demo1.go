package for_range

import "fmt"

func Demo1() {
	sehirler := []string{"ankara", "istanbul", "izmir"}

	for i := 0; i < len(sehirler); i++ {
		fmt.Println(sehirler[i])

	}
// 2 turlu for vardır ustteki ve alttaki

	for _, sehir := range sehirler {
		/*
		for i, sehir := range sehirler {
			bu sekilde de yazatdık ama o zaman indexe deger
			eklemeliydik 
			fmt.Println(i) 
		*/
		fmt.Println(sehir)
	}

}

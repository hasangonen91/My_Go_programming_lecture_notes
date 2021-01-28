package string_functions

//alias
import (
	"fmt"
	s "strings"
)

//case sensitive
func Demo1() {
	isim := "Hasan"
	fmt.Println(s.Count(isim, "a"))    //kac adet var
	fmt.Println(s.Contains(isim, "n")) //varmı yokmu dondurur
	fmt.Println(s.Index(isim, "a"))    //ilk gordugunun indexini dondurur kacinci indexte oldugunu dondurur
	fmt.Println(s.Index(isim, "k"))    //bulamadiginda -1 dondurur
	sonuc := s.Index(isim, "a")

	if sonuc != -1 {
		fmt.Println("a harfi var")
	} else {
		fmt.Println("a harfi yok")
	}

	fmt.Println(s.ToLower(isim)) //kucuk harfe cevir
	fmt.Println(s.ToUpper(isim)) // buyuk harfe cevir

}

package string_functions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	isim := "Hasan"
	fmt.Println(s.HasPrefix(isim, "Has")) //sart blogunda boolean olarak kullanilabilir
	fmt.Println(s.HasSuffix(isim, "an"))  // an ile bitiyormu kontrol ediyor bitis kismini kontrol ediyor
	fmt.Println(s.Index(isim, "sa"))      //sirasini soyluyor
	harfler := []string{"h", "a", "s", "a", "n"}
	sonuc := s.Join(harfler, "*") // harfleri ayirdi arasina istedigini koyabilirsin
	fmt.Println(sonuc)

	fmt.Println(s.Replace(sonuc, "*", "+", -1)) //-1 komple degistir yani kalmasin demek programlama dilinde
	// 1,2,3.. vs koyarsan o kadarini degistir demektir

	fmt.Println(s.Split(sonuc, "*")) //* lar silindi ona göre ayirdi
	fmt.Println(s.Split(sonuc, "-")) //ayracı bulamadi tek bir array yaptı
	fmt.Println(s.Repeat(sonuc, 5))  //5 defa tekrarla dedik
}

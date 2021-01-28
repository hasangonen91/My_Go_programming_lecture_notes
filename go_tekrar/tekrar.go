package go_tekrar

import (
	"fmt"
	"time"
)

func Tekrar() {
	// diziler
	// var a = []int{1, 2, 3, 4} //[6] vb. yaparsak ekstradan 2 sayi eklenir ama 0 0 olarak eklenir
	// fmt.Println(a)
	// var b = a[:3] // 0:3 ister böyle ister :3 ya da 1:3 aralarindaki sayialari alirlar
	// fmt.Println(b)
	// a[1] = 25 // a dizisinin 1.degerini 2 yerine 25 yaptik
	// fmt.Println(a)

	//degiskenler
	// var sayi1 int = 30
	// var sayi2 = 10
	// sayi3:= 15

	//sabitler

	// const sabit int = 10

	// sabit =20  //yeni deger atasan bile degismez
	//fmt.Println(sabit)

	//yazdir2("hasan")

	//fmt.Println(yazdir3("hasan"))
	// fmt.Println(islem(5, 10))
	// toplam := islem(5, 10)
	// fmt.Println(toplam)
	// toplam, carpim := islem(5, 10)//carpimi kullanmak isteöiyorsak _ alt cizgi koyariz
	// fmt.Println(toplam, carpim)
	// fmt.Println("selam")
	// fmt.Print("nasilsin")
	// fmt.Printf("%d , %f  ,  %s bu kisim turune gore basilir C gibi aynisi", 10, 20.5, "anladinmi")
	// giris := bufio.NewReader(os.Stdin)
	// yazi1, _ := giris.ReadString('\n')
	// fmt.Println(yazi1)  EKRANA Yazdigin sayiyiyi tekrar yaziyor "deneme" sonuc= deneme
	//if else sart bloklari
	// sayi1 := 5
	// if sayi1 == 10 {
	// 	fmt.Println("dogru")
	// } else if sayi1 == 5 {
	// 	fmt.Println("sayi 5 e esit")
	// } else {
	// 	fmt.Println("yanlis")
	// }

	// if sayi1 := 5; sayi1 == 10 {//bu sekilde de yazilabilir
	// 	fmt.Println("dogru")
	// } else if sayi1 == 5 {
	// 	fmt.Println("sayi 5 e esit")
	// } else {
	// 	fmt.Println("yanlis")
	// }
	//Donguler
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 10; i++ {
	//    if i==5{
	// 	   break
	//    }else{
	// 	fmt.Println(i)
	//    }
	// }

	//for-Range dongusu
	// a := []string{"hasan", "ali", "murat"}
	// for i, d := range a {
	// 	fmt.Println(i, d)
	// }

	//isaretciler
	// a := "merhaba"
	// fmt.Println(&a)
	// b := &a
	// fmt.Println(*b)

	// nesne yonelimli programlama

	// kisi1 := insan{isim:"hasan",yas: 21}   1.kisim
	// fmt.Println(kisi1.isim, kisi1.yas)

	//2.kisim
	// calisan := calisan{insan{isim: "hasan", yas: 21}, 500}
	// fmt.Println(calisan.isim, calisan.yas, calisan.maas)
	// kisi1 := insan{"hasan", 21}
	// fmt.Println(kisi1)
	// kisi1.yazdir()

	// sekil1 := daire{5}
	// fmt.Println(sekil1.alan())

	// sekil2 := dortgen{3, 5}
	// fmt.Println(sekil2.alan())

	// var a sekil
	// sekil1 := daire{5}
	// a = &sekil1
	// alanHesabi(a)

	//map kismi

	// a := []int{1, 2, 3, 4}
	// fmt.Println(a)
	// b := make(map[string]sehir)
	// b["istanbul"] = sehir{
	// 	34,
	// 	18000000,
	// }
	// fmt.Println(b)

	//time komutlari
	// go merhaba()
	// dunya()

	//chanel
	c := make(chan bool)
	go merhaba(c)
	dunya()
	//a := <-c
	 <-c//bu sekilde de kullanabiliriz
	//fmt.Println(a)
}
func merhaba(c chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("merhaba")
	c <- true
}
func dunya() {
	time.Sleep(1 * time.Second)
	fmt.Println("dunya")
}

// func merhaba(){
// 	time.Sleep(1 * time.Second)
// 	fmt.Println("merhaba")
// }
// func dunya(){
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("dunya")
// }

// type sehir struct {
// 	plaka int
// 	nufus int
// }

/*
bu kisim interface

// type sekil interface { // hangi structi calismasini ayarlar
// 	alan() float64
// }

// type dortgen struct {
// 	uzunkenar float64
// 	kisakenar float64
// }
// type daire struct {
// 	yaricap float64
// }

// func (d dortgen) alan() float64 {
// 	return d.kisakenar * d.uzunkenar
// }

// func (d daire) alan() float64 {
// 	return 3.14 * (d.yaricap * d.yaricap)
// }

// func alanHesabi(i sekil) {
// 	fmt.Println(i.alan())
// }
*/

// func (i insan) yazdir() {
// 	fmt.Println(i)
// }

// //structlara ozel fonksiyonlar
// type insan struct {
// 	isim string
// 	yas  int
// }

// type insan struct {
// 	isim string
// 	yas  int
// }
// type calisan struct {
// 	insan
// 	maas int
// }

// type insan struct {//1.kisim
// 	isim string
// 	yas  int
// }

// fonksiyonlar

// func yazdir() {
// 	fmt.Println("merhaba")

// }
// func yazdir2(yazi string) {
// 	fmt.Println(yazi)
// }
// func yazdir3(yazi string) string{

// 	return yazi

// }

// func islem(a int, b int) int {
// 	return a + b
// }

// func islem(a int, b int) (int, int) { //func islem(a,b int) (int, int)// bu sekilde de yazilabilir
// 	/*
// 		topla = a+b
// 		carp=a*b
// 		return    bu sekilde de yazilabilir
// 	*/

// 	return a + b, a * b

//}

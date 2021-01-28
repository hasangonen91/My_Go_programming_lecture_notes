package main

import (
	"golesson/go_tekrar"
)

func main() {
	// 	//slices.Demo1()
	// 	//slices.Demo2()
	// 	//functions.Topla(2,3)
	// 	//functions.SelamVer("Hasan")
	// 	//alt cizgi 4.sünü calistirmak istemiyorum demek
	// 	/*sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 6)
	// 	//sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(10, 6)
	// 	fmt.Println("toplam:", sonuc1)
	// 	fmt.Println("cikarma:", sonuc2)
	// 	fmt.Println("carpma:", sonuc3)
	// 	//fmt.Println("bolme:", sonuc4)
	// 	*/
	// 	// functions.Demo3()
	// 	/*	fmt.Println(functions.ToplaVariadic(1, 4, 6, 3, 10))
	// 		fmt.Println(functions.ToplaVariadic(1, 4))
	// 		fmt.Println(functions.ToplaVariadic(1))
	// 		fmt.Println(functions.ToplaVariadic())

	// 		sayilar:=[]int{0,0,0,0,11}
	// 		fmt.Println(functions.ToplaVariadic(sayilar...))
	// 	*/

	// 	//maps.Demo1()
	// 	//for_range.Demo2()
	// 	//for_range.Demo3()
	// 	//sayi := 20
	// 	//pointers.Demo1(&sayi)
	// 	//fmt.Println("maindeki sayi : ",sayi)
	// 	//sayilar :=[]int{1,2,3}// arraylar degerleri ile calismaz referans ile adres ile calisir
	// 	// pointers.Demo2(sayilar) ctrl+k+c
	// 	// fmt.Println("Maindeki sayi : ",sayilar[0])
	// 	//structs.Demo2()
	// 	//pointers.Test()

	// 	// go goroutines.CiftSayilar()//calismaz cunku zaman kalmadi biter program "go" nun avantajı
	// 	// go goroutines.TekSayilar()

	// 	// time.Sleep(5 * time.Second)
	// 	// fmt.Println("main bitti")

	// 	// ciftSayiCn := make(chan int)
	// 	// tekSayiCn := make(chan int)
	// 	// go channels.CiftSayilar(ciftSayiCn)
	// 	// go channels.TekSayilar(tekSayiCn)

	// 	// ciftSayiToplam,tekSayiToplam := <-ciftSayiCn, <- tekSayiCn
	// 	// carpim:=ciftSayiToplam* tekSayiToplam
	// 	// fmt.Println("Carpim : ",carpim)

	// 	//interfaces.Demo1()
	// 	//interfaces.Demo2()

	// 	//defer_statement.B()
	// 	//defer_statement.Test()
	// 	//defer_statement.Demo3()
	// 	//error_handling.Demo1()
	// 	//interfaces.Demo3()
	// 	//error_handling.Demo2()
	// 	// fmt.Println(error_handling.TahminEt2(99))
	// 	//string_functions.Demo1()
	// 	//string_functions.Demo2()
	// 	//restful.Demo1()
	// 	//restful.Demo2()
	// 	//project.GetAllProducts()

	// 	// product, _ := project.AddProduct()
	// 	// fmt.Println(product)

	// 	//  products, _ := project.GetAllProducts()
	// 	//  for i:=0; i<len(products);i++{
	// 	//  	fmt.Println(products[i].ProductName)
	// 	//  }
	// 	for i := 1; i <= 10; i = i + 2 {
	// 		fmt.Println(i)
	// 	}

	go_tekrar.Tekrar()

}

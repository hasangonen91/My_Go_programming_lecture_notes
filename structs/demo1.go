package structs

import "fmt"

func Demo1() {

	fmt.Println(product{"Bilgisayar",500,"XYZ"})
	fmt.Println(product{"Bilgisayar",500,"abs"})
	fmt.Println(product{name :"Bilgisayar",unitPrice: 500})
	fmt.Println(product{name :"Bilgisayar"})

	



}

type product struct {
	name         string
	unitPrice    float64
	brand        string
	
}
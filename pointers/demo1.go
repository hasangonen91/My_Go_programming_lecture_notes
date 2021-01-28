package pointers

import "fmt"

func Demo1(sayi *int) {
	*sayi = *sayi + 1
	//fmt.Println("Demo da ki sayi : ",*sayi)
	//genel tekrar denemesidir:D
	fmt.Println("testte ki sayi : ",*sayi)

}
func Test(){
	sayi:=10
	Demo1(&sayi)


}

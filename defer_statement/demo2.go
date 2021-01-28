package defer_statement

import "fmt"

func Demo2(sayi int32) string {
	defer DeferFunc()

	if sayi%2 == 0 {
		fmt.Println("cift sayi calisti")
		return "cift sayidir"
	}
	if sayi%2 != 0 {
		return "tek sayidir"
	}

	return "belli degil"
}

func Test() {
	sonuc := Demo2(10)
	fmt.Println(sonuc)
}

func DeferFunc() {

	fmt.Println("bitti")

}

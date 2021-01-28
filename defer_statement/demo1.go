package defer_statement

import "fmt"

func A() {
	fmt.Println("a fonksiyonu calisti")
}

func C() {
	fmt.Println("c fonksiyonu calisti")
}

func D() {
	fmt.Println("d fonksiyonu calisti")
}

func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("b fonksiyonu calisti")
	//kodların calismasi alttan yukarıya dogru gidiyor

}

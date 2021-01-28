package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Ürün kaydedildi : ", p.productName)
	defer Log()
}

func Log() {
	fmt.Println("Logland")
}

func Demo3() {
	p := product{productName: "DELL", unitPrice: 6250}
	p = product{productName: "logitec", unitPrice: 250}

	fmt.Println("islem basarili")
	fmt.Println(p.productName)
	defer p.save()
}

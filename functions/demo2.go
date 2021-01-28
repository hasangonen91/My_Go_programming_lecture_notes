package functions

func DortIslem(sayi1 int, sayi2 int) (int, int, int, float32) {

	toplam := sayi1 + sayi2
	cikarma := sayi1 - sayi2
	carpma := sayi1 * sayi2
	bolme := float32(sayi1 / sayi2)

	return toplam, cikarma, carpma, bolme
}

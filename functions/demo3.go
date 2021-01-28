package functions

func ToplaVariadic(sayilar ...int) int{
	toplam:=0
	for i:=0;i<len(sayilar);i++{
		toplam += sayilar[i]

	} 
	return toplam


}
package channels
import (
    "fmt"
    "time"    
)
	
func CiftSayilar(ciftSayiCn chan int){
	toplam := 0

	for i:=0 ; i<=10 ; i+=2{
		toplam+=i
	   
	}
	fmt.Println("Cift sayilar calisiyor")
	time.Sleep(1*time.Second)
	fmt.Println(toplam)
	ciftSayiCn <- toplam

}

func TekSayilar(tekSayiCn chan int){
	toplam := 0
	for i:=1 ; i<=10 ; i+=2{
		toplam+=i
		
	}
	fmt.Println("Tek sayilar calisiyor")
	time.Sleep(1*time.Second)
	fmt.Println(toplam)
	tekSayiCn <- toplam
}



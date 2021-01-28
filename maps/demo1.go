package maps

import "fmt"

func Demo1() {
	//key-value
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "kitap"
	sozluk["pencil"] = "kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("eleman sayisi : ", len(sozluk))
	fmt.Println("Sozluk : ", sozluk)
	delete(sozluk, "book")
	fmt.Println("eleman sayisi : ", len(sozluk))
	fmt.Println("Sozluk : ", sozluk)

	//deger, varMi := sozluk["table"]
	//deger, varMi := sozluk["lamb"]//bu yok
	deger, varMi := sozluk["pencil"]

	fmt.Println(deger)
	fmt.Println("Listede olma durumu : ", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "mikrofon"}
	fmt.Println(sozluk2)

}

package structs

import "fmt"

type customer struct{
	firstName string
	lastName  string
	age int

}

func (c customer) save(){
	fmt.Println("Eklendi : ",c.firstName)
}

func (c customer) update(){
	fmt.Println("Guncellendi : ",c.firstName,c.lastName)
}



func Demo2(){
	c:= customer{firstName:"Hasan",lastName:"Gonen",age:35}
	c.save()
	c.update()
}


package main

import "fmt"

type Person struct{
	Name string
	Age int
}

func Greet(p Person) string{
	return "HI"+ p.Name
}

func main(){
	

	p:= Person{
		Name:"rayaan",
		Age:2,
	}

	fmt.Println(Greet(p))


}
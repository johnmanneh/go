package main

import (
	"fmt"
	"reflect"
)

//TAGS
//here elkuch is a struck.. but when i create the typeof.. it asked for an interface "Elkuch{}"
type Elkuch struct {
	Name string `name :"sarah"`
}

//
type labeh struct {
	Id     int
	Name   string
	status []string
}

//
func Stuctss() {
	fam := struct{ Name string }{Name: "john"}
	fmt.Println(fam)
}
func main() {

	t := reflect.TypeOf(Elkuch{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

	Stuctss()
	manneh := labeh{
		Id:   1,
		Name: "John",
		status: []string{
			"werns",
			"Hauser",
		},
	}
	fmt.Println(manneh.status[1])

	a := []int{1, 2, 3, 4, 4}
	b := []int{32, 23}
	c := append(a, b...)
	fmt.Println(c)
	x := c[:len(c)-1] //last  element
	i := c[1:]        // first element
	fmt.Println(x)
	fmt.Println(i)
	fmt.Println("labeh")

	state()
	fmt.Println("Stretch")
	makes()

}

func state() {

	statePop := map[string]int{"Salone": 1, "lumley": 2, "aberdeen": 3}
	m := map[[2]int]string{}
	fmt.Println(statePop["Salone"])
	fmt.Println(m)
	pop, ok := statePop["salone"]
	fmt.Println(pop, ok)
}

func makes() {
	labeh := make(map[int]int)
	labeh = map[int]int{3: 2}
	fmt.Println(labeh)
}

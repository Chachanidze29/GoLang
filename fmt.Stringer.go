package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

//Person is also of type Stringer because of overloaded method String()
//fmt and also log looks for stringer type and calls it if it exsits

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type Developer struct {
	Person
	progLang string
}

func (d Developer) String() string {
	return d.Person.String() + fmt.Sprintf("Programming language %v", d.progLang)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	d1 := Developer{Person{"Avtandil chachanidze", 18}, "Go"}
	d2 := Developer{Person{"Vakho misuradze", 28}, "JS"}
	fmt.Println(d1, d2)
	arr := []fmt.Stringer{&a, &z, &d1, &d2}
	for _, v := range arr {
		fmt.Println(v) // fmt calls overloaded String() method
	}
}

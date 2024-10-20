package main

import "fmt"

type Employee struct {
	Name  string
	Email string
}

type data int

func (v1 data) demo(v2 data) {
	fmt.Println(v1, v2)
}

func (d *data) changeValue(v int) {
	*d = data(v)

}

func (emp *Employee) empDemo(em string) {
	// fmt.Println(em.Name, em.Email)
	(*emp).Email = em

}

func main() {

	a := data(20)
	b := data(10)
	a.demo(200)
	b.demo(100)

	c := Employee{Name: "wap", Email: "Wap@gmail.com"}
	c.empDemo("ins@gmail.com")
	fmt.Println(c.Email, c.Name)

	d := data(100)
	d.changeValue(108)

	fmt.Println(d)

	// fmt.Println(a, b)

}

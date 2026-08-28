package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	tamanhoLado float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.tamanhoLado * q.tamanhoLado
}

func (c circulo) area() float64 {
	return 2 * math.Pi * c.raio
}

type figura interface {
	area() float64
}

func info(f figura) float64 {
	return f.area()
}

func main() {
	q1 := quadrado{4.7}
	c1 := circulo{5.6}
	q2 := quadrado{tamanhoLado: 4.7}
	c2 := circulo{raio: 5.6}

	fmt.Println(info(q1))
	fmt.Println(info(c1))
	fmt.Println(info(q2))
	fmt.Println(info(c2))
}

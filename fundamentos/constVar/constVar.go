package main

import (
	"fmt"
	//Podemos criar referencia do import
	m "math"
	//para manter um import sem usar ele na aplicação
	_ "os"
)

func main() {
	/*Criar variavel constante*/
	const PI float64 = 3.1415
	//variavel auto tipada, tipo float inferido pelo compilador
	var radio = 3.2

	//Forma reduzida para criar variavel
	area := PI * m.Pow(radio, 2)
	fmt.Println("Area da circunferencia é", area)

	//declarar mais de uma variavel ou constantes
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	//criando 2 variavel e passando valores em um unico comando
	var e, f bool = true, false
	fmt.Println(e, f)

	//criando 2 variavel e passando valores em um unico comando de forma reduzida
	g, h, i := 1, false, "epa!"
	fmt.Println(g, h, i)
}

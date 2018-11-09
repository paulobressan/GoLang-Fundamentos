//Pacote da aplicação, programas executaveis iniciam pelo pacote main
package main

/*
Os código em Go são organizados em pacotes
e para usá-los é necessario declarar um ou vários imports
*/
import (
	"fmt"
)

/*
A função main é a porta de entrada do programa go
Um programa só pode ter uma unica função main
*/
func main() {
	fmt.Println("Primeiro")
	fmt.Println("Programa!")
}

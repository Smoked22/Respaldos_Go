package main

import (
	"fmt"
)

func main() {
	nombre := "Jona"
	edad := 24
	direccion := "Juan Davalos #150"
	fmt.Printf("Hola, mi nombre es %s y tengo %d años y vivo en %s\n", nombre, edad, direccion)

	altura := 1.65
	casado := false

	fmt.Printf("Nombre: %s\n", nombre)
	fmt.Printf("Edad: %d\n", edad)
	fmt.Printf("Altura: %f\n", altura)
	fmt.Printf("¿Casado?: %t\n", casado)

	resta := 20

	resta = edad - resta

	fmt.Printf("tengo menos años %d\n", resta)
}

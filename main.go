package main

import (
	"fmt"
	"usuarios-roles/models"
)

func main() {
	// Crear un usuario con rol de administrador
	user1 := models.Usuario{
		Nombre: "Lautaro",
		Email:  "lautaro@example.com",
		Rol:    models.Admin{},
	}

	// Crear un usuario con rol de cliente
	user2 := models.Usuario{
		Nombre: "Yair",
		Email:  "yair@example.com",
		Rol:    models.Cliente{},
	}

	// Crear un usuario con rol de empleado
	user3 := models.Usuario{
		Nombre: "Ulises",
		Email:  "ulises@example.com",
		Rol:    models.Empleado{},
	}

	user1.MostrarInfo() // Mostrar información del usuario 1
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = = = = = =")
	user2.MostrarInfo() // Mostrar información del usuario 2
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = = = = = =")
	user3.MostrarInfo() // Mostrar información del usuario 3

	fmt.Println("\n= = = = = = = = = = = = = = = = = = = = = = = = = =\n")
	user1.CambiarRol(models.Empleado{}) // Cambiar el rol del usuario 1 a empleado
	user1.MostrarInfo()                 // Mostrar información del usuario 1
}

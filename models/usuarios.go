package models

import (
	"fmt"
)

// Usuarios es una estructura que representa a los usuarios
type Usuario struct {
	Nombre string
	Email  string
	Rol    Rol
}

// MostrarInfo es un método que muestra la información de un usuario
func (u Usuario) MostrarInfo() {
	fmt.Printf("Nombre: %s\nCorreo: %s\nRol: %s\n", u.Nombre, u.Email, u.Rol.ObtenerNombre())
	u.Rol.MostrarPermisos()
}

/*
	1️⃣ Un Cliente solo puede convertirse en Empleado, pero no en Administrador.
	2️⃣ Un Empleado puede convertirse en Administrador o volver a ser Cliente.
	3️⃣ Un Administrador puede cambiar a cualquier otro rol.
*/

// CambiarRol es un método que cambia el rol de un usuario
func (u *Usuario) CambiarRol(nuevoRol Rol) {
	rolActual := u.Rol.ObtenerNombre()         // Obtener el nombre del rol actual
	nuevoRolNombre := nuevoRol.ObtenerNombre() // Obtener el nombre del nuevo rol

	// Verificar si el usuario ya tiene el nuevo rol
	if rolActual == nuevoRolNombre {
		fmt.Printf("El usuario ya es %s\n", rolActual)
		return
	}

	// Verificar si el usuario es cliente y el nuevo rol es administrador
	if rolActual == "Cliente" && nuevoRolNombre == "Administrador" {
		fmt.Println("❌ error: un cliente no puede convertirse en administrador")
		return
	}

	// Verificar si el usuario es empleado y el nuevo rol es cliente o administrador
	if rolActual == "Empleado" && nuevoRolNombre == "Cliente" {
		fmt.Printf("❌ error: un empleado no puede convertirse en %s\n", nuevoRolNombre)
		return
	}

	fmt.Printf("El rol ha sido cambiado de: %s a %s\n", rolActual, nuevoRolNombre) // Mostrar mensaje de éxito
	u.Rol = nuevoRol                                                               // Cambiar el rol del usuario
}

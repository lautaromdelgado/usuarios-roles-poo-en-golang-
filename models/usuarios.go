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

func (u *Usuario) CambiarRol(nuevoRol Rol) {
	u.Rol = nuevoRol
	fmt.Printf("El rol ha sido cambiado a: %s\n", u.Rol.ObtenerNombre())
}

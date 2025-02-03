package models

import (
	"fmt"
)

// Rol es una interfaz que define los métodos que deben implementar los roles
type Rol interface {
	ObtenerNombre() string
	MostrarPermisos()
}

// Admin es una estructura que representa a un administrador
type Admin struct{}

// ObtenerNombre es un método que devuelve el nombre del rol
func (a Admin) ObtenerNombre() string {
	return "Administrador"
}

// MostrarPermisos es un método que muestra los permisos del rol
func (a Admin) MostrarPermisos() {
	fmt.Println("Permisos: Puede crear, editar y eliminar usuarios.")
}

// Cliente es una estructura que representa a un cliente
type Cliente struct{}

// ObtenerNombre es un método que devuelve el nombre del rol
func (c Cliente) ObtenerNombre() string {
	return "Cliente"
}

// MostrarPermisos es un método que muestra los permisos del rol
func (c Cliente) MostrarPermisos() {
	fmt.Println("Permisos: Puede ver productos y realizar compras.")
}

// Empleado es una estructura que representa a un empleado
type Empleado struct{}

// ObtenerNombre es un método que devuelve el nombre del rol
func (e Empleado) ObtenerNombre() string {
	return "Empleado"
}

// MostrarPermisos es un método que muestra los permisos del rol
func (e Empleado) MostrarPermisos() {
	fmt.Println("Permisos: Puede atender clientes y gestionar pedidos.")
}

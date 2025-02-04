package models

import (
	"encoding/json"
	"fmt"
)

// Usuarios es una estructura que representa a los usuarios
type Usuario struct {
	Nombre string `json:"Nombre"` // Nombre del usuario
	Email  string `json:"Email"`  // Correo electrónico del usuario
	Rol    Rol    `json:"Rol"`    // Rol del usuario
}

// MostrarInfo es un método que muestra la información de un usuario
func (u Usuario) MostrarInfo() {
	fmt.Printf("👤 Nombre: %s\n📧 Correo: %s\n", u.Nombre, u.Email)
	if u.Rol != nil {
		fmt.Printf("🔑 Rol: %s\n", u.Rol.ObtenerNombre())
	} else {
		fmt.Println("🔑 Rol: Sin rol asignado")
	}
}

// usuarioAux es una estructura auxiliar para el método CambiarRol
type usuarioAux struct {
	Nombre string          `json:"Nombre"` // Nombre del usuario
	Email  string          `json:"Email"`  // Apellido del usuario
	Rol    json.RawMessage `json:"Rol"`    // Rol del usuario
}

func (u *Usuario) UnmarshalJson(data []byte) error {
	var aux usuarioAux
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return fmt.Errorf("❌ error al deserializar el usuario: %v", err)
	}

	u.Nombre = aux.Nombre
	u.Email = aux.Email

	type rolTipo struct {
		Tipo string `json:"Tipo"`
	}

	var rt rolTipo
	if err := json.Unmarshal(aux.Rol, &rt); err != nil {
		return fmt.Errorf("❌ Error al deserializar el rol: %v", err)
	}

	switch rt.Tipo {
	case "Admin":
		u.Rol = Admin{}
	case "Cliente":
		u.Rol = Cliente{}
	case "Empleado":
		u.Rol = Empleado{}
	default:
		u.Rol = nil
	}
	return nil
}

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

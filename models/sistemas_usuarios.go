package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// SistemasUsuarios es una estructura que representa
// a un conjunto de usuarios
type SistemasUsuarios struct {
	Usuarios []Usuario
}

// AgregarUsuario es un método que agrega un usuario al sistema
func (s *SistemasUsuarios) AgregarUsuario(u Usuario) {
	s.Usuarios = append(s.Usuarios, u)             // Agregamos el usuario al sistema
	fmt.Printf("✅ Usuario agregado: %s", u.Nombre) // Mostramos un mensaje de éxito
}

// BuscarUsuarioNombre es un método que busca un usuario por su nombre
func (s *SistemasUsuarios) BuscarUsuarioNombre(n string) *Usuario {
	for i, u := range s.Usuarios { // Recorremos los usuarios
		if u.Nombre == n { // Si el nombre del usuario es igual al nombre buscado
			return &s.Usuarios[i] // Retornamos el usuario encontrado
		}
	}
	return nil // Si no se encontró el usuario, retornamos nil
}

// EliminarUsuario es un método que elimina un usuario del sistema
func (s *SistemasUsuarios) EliminarUsuario(n string) {
	for i, u := range s.Usuarios {
		if u.Nombre == n {
			s.Usuarios = append(s.Usuarios[:i], s.Usuarios[i+1:]...) // Eliminamos el usuario del slice
			fmt.Printf("🗑️ Usuario eliminado: %s", u.Nombre)
			return
		}
	}
	fmt.Println("❌ Usuario no encontrado")
}

// MostrarUsuarios es un método que muestra la información de los usuarios
func (s *SistemasUsuarios) MostrarUsuarios() {
	if len(s.Usuarios) > 0 { // Si hay usuarios registrados
		fmt.Println("📭 Usuarios registrados:")
		for _, u := range s.Usuarios { // Recorremos los usuarios
			u.MostrarInfo() // Mostramos la información del usuario
			fmt.Println("-----------------------")
		}
		return
	}
	fmt.Println("📭 No hay usuarios registrados.") // Si no hay usuarios registrados
}

// GuardarEnArchivos es un método que guarda los usuarios en un archivo
func (s *SistemasUsuarios) GuardarEnArchivos(nombreArchivo string) {
	// Convertimos los usuarios en JSON
	data, err := json.MarshalIndent(s.Usuarios, "", " ")
	if err != nil { // Si hay un error al convertir los usuarios
		fmt.Printf("❌ Error al convertir los usuarios en JSON: %s\n", err)
		return
	}
	// Guardamos los usuarios en un archivo
	err = os.WriteFile(nombreArchivo, data, 0644)
	if err != nil { // Si hay un error al guardar los usuarios
		fmt.Printf("❌ Error al guardar los usuarios en el archivo: %s\n", err)
		return
	}
	fmt.Printf("✅ Usuarios guardados en el archivo %s\n", nombreArchivo)
}

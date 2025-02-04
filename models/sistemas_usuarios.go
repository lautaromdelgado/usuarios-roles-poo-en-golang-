package models

import (
	"encoding/json"
	"fmt"
	"os"
)

// SistemaUsuarios es una estructura que maneja la lista de usuarios
type SistemaUsuarios struct {
	Usuarios []Usuario
}

// AgregarUsuario agrega un usuario a la lista
func (s *SistemaUsuarios) AgregarUsuario(usuario Usuario) {
	s.Usuarios = append(s.Usuarios, usuario)
	fmt.Println("✅ Usuario agregado:", usuario.Nombre)
}

// BuscarUsuarioPorNombre busca un usuario en la lista
func (s *SistemaUsuarios) BuscarUsuarioPorNombre(nombre string) *Usuario {
	for i := range s.Usuarios {
		if s.Usuarios[i].Nombre == nombre {
			return &s.Usuarios[i] // Retorna un puntero al usuario
		}
	}
	return nil
}

// EliminarUsuario elimina un usuario de la lista por su nombre
func (s *SistemaUsuarios) EliminarUsuario(nombre string) {
	for i, usuario := range s.Usuarios {
		if usuario.Nombre == nombre {
			s.Usuarios = append(s.Usuarios[:i], s.Usuarios[i+1:]...) // Remueve el usuario
			fmt.Println("🗑️ Usuario eliminado:", nombre)
			return
		}
	}
	fmt.Println("❌ Usuario no encontrado:", nombre)
}

// MostrarUsuarios imprime todos los usuarios de la lista
func (s *SistemaUsuarios) MostrarUsuarios() {
	if len(s.Usuarios) == 0 {
		fmt.Println("📭 No hay usuarios en el sistema.")
		return
	}
	fmt.Println("📋 Lista de Usuarios:")
	for _, usuario := range s.Usuarios {
		usuario.MostrarInfo()
		fmt.Println("-----------------------")
	}
}

// GuardarEnArchivo guarda los usuarios en un archivo JSON
func (s *SistemaUsuarios) GuardarEnArchivo(nombreArchivo string) {
	data, err := json.MarshalIndent(s.Usuarios, "", "  ")
	if err != nil {
		fmt.Println("❌ Error al serializar usuarios:", err)
		return
	}
	err = os.WriteFile(nombreArchivo, data, 0644)
	if err != nil {
		fmt.Println("❌ Error al guardar archivo:", err)
		return
	}
	fmt.Println("✅ Usuarios guardados en", nombreArchivo)
}

// CargarDesdeArchivo carga los usuarios desde un archivo JSON
func (s *SistemaUsuarios) CargarDesdeArchivo(nombreArchivo string) {
	data, err := os.ReadFile(nombreArchivo)
	if err != nil {
		fmt.Println("⚠️ No se pudo leer el archivo, iniciando con lista vacía.")
		return
	}
	err = json.Unmarshal(data, &s.Usuarios)
	if err != nil {
		fmt.Println("❌ Error al deserializar usuarios:", err)
	} else {
		fmt.Println("✅ Usuarios cargados desde", nombreArchivo)
	}
}

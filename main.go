package main

// TESTEOS CON PRUEBAS DE EJEMPLO
// SETEANDO ROLES Y CAMBIANDO LOS ROLES DE LOS USUARIOS

/* PRUEBA 1° */
// func main() {
// 	// Creamos un usuario con el rol de Administrador
// 	usuario1 := models.Usuario{
// 		Nombre: "Lautaro",
// 		Email:  "lautaro@example.com",
// 		Rol:    models.Admin{}, // Asignamos el rol Admin
// 	}

// 	// Creamos otro usuario con el rol de Cliente
// 	usuario2 := models.Usuario{
// 		Nombre: "Carlos",
// 		Email:  "carlos@example.com",
// 		Rol:    models.Cliente{}, // Asignamos el rol Cliente
// 	}

// 	// Mostramos la información de los usuarios
// 	usuario1.MostrarInfo()
// 	fmt.Println("-----------------------")
// 	usuario2.MostrarInfo()
// }

/*

SALIDA ESPERADA:
Nombre: Lautaro
Correo: lautaro@example.com
Rol: Administrador
Permisos: Puede crear, editar y eliminar usuarios.
-----------------------
Nombre: Carlos
Correo: carlos@example.com
Rol: Cliente
Permisos: Puede ver su perfil y realizar compras.

*/

/* - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - */

/*PRUEBA 2°*/
// func main() {
// 	// Creamos un usuario con el rol de Cliente
// 	usuario := models.Usuario{
// 		Nombre: "Lautaro",
// 		Email:  "lautaro@example.com",
// 		Rol:    models.Cliente{}, // Inicialmente es Cliente
// 	}

// 	// Mostramos la información inicial
// 	usuario.MostrarInfo()
// 	fmt.Println("-----------------------")

// 	// Cambiamos el rol del usuario a Administrador
// 	usuario.CambiarRol(models.Admin{})

// 	// Mostramos la información después del cambio
// 	fmt.Println("-----------------------")
// 	usuario.MostrarInfo()
// }

/*

SALIDA ESPERADA:
Nombre: Lautaro
Correo: lautaro@example.com
Rol: Cliente
Permisos: Puede ver productos y realizar compras.
-----------------------
❌ error: un cliente no puede convertirse en administrador
-----------------------
Nombre: Lautaro
Correo: lautaro@example.com
Rol: Cliente
Permisos: Puede ver productos y realizar compras.

*/

/* PRUEBA 3° */
// func main() {
// 	usuario := models.Usuario{
// 		Nombre: "Lautaro",
// 		Email:  "lautaro@example.com",
// 		Rol:    models.Cliente{}, // Inicialmente es Cliente
// 	}

// 	usuario.MostrarInfo()
// 	fmt.Println("-----------------------")

// 	// Intentamos cambiar de Cliente a Administrador (No permitido)
// 	usuario.CambiarRol(models.Admin{})
// 	fmt.Println("-----------------------")

// 	// Cambiamos de Cliente a Empleado (Permitido)
// 	usuario.CambiarRol(models.Empleado{})
// 	usuario.MostrarInfo()
// 	fmt.Println("-----------------------")

// 	// Intentamos cambiar de Empleado a Cliente (No permitido)
// 	usuario.CambiarRol(models.Cliente{})
// 	fmt.Println("-----------------------")

// 	// Cambiamos de Empleado a Administrador (Permitido)
// 	usuario.CambiarRol(models.Admin{})
// 	usuario.MostrarInfo()
// }

/*
SALIDA ESPERADA:
Nombre: Lautaro
Correo: lautaro@example.com
Rol: Cliente
Permisos: Puede ver su perfil y realizar compras.
-----------------------
❌ Error: Un Cliente no puede convertirse en Administrador.
-----------------------
✅ El rol ha sido cambiado a: Empleado
Nombre: Lautaro
Correo: lautaro@example.com
Rol: Empleado
Permisos: Puede atender clientes y gestionar pedidos.
-----------------------
❌ Error: Un Empleado no puede convertirse en Cliente.
-----------------------
✅ El rol ha sido cambiado a: Administrador
Nombre: Lautaro
Correo: lautaro@example.com
Rol: Administrador
Permisos: Puede crear, editar y eliminar usuarios.

*/

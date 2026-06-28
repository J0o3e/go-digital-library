# Servicios Web JSON - Go Digital Library

Este documento describe los servicios web del proyecto **Go Digital Library**.

La API fue desarrollada en Go usando los paquetes estándar:

- `net/http`
- `encoding/json`

La serialización de datos se realiza mediante JSON.

## URL base

```text
http://localhost:8080

1. Listar libros
Método
GET /api/books
Descripción

Devuelve todos los libros registrados en la base de datos.

Respuesta ejemplo
[
  {
    "id": 1,
    "titulo": "La casa de los espiritus",
    "autor": "Isabel Allende",
    "categoria": "Literatura latinoamericana",
    "anio": 1982,
    "disponible": true
  }
]

2. Registrar libro
Método
POST /api/books
Descripción

Registra un libro nuevo en la base de datos.

Cuerpo JSON
{
  "titulo": "Clean Code",
  "autor": "Robert C. Martin",
  "categoria": "Programacion",
  "anio": 2008
}
Respuesta ejemplo
{
  "id": 2,
  "titulo": "Clean Code",
  "autor": "Robert C. Martin",
  "categoria": "Programacion",
  "anio": 2008,
  "disponible": true
}

3. Buscar libros
Método
GET /api/books/search?q=clean
Descripción

Busca libros por título, autor o categoría.

Respuesta ejemplo
[
  {
    "id": 2,
    "titulo": "Clean Code",
    "autor": "Robert C. Martin",
    "categoria": "Programacion",
    "anio": 2008,
    "disponible": true
  }
]

4. Listar usuarios
Método
GET /api/users
Descripción

Devuelve todos los usuarios registrados.

Respuesta ejemplo
[
  {
    "id": 1,
    "nombre": "Joseph Solano",
    "correo": "joseph@example.com",
    "activo": true
  }
]

5. Registrar usuario
Método
POST /api/users
Descripción

Registra un usuario nuevo en la base de datos.

Cuerpo JSON
{
  "nombre": "Joseph Solano",
  "correo": "joseph@example.com"
}
Respuesta ejemplo
{
  "id": 1,
  "nombre": "Joseph Solano",
  "correo": "joseph@example.com",
  "activo": true
}

6. Buscar usuarios
Método
GET /api/users/search?q=joseph
Descripción

Busca usuarios por nombre o correo.

Respuesta ejemplo
[
  {
    "id": 1,
    "nombre": "Joseph Solano",
    "correo": "joseph@example.com",
    "activo": true
  }
]

7. Registrar préstamo
Método
POST /api/loans
Descripción

Registra un préstamo entre un libro disponible y un usuario activo.

Cuando se registra el préstamo, el libro pasa a estado no disponible.

Cuerpo JSON
{
  "libro_id": 1,
  "usuario_id": 1
}
Respuesta ejemplo
{
  "id": 1,
  "libro_id": 1,
  "libro_titulo": "La casa de los espiritus",
  "usuario_id": 1,
  "usuario_nombre": "Joseph Solano",
  "fecha_prestamo": "2026-06-27 15:30",
  "fecha_devolucion": "Pendiente",
  "activo": true
}

8. Listar préstamos activos
Método
GET /api/loans/active
Descripción

Devuelve los préstamos que todavía no han sido devueltos.

Respuesta ejemplo
[
  {
    "id": 1,
    "libro_id": 1,
    "libro_titulo": "La casa de los espiritus",
    "usuario_id": 1,
    "usuario_nombre": "Joseph Solano",
    "fecha_prestamo": "2026-06-27 15:30",
    "fecha_devolucion": "Pendiente",
    "activo": true
  }
]

9. Historial de préstamos
Método
GET /api/loans/history
Descripción

Devuelve todos los préstamos registrados, tanto activos como devueltos.

Respuesta ejemplo
[
  {
    "id": 1,
    "libro_id": 1,
    "libro_titulo": "La casa de los espiritus",
    "usuario_id": 1,
    "usuario_nombre": "Joseph Solano",
    "fecha_prestamo": "2026-06-27 15:30",
    "fecha_devolucion": "2026-06-27 15:45",
    "activo": false
  }
]

10. Devolver préstamo
Método
POST /api/loans/return
Descripción

Marca un préstamo como devuelto y cambia el libro nuevamente a disponible.

Cuerpo JSON
{
  "prestamo_id": 1
}
Respuesta ejemplo
{
  "message": "prestamo devuelto correctamente"
}

Ejemplos de prueba con PowerShell

Crear libro

Invoke-RestMethod -Uri "http://localhost:8080/api/books" -Method Post -ContentType "application/json" -Body '{"titulo":"Clean Code","autor":"Robert C. Martin","categoria":"Programacion","anio":2008}'

Crear usuario

Invoke-RestMethod -Uri "http://localhost:8080/api/users" -Method Post -ContentType "application/json" -Body '{"nombre":"Joseph Solano","correo":"joseph@example.com"}'

Crear préstamo

Invoke-RestMethod -Uri "http://localhost:8080/api/loans" -Method Post -ContentType "application/json" -Body '{"libro_id":1,"usuario_id":1}'

Devolver préstamo

Invoke-RestMethod -Uri "http://localhost:8080/api/loans/return" -Method Post -ContentType "application/json" -Body '{"prestamo_id":1}'


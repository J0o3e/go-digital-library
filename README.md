# Go Digital Library

Go Digital Library es un sistema de biblioteca digital desarrollado en el lenguaje Go.

El proyecto permite registrar libros, registrar usuarios, realizar préstamos, devolver libros, buscar información y consultar el historial de préstamos. Los datos se almacenan de forma persistente en una base de datos PostgreSQL.

## Objetivo del proyecto

El objetivo principal del proyecto es aplicar conceptos de programación orientada a objetos y estructuras de datos usando Go, organizando el sistema en paquetes y separando responsabilidades.

El proyecto aplica conceptos como:

- Structs
- Métodos
- Constructores
- Encapsulación
- Punteros
- Slices
- Maps
- Funciones
- Manejo de errores
- Condicionales `if`
- Ciclos `for`
- Estructura `switch`
- Conexión a base de datos PostgreSQL

## Tecnologías utilizadas

- Go
- PostgreSQL
- pgx
- godotenv
- Git
- GitHub
- Visual Studio Code

## Funcionalidades principales

El sistema permite:

- Registrar libros.
- Listar libros.
- Buscar libros por título, autor o categoría.
- Registrar usuarios.
- Listar usuarios.
- Buscar usuarios por nombre o correo.
- Registrar préstamos.
- Listar préstamos activos.
- Devolver préstamos.
- Consultar historial de préstamos.

## Arquitectura del proyecto

El proyecto está organizado en paquetes para separar responsabilidades:

```text
go-digital-library/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── books/
│   ├── config/
│   ├── console/
│   ├── database/
│   ├── loans/
│   ├── models/
│   ├── repositories/
│   └── users/
├── docs/
├── scripts/
├── .env.example
├── go.mod
└── README.md


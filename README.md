\# Go Digital Library



Sistema de Biblioteca Digital desarrollado en Go con conexión a PostgreSQL.



\## Descripción



Este proyecto tiene como objetivo desarrollar una biblioteca digital usando el lenguaje de programación Go. El sistema permitirá gestionar libros, usuarios y préstamos, almacenando la información en una base de datos externa PostgreSQL.



\## Objetivo general



Desarrollar un sistema de Biblioteca Digital en Go que permita registrar, consultar, buscar, prestar y devolver libros, usando PostgreSQL como base de datos externa.



\## Módulos del sistema



\- Gestión de libros

\- Gestión de usuarios

\- Gestión de préstamos

\- Gestión de validaciones

\- Conexión a PostgreSQL

\- Acceso a datos

\- Interfaz por consola



\## Temas aplicados de la Unidad 1



\- Condicionales `if` e `if else`

\- Estructura `switch`

\- Bucle `for`

\- Funciones con y sin parámetros

\- Funciones con uno o varios retornos

\- Uso de punteros



\## Tecnologías utilizadas



\- Go

\- PostgreSQL

\- Git

\- GitHub



\## Paquetes principales



\### Paquetes estándar de Go



\- `fmt`

\- `bufio`

\- `os`

\- `strings`

\- `strconv`

\- `time`

\- `context`

\- `database/sql`

\- `errors`



\### Paquete externo



\- `github.com/jackc/pgx/v5`



\## Estructura del proyecto



```text

go-digital-library/

│

├── cmd/

│   └── app/

│       └── main.go

│

├── internal/

│   ├── models/

│   ├── console/

│   ├── books/

│   ├── users/

│   ├── loans/

│   ├── database/

│   ├── repositories/

│   └── validations/

│

├── docs/

├── scripts/

├── .gitignore

├── .env.example

├── README.md

└── go.mod


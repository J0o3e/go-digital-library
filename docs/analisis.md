\# Análisis del Sistema



\## Nombre del proyecto



\*\*Go Digital Library\*\*



\## Descripción del sistema



Go Digital Library es un sistema de Biblioteca Digital desarrollado en el lenguaje de programación Go. Su propósito es gestionar libros, usuarios y préstamos dentro de una biblioteca, utilizando PostgreSQL como base de datos externa para almacenar la información de forma persistente.



El sistema será desarrollado inicialmente como una aplicación de consola, permitiendo al usuario interactuar mediante un menú principal.



\## Objetivo general



Desarrollar un sistema de Biblioteca Digital en Go que permita registrar, consultar, buscar, prestar y devolver libros, usando PostgreSQL como base de datos externa.



\## Objetivos específicos



\- Registrar libros dentro del catálogo de la biblioteca.

\- Consultar la lista de libros registrados.

\- Buscar libros por título, autor o categoría.

\- Registrar usuarios que podrán solicitar préstamos.

\- Gestionar préstamos de libros disponibles.

\- Registrar devoluciones de libros prestados.

\- Actualizar el estado de disponibilidad de cada libro.

\- Conectarse a una base de datos externa PostgreSQL.

\- Organizar el sistema mediante paquetes propios en Go.

\- Aplicar los temas vistos en la Unidad 1.

\- Aplicar principios de programación funcional.



\## Módulos del sistema



\### 1. Gestión de libros



Este módulo permitirá registrar, consultar, buscar y actualizar el estado de los libros.



Funciones principales:



\- Registrar libro.

\- Listar libros.

\- Buscar libro.

\- Verificar disponibilidad.

\- Actualizar estado del libro.



\### 2. Gestión de usuarios



Este módulo permitirá administrar los usuarios registrados en la biblioteca.



Funciones principales:



\- Registrar usuario.

\- Listar usuarios.

\- Buscar usuario.

\- Validar existencia del usuario.



\### 3. Gestión de préstamos



Este módulo permitirá controlar los préstamos y devoluciones de libros.



Funciones principales:



\- Registrar préstamo.

\- Registrar devolución.

\- Consultar préstamos activos.

\- Consultar historial de préstamos.



\### 4. Gestión de validaciones



Este módulo se encargará de verificar que los datos ingresados sean correctos.



Validaciones principales:



\- Validar campos vacíos.

\- Validar año de publicación.

\- Validar identificadores.

\- Validar disponibilidad del libro.

\- Validar existencia de usuario.

\- Validar existencia de préstamo.



\### 5. Conexión a PostgreSQL



Este módulo permitirá conectar el sistema con una base de datos externa PostgreSQL.



Funciones principales:



\- Abrir conexión con PostgreSQL.

\- Verificar conexión.

\- Cerrar conexión.

\- Proporcionar la conexión a los módulos que la necesiten.



\### 6. Acceso a datos



Este módulo se encargará de realizar operaciones sobre la base de datos.



Funciones principales:



\- Insertar libros.

\- Consultar libros.

\- Insertar usuarios.

\- Consultar usuarios.

\- Registrar préstamos.

\- Actualizar estados.

\- Consultar préstamos.



\### 7. Interfaz por consola



Este módulo permitirá la interacción entre el usuario y el sistema.



Funciones principales:



\- Mostrar menú principal.

\- Leer opciones del usuario.

\- Solicitar datos.

\- Mostrar resultados.

\- Mostrar mensajes de error o confirmación.



\## Estructura de gestión del sistema



El sistema estará organizado en módulos funcionales. Cada módulo tendrá una responsabilidad específica para facilitar el mantenimiento y crecimiento del proyecto.



```text

Biblioteca Digital en Go

│

├── Interfaz por consola

│   ├── Mostrar menú

│   ├── Leer datos

│   └── Mostrar resultados

│

├── Gestión de libros

│   ├── Registrar libro

│   ├── Listar libros

│   ├── Buscar libro

│   └── Actualizar disponibilidad

│

├── Gestión de usuarios

│   ├── Registrar usuario

│   ├── Listar usuarios

│   └── Buscar usuario

│

├── Gestión de préstamos

│   ├── Registrar préstamo

│   ├── Registrar devolución

│   ├── Consultar préstamos activos

│   └── Consultar historial

│

├── Validaciones

│   ├── Validar libro

│   ├── Validar usuario

│   └── Validar préstamo

│

├── Acceso a datos

│   ├── Consultas de libros

│   ├── Consultas de usuarios

│   └── Consultas de préstamos

│

└── Conexión PostgreSQL

&#x20;   ├── Abrir conexión

&#x20;   ├── Verificar conexión

&#x20;   └── Cerrar conexión


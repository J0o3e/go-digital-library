# Ejecución del proyecto Go Digital Library

Este documento explica los pasos necesarios para instalar, configurar y ejecutar el proyecto **Go Digital Library**.

## 1. Requisitos

Para ejecutar el proyecto se necesita tener instalado:

* Go
* PostgreSQL
* pgAdmin 4, opcional pero recomendado
* Git
* Visual Studio Code, opcional pero recomendado

## 2. Clonar el repositorio

Para descargar el proyecto desde GitHub:

```bash
git clone URL_DEL_REPOSITORIO
```

Luego entrar a la carpeta del proyecto:

```bash
cd go-digital-library
```

## 3. Configurar PostgreSQL

Crear una base de datos en PostgreSQL con el siguiente nombre:

```text
go_digital_library
```

La base de datos puede crearse desde pgAdmin 4:

1. Abrir pgAdmin 4.
2. Ingresar al servidor PostgreSQL.
3. Clic derecho en `Databases`.
4. Seleccionar `Create > Database`.
5. Escribir el nombre `go_digital_library`.
6. Guardar.

## 4. Crear las tablas

Abrir el archivo:

```text
scripts/schema.sql
```

Copiar su contenido y ejecutarlo en el Query Tool de pgAdmin sobre la base de datos:

```text
go_digital_library
```

Este script crea las tablas principales del sistema:

* `libros`
* `usuarios`
* `prestamos`

## 5. Configurar variables de entorno

Crear un archivo llamado `.env` en la raíz del proyecto.

El archivo debe tener esta estructura:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=TU_PASSWORD_DE_POSTGRES
DB_NAME=go_digital_library
DB_SSLMODE=disable
```

Ejemplo:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=12345
DB_NAME=go_digital_library
DB_SSLMODE=disable
```

Importante: el archivo `.env` no debe subirse a GitHub porque contiene datos sensibles, como la contraseña de PostgreSQL.

Por esa razón, el proyecto incluye un archivo `.env.example` como referencia.

## 6. Instalar dependencias

Desde la raíz del proyecto, ejecutar:

```bash
go mod tidy
```

Este comando descarga y organiza las dependencias necesarias del proyecto.

Entre las dependencias usadas están:

* `github.com/jackc/pgx/v5`, para conectar Go con PostgreSQL.
* `github.com/joho/godotenv`, para leer variables desde el archivo `.env`.

## 7. Ejecutar el proyecto

Desde la raíz del proyecto, ejecutar:

```bash
go run ./cmd/app
```

Si la conexión con PostgreSQL funciona correctamente, debe aparecer un mensaje parecido a:

```text
Conexion a PostgreSQL exitosa.
```

Luego se mostrará el menú principal del sistema:

```text
===== GO DIGITAL LIBRARY =====
1. Registrar libro
2. Listar libros
3. Registrar usuario
4. Listar usuarios
5. Registrar prestamo
6. Listar prestamos activos
7. Devolver prestamo
8. Buscar libros
9. Buscar usuarios
10. Historial de prestamos
11. Salir
```

## 8. Funcionalidades principales

El sistema permite:

* Registrar libros.
* Listar libros.
* Buscar libros por título, autor o categoría.
* Registrar usuarios.
* Listar usuarios.
* Buscar usuarios por nombre o correo.
* Registrar préstamos.
* Listar préstamos activos.
* Devolver préstamos.
* Ver el historial de préstamos.

## 9. Verificar que el proyecto compila

Para verificar que todos los paquetes compilan correctamente:

```bash
go test ./...
```

Aunque el proyecto todavía no tenga pruebas automatizadas, este comando ayuda a confirmar que no existen errores de compilación.

## 10. Comandos básicos de Git usados

Ver el estado del repositorio:

```bash
git status
```

Agregar cambios:

```bash
git add .
```

Crear un commit:

```bash
git commit -m "Mensaje del commit"
```

Subir cambios a GitHub:

```bash
git push
```

Ver ramas:

```bash
git branch
```

Cambiar de rama:

```bash
git checkout nombre-rama
```

## 11. Notas importantes

El proyecto usa PostgreSQL para almacenar los datos de forma persistente.

Las tablas principales son:

* `libros`
* `usuarios`
* `prestamos`

La conexión a PostgreSQL se configura mediante variables de entorno en el archivo `.env`.

El archivo `.env` debe existir localmente, pero no debe subirse al repositorio.

El archivo `scripts/schema.sql` debe ejecutarse antes de iniciar el sistema, para que existan las tablas necesarias.

## 12. Comando principal de ejecución

El comando principal para ejecutar el sistema es:

```bash
go run ./cmd/app
```

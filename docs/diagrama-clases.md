\# Diagrama de Clases del Sistema



\## Descripción



El sistema \*\*Go Digital Library\*\* será desarrollado en el lenguaje Go. Aunque Go no utiliza clases de forma tradicional como otros lenguajes orientados a objetos, el sistema puede representarse mediante un diagrama de clases conceptual.



En este diagrama, las clases representan las estructuras, entidades y módulos principales que serán incorporados en el sistema.



\## Diagrama de clases conceptual



```mermaid

classDiagram



class Libro {

&#x20;   +int ID

&#x20;   +string Titulo

&#x20;   +string Autor

&#x20;   +string Categoria

&#x20;   +int Anio

&#x20;   +string Estado

}



class Usuario {

&#x20;   +int ID

&#x20;   +string Nombre

&#x20;   +string Correo

&#x20;   +string Estado

}



class Prestamo {

&#x20;   +int ID

&#x20;   +int LibroID

&#x20;   +int UsuarioID

&#x20;   +date FechaPrestamo

&#x20;   +date FechaDevolucion

&#x20;   +string Estado

}



class SistemaBiblioteca {

&#x20;   +RegistrarLibro()

&#x20;   +ListarLibros()

&#x20;   +BuscarLibro()

&#x20;   +RegistrarUsuario()

&#x20;   +ListarUsuarios()

&#x20;   +PrestarLibro()

&#x20;   +DevolverLibro()

&#x20;   +ConsultarPrestamos()

}



class PostgreSQL {

&#x20;   +GuardarDatos()

&#x20;   +ConsultarDatos()

&#x20;   +ActualizarDatos()

}



Usuario "1" --> "0..\*" Prestamo : realiza

Libro "1" --> "0..\*" Prestamo : es prestado en



SistemaBiblioteca --> Libro : gestiona

SistemaBiblioteca --> Usuario : gestiona

SistemaBiblioteca --> Prestamo : gestiona

SistemaBiblioteca --> PostgreSQL : almacena datos


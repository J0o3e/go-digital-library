CREATE TABLE IF NOT EXISTS libros (
    id SERIAL PRIMARY KEY,
    titulo VARCHAR(150) NOT NULL,
    autor VARCHAR(100) NOT NULL,
    categoria VARCHAR(80) NOT NULL,
    anio INT NOT NULL,
    estado VARCHAR(20) NOT NULL DEFAULT 'disponible'
);

CREATE TABLE IF NOT EXISTS usuarios (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(120) NOT NULL,
    correo VARCHAR(120) NOT NULL UNIQUE,
    estado VARCHAR(20) NOT NULL DEFAULT 'activo'
);

CREATE TABLE IF NOT EXISTS prestamos (
    id SERIAL PRIMARY KEY,
    libro_id INT NOT NULL,
    usuario_id INT NOT NULL,
    fecha_prestamo DATE NOT NULL,
    fecha_devolucion DATE,
    estado VARCHAR(20) NOT NULL DEFAULT 'activo',

    CONSTRAINT fk_libro
        FOREIGN KEY (libro_id)
        REFERENCES libros(id),

    CONSTRAINT fk_usuario
        FOREIGN KEY (usuario_id)
        REFERENCES usuarios(id)
);
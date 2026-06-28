const booksTableBody = document.getElementById("booksTableBody");
const bookForm = document.getElementById("bookForm");
const message = document.getElementById("message");
const reloadBooksButton = document.getElementById("reloadBooks");

async function loadBooks() {
    try {
        const response = await fetch("/api/books");
        const books = await response.json();

        booksTableBody.innerHTML = "";

        if (books.length === 0) {
            booksTableBody.innerHTML = `
                <tr>
                    <td colspan="6" class="empty">No hay libros registrados.</td>
                </tr>
            `;
            return;
        }

        books.forEach((book) => {
            const row = document.createElement("tr");

            const estado = book.disponible ? "Disponible" : "Prestado";

            row.innerHTML = `
                <td>${book.id}</td>
                <td>${book.titulo}</td>
                <td>${book.autor}</td>
                <td>${book.categoria}</td>
                <td>${book.anio}</td>
                <td>${estado}</td>
            `;

            booksTableBody.appendChild(row);
        });
    } catch (error) {
        message.textContent = "Error al cargar libros.";
    }
}

bookForm.addEventListener("submit", async function (event) {
    event.preventDefault();

    const book = {
        titulo: document.getElementById("titulo").value,
        autor: document.getElementById("autor").value,
        categoria: document.getElementById("categoria").value,
        anio: Number(document.getElementById("anio").value)
    };

    try {
        const response = await fetch("/api/books", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(book)
        });

        const data = await response.json();

        if (!response.ok) {
            message.textContent = data.error || "Error al registrar libro.";
            return;
        }

        message.textContent = "Libro registrado correctamente.";
        bookForm.reset();
        loadBooks();
    } catch (error) {
        message.textContent = "Error al conectar con la API.";
    }
});

reloadBooksButton.addEventListener("click", loadBooks);

loadBooks();
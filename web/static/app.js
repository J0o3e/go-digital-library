const booksTableBody = document.getElementById("booksTableBody");
const bookForm = document.getElementById("bookForm");
const message = document.getElementById("message");
const reloadBooksButton = document.getElementById("reloadBooks");

const usersTableBody = document.getElementById("usersTableBody");
const userForm = document.getElementById("userForm");
const userMessage = document.getElementById("userMessage");
const reloadUsersButton = document.getElementById("reloadUsers");

const loanForm = document.getElementById("loanForm");
const loanMessage = document.getElementById("loanMessage");
const activeLoansTableBody = document.getElementById("activeLoansTableBody");
const loanHistoryTableBody = document.getElementById("loanHistoryTableBody");
const reloadActiveLoansButton = document.getElementById("reloadActiveLoans");
const reloadLoanHistoryButton = document.getElementById("reloadLoanHistory");

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

async function loadUsers() {
    try {
        const response = await fetch("/api/users");
        const users = await response.json();

        usersTableBody.innerHTML = "";

        if (users.length === 0) {
            usersTableBody.innerHTML = `
                <tr>
                    <td colspan="4" class="empty">No hay usuarios registrados.</td>
                </tr>
            `;
            return;
        }

        users.forEach((user) => {
            const row = document.createElement("tr");

            const estado = user.activo ? "Activo" : "Inactivo";

            row.innerHTML = `
                <td>${user.id}</td>
                <td>${user.nombre}</td>
                <td>${user.correo}</td>
                <td>${estado}</td>
            `;

            usersTableBody.appendChild(row);
        });
    } catch (error) {
        userMessage.textContent = "Error al cargar usuarios.";
    }
}

async function loadActiveLoans() {
    try {
        const response = await fetch("/api/loans/active");
        const loans = await response.json();

        activeLoansTableBody.innerHTML = "";

        if (loans.length === 0) {
            activeLoansTableBody.innerHTML = `
                <tr>
                    <td colspan="7" class="empty">No hay préstamos activos.</td>
                </tr>
            `;
            return;
        }

        loans.forEach((loan) => {
            const row = document.createElement("tr");

            const estado = loan.activo ? "Activo" : "Devuelto";

            row.innerHTML = `
                <td>${loan.id}</td>
                <td>${loan.libro_titulo}</td>
                <td>${loan.usuario_nombre}</td>
                <td>${loan.fecha_prestamo}</td>
                <td>${loan.fecha_devolucion}</td>
                <td>${estado}</td>
                <td>
                    <button onclick="returnLoan(${loan.id})">Devolver</button>
                </td>
            `;

            activeLoansTableBody.appendChild(row);
        });
    } catch (error) {
        loanMessage.textContent = "Error al cargar préstamos activos.";
    }
}

async function loadLoanHistory() {
    try {
        const response = await fetch("/api/loans/history");
        const loans = await response.json();

        loanHistoryTableBody.innerHTML = "";

        if (loans.length === 0) {
            loanHistoryTableBody.innerHTML = `
                <tr>
                    <td colspan="6" class="empty">No hay préstamos registrados.</td>
                </tr>
            `;
            return;
        }

        loans.forEach((loan) => {
            const row = document.createElement("tr");

            const estado = loan.activo ? "Activo" : "Devuelto";

            row.innerHTML = `
                <td>${loan.id}</td>
                <td>${loan.libro_titulo}</td>
                <td>${loan.usuario_nombre}</td>
                <td>${loan.fecha_prestamo}</td>
                <td>${loan.fecha_devolucion}</td>
                <td>${estado}</td>
            `;

            loanHistoryTableBody.appendChild(row);
        });
    } catch (error) {
        loanMessage.textContent = "Error al cargar historial de préstamos.";
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

userForm.addEventListener("submit", async function (event) {
    event.preventDefault();

    const user = {
        nombre: document.getElementById("nombre").value,
        correo: document.getElementById("correo").value
    };

    try {
        const response = await fetch("/api/users", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(user)
        });

        const data = await response.json();

        if (!response.ok) {
            userMessage.textContent = data.error || "Error al registrar usuario.";
            return;
        }

        userMessage.textContent = "Usuario registrado correctamente.";
        userForm.reset();
        loadUsers();
    } catch (error) {
        userMessage.textContent = "Error al conectar con la API.";
    }
});

loanForm.addEventListener("submit", async function (event) {
    event.preventDefault();

    const loan = {
        libro_id: Number(document.getElementById("libroId").value),
        usuario_id: Number(document.getElementById("usuarioId").value)
    };

    try {
        const response = await fetch("/api/loans", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(loan)
        });

        const data = await response.json();

        if (!response.ok) {
            loanMessage.textContent = data.error || "Error al registrar préstamo.";
            return;
        }

        loanMessage.textContent = "Préstamo registrado correctamente.";
        loanForm.reset();

        loadBooks();
        loadActiveLoans();
        loadLoanHistory();
    } catch (error) {
        loanMessage.textContent = "Error al conectar con la API.";
    }
});

async function returnLoan(prestamoId) {
    try {
        const response = await fetch("/api/loans/return", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                prestamo_id: prestamoId
            })
        });

        const data = await response.json();

        if (!response.ok) {
            loanMessage.textContent = data.error || "Error al devolver préstamo.";
            return;
        }

        loanMessage.textContent = data.message || "Préstamo devuelto correctamente.";

        loadBooks();
        loadActiveLoans();
        loadLoanHistory();
    } catch (error) {
        loanMessage.textContent = "Error al conectar con la API.";
    }
}

reloadBooksButton.addEventListener("click", loadBooks);
reloadUsersButton.addEventListener("click", loadUsers);
reloadActiveLoansButton.addEventListener("click", loadActiveLoans);
reloadLoanHistoryButton.addEventListener("click", loadLoanHistory);

loadBooks();
loadUsers();
loadActiveLoans();
loadLoanHistory();



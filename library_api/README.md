# Library API

This project implements a RESTful API for managing a simple library system. It allows performing CRUD operations on books, with data stored in memory.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/opchaudhary/bookstore.git
    ```

2. Navigate to the project directory:

    ```bash
    cd library-api
    ```

3. Initialize the Go module:

    ```bash
    go mod init https://github.com/opchaudhary/bookstore.git
    ```

4. Run the application:

    ```bash
    go run main.go
    ```

## Usage

The API supports the following endpoints:

- `GET /books`: Retrieve all books.
- `GET /books/{id}`: Retrieve a specific book by its ID.
- `POST /books/add`: Add a new book.
- `PUT /books/update{id}`: Update an existing book by its ID.
- `DELETE /books/delete{id}`: Delete a book by its ID.

### Example

To add a book:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"id": "1", "title": "Sample Book", "author": "John Doe"}' http://localhost:8080/books/add

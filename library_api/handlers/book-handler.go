package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"omprakash/library_api/models"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Database connection string
const (
	host     = "localhost"
	port     = 5432
	user     = "omprakash"
	password = "omprakash"
	dbname   = "library"
)

// DB represents the database connection
var DB *sql.DB

func init() {
	var err error
	// Create a connection pool
	connStr := "postgres://" + user + ":" + password + "@localhost:" + strconv.Itoa(port) + "/" + dbname + "?sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// Check if the connection is successful
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}
	log.Println("Connected to the database")
}

func NewRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	return router
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := DB.Query("SELECT id, title, author FROM books")
	if err != nil {
		log.Println("Error querying database:", err)
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			log.Println("Error scanning row:", err)
			http.Error(w, "Error scanning row", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error iterating over rows:", err)
		http.Error(w, "Error iterating over rows", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var book models.Book
	err := DB.QueryRow("SELECT id, title, author FROM books WHERE id = $1", params["id"]).Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		log.Println("Error querying database:", err)
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	_, err := DB.Exec("INSERT INTO books (title, author) VALUES ($1, $2)", book.Title, book.Author)
	if err != nil {
		log.Println("Error inserting into database:", err)
		http.Error(w, "Error inserting into database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	_, err := DB.Exec("UPDATE books SET title = $1, author = $2 WHERE id = $3", book.Title, book.Author, params["id"])
	if err != nil {
		log.Println("Error updating database:", err)
		http.Error(w, "Error updating database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	_, err := DB.Exec("DELETE FROM books WHERE id = $1", params["id"])
	if err != nil {
		log.Println("Error deleting from database:", err)
		http.Error(w, "Error deleting from database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Book deleted successfully")
}

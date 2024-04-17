// models/book.go

package models

import "time"

// Book represents a book in the library.
type Book struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"published_date"`
}

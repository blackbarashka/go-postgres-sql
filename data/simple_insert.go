package data

import (
	"context"

	"postgres/book"

	"github.com/jackc/pgx/v5"
)

func InsertTable(ctx context.Context, conn *pgx.Conn, newBook *book.Book) error {
	sqlQuery := `
	INSERT INTO booksunique (title, author, description, comleted, created_at) 
	VALUES ($1, $2, $3, $4, $5)
;
`
	_, err := conn.Exec(ctx, sqlQuery, newBook.Title, newBook.Author, newBook.Description, newBook.Completed, newBook.Created_at)
	if err != nil {
		return err
	}
	return nil
}

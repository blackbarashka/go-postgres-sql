package data

import (
	"context"
	"postgres/book"

	"github.com/jackc/pgx/v5"
)

func UpdateByIdTable(ctx context.Context, conn *pgx.Conn, id int, newBook *book.Book) error {
	sqlQuery := `
	UPDATE booksunique
	SET title=$1, author=$2, description=$3, comleted=$4, created_at=$5, completed_at=$6
	WHERE id=$7
`
	_, err := conn.Exec(ctx, sqlQuery, newBook.Title, newBook.Author, newBook.Description, newBook.Completed, newBook.Created_at, newBook.Completed_at, id)
	return err
}

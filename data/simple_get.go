package data

import (
	"context"
	"postgres/book"

	"github.com/jackc/pgx/v5"
)

func GetTable(ctx context.Context, conn *pgx.Conn) ([]book.Book, error) {
	sqlQuery := `
	SELECT title, author, description, comleted, created_at, completed_at 
	FROM booksunique
`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []book.Book
	for rows.Next() {
		var b book.Book
		err := rows.Scan(&b.Title, &b.Author, &b.Description, &b.Completed, &b.Created_at, &b.Completed_at)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

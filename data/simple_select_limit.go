package data

import (
	"context"
	"postgres/book"

	"github.com/jackc/pgx/v5"
)

func SelectLimitTable(ctx context.Context, conn *pgx.Conn, limit, offset int) ([]book.Book, error) {
	sqlQuery := `
    SELECT title, author, description, comleted, created_at, completed_at
    FROM booksunique
	LIMIT $1
	OFFSET $2
	`
	rows, err := conn.Query(ctx, sqlQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	var books []book.Book
	for rows.Next() {
		var b book.Book
		err := rows.Scan(
			&b.Title,
			&b.Author,
			&b.Description,
			&b.Completed,
			&b.Created_at,
			&b.Completed_at)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func GetCountRow(ctx context.Context, conn *pgx.Conn) (int, error) {
	var count int
	err := conn.QueryRow(ctx, "SELECT COUNT(*) FROM booksunique").Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

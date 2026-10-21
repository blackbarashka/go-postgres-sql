package data

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteByIds(ctx context.Context, conn *pgx.Conn, ids []int) error {
	sqlQuery := `
    DELETE FROM booksunique 
    WHERE id=$1`

	for _, id := range ids {
		_, err := conn.Exec(ctx, sqlQuery, id)
		if err != nil {
			return err
		}
	}
	return nil
}

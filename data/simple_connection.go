package data

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {
	//change.
	conn, err := pgx.Connect(ctx, "postgres://postgres:admin@localhost:5432/postgres")
	if err != nil {
		return nil, err
	}
	err_ := conn.Ping(ctx)
	if err_ != nil {
		return nil, err_
	}
	fmt.Println("Successfully connected to postgres")
	return conn, err
}
